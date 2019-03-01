package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

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
	master(ctx context.Context, entry *pbd.RegistryEntry, master bool) (bool, error)
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
						if r.Category == req.Category && r.Properties == req.Properties {
							localmatch = true
							s.Log(fmt.Sprintf("MATCH %v and %v for %v on %v", r, req, job.Name, services.Services[i].Identifier))
						}
					}

					if !localmatch {
						allmatch = false
					}
				}
				if allmatch {
					s.Log(fmt.Sprintf("MATCH %v and %v for %v on %v", job.Requirements, requirements, job.Name, services.Services[i].Identifier))
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

func loadConfig(f string) (*pb.Config, error) {
	toload := &pb.Config{}
	bytes, _ := ioutil.ReadFile(f)
	proto.UnmarshalText(string(bytes), toload)
	return toload, nil
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
