package main

import "testing"
import pb "github.com/brotherlogic/gobuildmaster/proto"
import pbs "github.com/brotherlogic/gobuildslave/proto"
import pbd "github.com/brotherlogic/discovery/proto"

type testChecker struct{}

func (t *testChecker) assess(server string) *pbs.JobList {
	if server == "server1" {
		return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test1"}}}}
	}
	return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test2"}}}}
}

func (t *testChecker) discover() *pbd.ServiceList {
	return &pbd.ServiceList{Services: []*pbd.RegistryEntry{&pbd.RegistryEntry{Identifier: "server1", Name: "gobuildslave"}, &pbd.RegistryEntry{Identifier: "server2", Name: "gobuildslave"}}}
}

func TestPullData(t *testing.T) {
	status := getFleetStatus(&testChecker{})
	if val, ok := status["server1"]; !ok || len(val.Details) != 1 {
		t.Errorf("Status has come back bad: %v", status)
	}
}

func TestRunJob(t *testing.T) {
	i1 := &pb.Intent{Spec: &pbs.JobSpec{Name: "testing"}, Masters: 2}
	jobs := runJobs(&pb.Config{Intents: []*pb.Intent{i1}})
	if len(jobs) != 2 || jobs[0].Name != "testing" || jobs[1].Name != "testing" {
		t.Errorf("Run jobs produced bad result: %v", jobs)
	}
}

func TestDiff(t *testing.T) {
	i1 := &pb.Intent{Spec: &pbs.JobSpec{Name: "testing"}}
	c1 := &pb.Config{Intents: []*pb.Intent{i1}}
	c2 := &pb.Config{Intents: []*pb.Intent{}}

	diff := configDiff(c1, c2)
	if len(diff.Intents) != 1 || diff.Intents[0].Spec.Name != "testing" {
		t.Errorf("Error in diff: %v", diff)
	}
}

func TestClean(t *testing.T) {
	blank()
}
