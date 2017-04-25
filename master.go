package main

import pb "github.com/brotherlogic/gobuildmaster/proto"
import pbs "github.com/brotherlogic/gobuildslave/proto"

func configDiff(cm, cs *pb.Config) *pb.Config {
	return cm
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
