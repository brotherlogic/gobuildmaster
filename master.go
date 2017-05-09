package main

import (
	"io/ioutil"
	"log"

	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
	"github.com/golang/protobuf/proto"

	pbd "github.com/brotherlogic/discovery/proto"
)

type checker interface {
	assess(server string) *pbs.JobList
	discover() *pbd.ServiceList
}

func getFleetStatus(c checker) map[string]*pbs.JobList {
	res := make(map[string]*pbs.JobList)

	for _, service := range c.discover().Services {
		log.Printf("Found: %v", service)
		if service.Name == "gobuildslave" {
			joblist := c.assess(service.Identifier)
			log.Printf("LIST %v", joblist)
			res[service.Identifier] = joblist
		}
	}

	return res
}

func configDiff(cm, cs *pb.Config) *pb.Config {
	return cm
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

func blank() {

}
