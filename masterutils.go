package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

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
	}

	s.lastWorldRun = time.Now().Unix()

	for server, seen := range s.serverMap {
		if time.Now().Sub(seen) > s.timeChange {
			s.RaiseIssue(ctx, "Missing Server", fmt.Sprintf("%v is missing.", server), false)
		}
	}

	return nil
}
