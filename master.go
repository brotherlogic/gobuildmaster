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

	for _, service := range c.discover().Services {
		log.Printf("Found: %v", service)
		if service.Name == "gobuildslave" {
			log.Printf("HERE: %v", service)
			joblist, config := c.assess(service.Identifier)
			resJ[service.Identifier] = joblist
			resC[service.Identifier] = config

			sub := 0
			for i := range resJ[service.Identifier].Details {
				if !resJ[service.Identifier].Details[i-sub].Running {
					resJ[service.Identifier].Details = append(resJ[service.Identifier].Details[:(i-sub)], resJ[service.Identifier].Details[(i-sub)+1:]...)
					sub++
				}
			}
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
			_, sc := c.assess(service.Identifier)
			if sc.Disk > job.Disk {
				return service.Identifier
			}
		}
	}
	return ""
}

func configDiff(cm, cs *pb.Config) *pb.Config {
	retConfig := &pb.Config{}

	log.Printf("COMPARE %v to %v", cm, cs)
	for _, entry := range cm.Intents {
		nIntent := &pb.Intent{}
		nIntent.Spec = entry.Spec
		nIntent.Masters = entry.Masters
		nIntent.Slaves = entry.Slaves
		retConfig.Intents = append(retConfig.Intents, nIntent)
	}

	log.Printf("FIRST PASS: %v", retConfig)
	for _, entry := range cs.Intents {
		for _, pair := range retConfig.Intents {
			if entry.Spec.Name == pair.Spec.Name {
				log.Printf("FOUND: %v", pair)
				pair.Masters -= entry.Masters
				pair.Slaves -= entry.Slaves
				log.Printf("RESULT: %v", retConfig)
			}
		}
	}

	log.Printf("RETURNING: %v", retConfig)
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
		for i := 0; i < int(j.Masters); i++ {
			jobs = append(jobs, j.Spec)
		}
	}

	return jobs
}
