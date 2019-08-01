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

	s.worldMutex.Lock()
	for _, job := range jobs {
		if _, ok := s.world[job.Job.GetName()]; !ok {
			s.world[job.Job.GetName()] = make(map[string]struct{})
		}
		s.world[job.Job.GetName()][server.GetIdentifier()] = struct{}{}
	}
	s.worldMutex.Unlock()
	return nil
}

func (s *Server) buildWorld(ctx context.Context) error {
	s.worldMutex.Lock()
	s.world = make(map[string]map[string]struct{})
	slaves, err := s.getter.getSlaves()
	if err != nil {
		s.worldMutex.Unlock()
		return err
	}
	s.worldMutex.Unlock()

	for _, server := range slaves.GetServices() {
		err := s.updateWorld(ctx, server)
		if err != nil {
			return err
		}
	}

	s.lastWorldRun = time.Now().Unix()

	for server, seen := range s.serverMap {
		if time.Now().Sub(seen) > s.timeChange {
			info, _ := s.State(ctx, &pb.Empty{})
			infoString := ""
			for _, str := range info.GetStates() {
				infoString += fmt.Sprintf("%v = %v\n", str.Key, str)
			}
			s.RaiseIssue(ctx, "Missing Server", fmt.Sprintf("%v is missing duration: %v.\n%v", server, time.Now().Sub(seen), infoString), false)
		}
	}

	for job, versions := range s.world {
		count := int32(0)
		found := false
		for _, cjob := range s.config.Nintents {
			if cjob.Job.Name == job {
				found = true
				count = cjob.Count
			}
		}

		if !found {
			s.Log(fmt.Sprintf("Could not find %v", job))
		}

		if count > 0 && int32(len(versions))-count > 1 {
			s.RaiseIssue(ctx, "Too many jobs", fmt.Sprintf("%v has too many versions running", job), false)
		}
	}

	return nil
}
