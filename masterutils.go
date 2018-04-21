package main

import (
	"fmt"
	"time"
)

func (s *Server) buildWorld() {
	s.worldMutex.Lock()
	s.world = make(map[string]map[string]struct{})
	slaves, err := s.getter.getSlaves()
	if err != nil {
		s.Log(fmt.Sprintf("Error getting slaves: %v", err))
		return
	}

	for _, server := range slaves.GetServices() {
		jobs, err := s.getter.getJobs(server)
		if err != nil {
			return
		}
		for _, job := range jobs {
			if _, ok := s.world[job.Job.GetName()]; !ok {
				s.world[job.Job.GetName()] = make(map[string]struct{})
			}
			s.world[job.Job.GetName()][server.GetIdentifier()] = struct{}{}
		}
	}

	s.lastWorldRun = time.Now().Unix()
	s.worldMutex.Unlock()
}
