package main

import (
	"fmt"
	"time"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			s.Log(fmt.Sprintf("Unable to reach %v -> %v", server, err))
			continue
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
		time.Sleep(time.Second * 2)
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
				s.Log(fmt.Sprintf("Running %v -> %v", intent.GetJob().GetName(), err))
				code := status.Convert(err).Code()
				if code != codes.OK {
					s.decisions[intent.GetJob().GetName()] = fmt.Sprintf("Error running job: %v", err)
				}
				if code != codes.OK && code != codes.FailedPrecondition {
					return err
				}
			} else {
				s.decisions[intent.GetJob().GetName()] = fmt.Sprintf("Missing requirement")
				s.Log(fmt.Sprintf("Missing requirements for %v -> %v vs %v", intent.GetJob().GetName(), intent.GetJob().GetRequirements(), localConfig))
			}
		}
	}

	return nil
}

func (s *Server) check(ctx context.Context, i *pb.NIntent, counts map[string]int, ls *pbd.RegistryEntry) error {
	// We register as best effort - and throw it into the background
	go func() {
		ctx, cancel := utils.ManualContext("gmb-register", time.Minute)
		defer cancel()
		s.registerJob(ctx, i)
	}()

	if i.Redundancy == pb.Redundancy_GLOBAL {
		if s.Registry.Identifier != i.NotOnServer {
			return s.runJob(ctx, i.GetJob(), ls)
		}
		return nil
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
