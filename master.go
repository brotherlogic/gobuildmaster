package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

type getter interface {
	getSlaves() (*pbd.ServiceList, error)
	getJobs(context.Context, *pbd.RegistryEntry) ([]*pbs.JobAssignment, error)
	getConfig(context.Context, *pbd.RegistryEntry) ([]*pbs.Requirement, error)
}

type checker interface {
	assess(ctx context.Context, server string) (*pbs.JobList, *pbs.Config)
	discover() *pbd.ServiceList
	master(entry *pbd.RegistryEntry, master bool) (bool, error)
	setprev([]string)
	getprev() []string
}

func getFleetStatus(ctx context.Context, c checker) (map[string]*pbs.JobList, map[string]*pbs.Config) {
	resJ := make(map[string]*pbs.JobList)
	resC := make(map[string]*pbs.Config)

	curr := make([]string, 0)
	for _, service := range c.discover().Services {
		if service.Name == "gobuildslave" {
			curr = append(curr, service.Identifier)
			joblist, config := c.assess(ctx, service.Identifier)
			resJ[service.Identifier] = joblist
			resC[service.Identifier] = config
		}
	}

	c.setprev(curr)

	return resJ, resC
}

// Find the first available server
func chooseServer(ctx context.Context, job *pbs.JobSpec, c checker) string {
	services := c.discover().Services
	for _, i := range rand.Perm(len(services)) {
		service := services[i]
		if service.Name == "gobuildslave" && (job.GetServer() == "" || job.GetServer() == service.GetIdentifier()) {
			jobs, sc := c.assess(ctx, service.Identifier)

			//Don't accept a server which is already running this job
			jobfine := true
			for _, j := range jobs.Details {
				if j.Spec.Name == job.Name {
					jobfine = false
				}

			}
			if jobfine {
				if sc.Disk > job.Disk && (!job.External || sc.External) && (!job.GetCds() || sc.GetSupportsCds()) {
					return service.Identifier
				}
			}
		}
	}
	return ""
}

func (s *Server) addAccessPoint(ctx context.Context, ap string) {
	s.accessPointsMutex.Lock()
	defer s.accessPointsMutex.Unlock()

	for key, val := range s.accessPoints {
		if time.Now().Sub(val) > time.Hour*24 {
			s.RaiseIssue(fmt.Sprintf("Access point Missing"), fmt.Sprintf("%v has been missing since %v", key, val))
		}
	}

	switch ap {
	case "70:3A:CB:17:CF:BB":
		s.accessPoints["LR2"] = time.Now()
	case "70:3A:CB:17:CC:CF":
		s.accessPoints["Bedroom"] = time.Now()
	case "70:3A:CB:17:CE:E3":
		s.accessPoints["LR"] = time.Now()
	case "70:3A:CB:17:CF:BF":
		s.accessPoints["LR2"] = time.Now()
	}
}

// Find the first available server
func (s *Server) selectServer(ctx context.Context, job *pbs.Job, g getter) string {
	services, _ := g.getSlaves()
	for _, i := range rand.Perm(len(services.Services)) {
		jobs, _ := g.getJobs(ctx, services.Services[i])
		//Don't accept a server which is already running this job
		jobfine := true
		for _, j := range jobs {
			if j.Job.Name == job.Name {
				jobfine = false
			}

		}
		if jobfine {
			requirements, err := g.getConfig(ctx, services.Services[i])
			if err == nil {
				allmatch := true
				for _, req := range job.Requirements {
					localmatch := false
					for _, r := range requirements {
						if r.Category == pbs.RequirementCategory_ACCESS_POINT {
							s.addAccessPoint(ctx, r.Properties)
						}
						if r.Category == req.Category && r.Properties == req.Properties {
							localmatch = true

						}
					}

					if !localmatch {
						allmatch = false
					}
				}
				if allmatch {
					return services.Services[i].Identifier
				}
			}
		}
	}
	return ""
}

func configDiff(cm, cs *pb.Config) *pb.Config {
	retConfig := &pb.Config{}

	for _, entry := range cm.Intents {
		nIntent := &pb.Intent{}
		nIntent.Spec = entry.Spec
		nIntent.Count = entry.Count
		retConfig.Intents = append(retConfig.Intents, nIntent)
	}

	for _, entry := range cs.Intents {
		for _, pair := range retConfig.Intents {
			if entry.Spec.Name == pair.Spec.Name {
				pair.Count -= entry.Count
			}
		}
	}

	return retConfig
}

func loadConfig() (*pb.Config, error) {
	toload := &pb.Config{}
	data, _ := Asset("config.pb")
	err := proto.UnmarshalText(string(data), toload)
	return toload, err
}

func runJobs(c *pb.Config) []*pbs.JobSpec {
	var jobs []*pbs.JobSpec
	for _, j := range c.Intents {
		for i := 0; i < int(j.Count); i++ {
			jobs = append(jobs, j.Spec)
		}
	}

	return jobs
}
