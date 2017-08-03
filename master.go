package main

import (
	"io/ioutil"
	"log"
	"math/rand"

	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
	"github.com/golang/protobuf/proto"

	pbd "github.com/brotherlogic/discovery/proto"
)

type checker interface {
	assess(server string) (*pbs.JobList, *pbs.Config)
	discover() *pbd.ServiceList
}

func getFleetStatus(c checker) (map[string]*pbs.JobList, map[string]*pbs.Config) {
	resJ := make(map[string]*pbs.JobList)
	resC := make(map[string]*pbs.Config)

	log.Printf("Discovering")
	for _, service := range c.discover().Services {
		if service.Name == "gobuildslave" {
			log.Printf("Found server on %v", service.Identifier)
			joblist, config := c.assess(service.Identifier)
			resJ[service.Identifier] = joblist
			resC[service.Identifier] = config
		}
	}

	return resJ, resC
}

// Find the first available server
func chooseServer(job *pbs.JobSpec, c checker) string {
	services := c.discover().Services
	for i := range rand.Perm(len(services)) {
		service := services[i]
		if service.Name == "gobuildslave" {
			jobs, sc := c.assess(service.Identifier)

			//Don't accept a server which is already running this job
			jobfine := true
			for _, j := range jobs.Details {
				if j.Spec.Name == job.Name {
					jobfine = false
				}
			}
			if jobfine {
				if sc.Disk > job.Disk && (!job.External || sc.External) {
					return service.Identifier
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
