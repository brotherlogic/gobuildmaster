package main

import (
	"fmt"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	"golang.org/x/net/context"
)

func (s *Server) updateWorld(ctx context.Context, server *pbd.RegistryEntry) ([]string, error) {
	jobs, err := s.getter.getJobs(ctx, server)
	if err != nil {
		return []string{}, err
	}

	slaveMap := []string{}
	for _, job := range jobs {
		slaveMap = append(slaveMap, job.GetJob().GetName())
	}

	return slaveMap, nil
}

func (s *Server) adjustWorld(ctx context.Context) error {
	slaves, err := s.getter.getSlaves()
	if err != nil {
		return err
	}

	var ourSlave *pbd.RegistryEntry
	for _, slave := range slaves.GetServices() {
		if slave.Identifier == s.Registry.Identifier {
			ourSlave = slave
		}
	}
	if ourSlave == nil {
		return fmt.Errorf("Cannot locate local gbs from %v", slaves)
	}

	if len(slaves.GetServices()) == 0 {
		return fmt.Errorf("Unable to locate any slaves")
	}

	jobCount := make(map[string]int)
	ourjobs := make(map[string]bool)
	for _, server := range slaves.GetServices() {
		slaves, err := s.updateWorld(ctx, server)
		if err != nil {
			return err
		}

		for _, j := range slaves {
			jobCount[j]++
			if server.Identifier == s.Registry.Identifier {
				ourjobs[j] = true
			}
		}
	}

	localConfig, err := s.getter.getConfig(ctx, ourSlave)
	if err != nil {
		return err
	}

	for _, intent := range s.config.Nintents {
		if !ourjobs[intent.GetJob().GetName()] {
			allmatch := true
			for _, req := range intent.GetJob().GetRequirements() {
				localmatch := false
				for _, r := range localConfig {
					if r.Category == req.Category && r.Properties == req.Properties {
						localmatch = true
					}
				}

				if !localmatch {
					allmatch = false
				}
			}

			if allmatch {
				err := s.check(ctx, intent, jobCount, ourSlave)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *Server) check(ctx context.Context, i *pb.NIntent, counts map[string]int, ls *pbd.RegistryEntry) error {
	if i.Redundancy == pb.Redundancy_GLOBAL {
		return s.runJob(ctx, i.GetJob(), ls)
	}

	if i.Redundancy == pb.Redundancy_REDUNDANT {
		if counts[i.GetJob().GetName()] < 3 {
			return s.runJob(ctx, i.GetJob(), ls)
		}
	}

	if counts[i.GetJob().GetName()] < int(i.Count) {
		return s.runJob(ctx, i.GetJob(), ls)
	}

	return nil
}
