package main

import (
	"fmt"
	"time"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/goserver/proto"
	"golang.org/x/net/context"
)

func (s *Server) updateWorld(ctx context.Context, server *pbd.RegistryEntry) error {
	s.serverMap[server.Identifier] = time.Now()

	jobs, err := s.getter.getJobs(ctx, server)
	if err != nil {
		return err
	}

	s.slaveMap[server.GetIdentifier()] = []string{}
	for _, job := range jobs {
		s.slaveMap[server.GetIdentifier()] = append(s.slaveMap[server.GetIdentifier()], job.GetJob().GetName())
	}
	return nil
}

func (s *Server) buildWorld(ctx context.Context) error {
	slaves, err := s.getter.getSlaves()
	if err != nil {
		return err
	}

	if len(slaves.GetServices()) == 0 {
		return fmt.Errorf("Unable to locate any slaves")
	}

	for _, server := range slaves.GetServices() {
		err := s.updateWorld(ctx, server)
		if err != nil {
			s.Log(fmt.Sprintf("Error getting world: %v", err))
		}
	}

	// Rebuild world
	s.worldMutex.Lock()
	s.world = make(map[string]map[string]struct{})
	for server, jobs := range s.slaveMap {
		for _, job := range jobs {
			if _, ok := s.world[job]; !ok {
				s.world[job] = make(map[string]struct{})
			}
			s.world[job][server] = struct{}{}
		}
	}
	s.worldMutex.Unlock()
	s.lastWorldRun = time.Now().Unix()

	for server, seen := range s.serverMap {
		if time.Now().Sub(seen) > s.timeChange {
			info, _ := s.State(ctx, &pb.Empty{})
			infoString := fmt.Sprintf("%v\n\n", slaves)
			for _, str := range info.GetStates() {
				infoString += fmt.Sprintf("%v = %v\n", str.Key, str)
			}
			s.RaiseIssue(ctx, "Missing Server", fmt.Sprintf("%v is missing duration: %v.\n%v", server, time.Now().Sub(seen), infoString), false)
		}
	}

	for job, versions := range s.world {
		count := int32(0)
		for _, cjob := range s.config.Nintents {
			if cjob.Job.Name == job {
				count = cjob.Count
			}
		}

		if count > 0 && int32(len(versions))-count > 1 {
			s.RaiseIssue(ctx, "Too many jobs", fmt.Sprintf("%v has too many versions running", job), false)
		}
	}

	return nil
}
