package main

import (
	"log"
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

type testChecker struct {
	machines []*pbs.Config
	running  bool
}

func (t testChecker) assess(server string) (*pbs.JobList, *pbs.Config) {
	if server == "server1" {
		return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test1"}, Running: true}}}, t.machines[0]
	}
	return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test2"}, Running: t.running}}}, t.machines[1]
}

func (t testChecker) discover() *pbd.ServiceList {
	return &pbd.ServiceList{Services: []*pbd.RegistryEntry{&pbd.RegistryEntry{Identifier: "server1", Name: "gobuildslave"}, &pbd.RegistryEntry{Identifier: "server2", Name: "gobuildslave"}}}
}

func TestPullData(t *testing.T) {
	status, _ := getFleetStatus(&testChecker{machines: []*pbs.Config{&pbs.Config{}, &pbs.Config{}}})
	if val, ok := status["server1"]; !ok || len(val.Details) != 1 {
		t.Errorf("Status has come back bad: %v", status)
	}
}

func TestFleetCount(t *testing.T) {
	status, _ := getFleetStatus(&testChecker{machines: []*pbs.Config{&pbs.Config{}, &pbs.Config{}}, running: false})
	if val, ok := status["server2"]; !ok || len(val.Details) != 0 {
		t.Errorf("Status has come back good when not running: %v", val)
	}
}

func TestLoadConfig(t *testing.T) {
	f := "testdata/testconfig.pb"
	c, err := loadConfig(f)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}

	if len(c.Intents) != 2 {
		t.Errorf("Config parsing failed: %v", c)
	}
}

func TestLoadNamedConfig(t *testing.T) {
	f := "testdata/config_test.pb"
	c, err := loadConfig(f)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}

	if len(c.Intents) != 1 || c.Intents[0].Spec.Name != "github.com/brotherlogic/cardserver" {
		t.Errorf("Config parsing failed: %v", c)
	}
}

func TestLoadMainConfig(t *testing.T) {
	f := "config.pb"
	c, err := loadConfig(f)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}

	if len(c.Intents) == 0 {
		t.Errorf("Config parsing failed: %v", c)
	}

	log.Printf("READ CONFIG")
	log.Printf("%v", c)
}

func TestRunJob(t *testing.T) {
	i1 := &pb.Intent{Spec: &pbs.JobSpec{Name: "testing"}, Count: 2}
	jobs := runJobs(&pb.Config{Intents: []*pb.Intent{i1}})
	log.Printf("RUN ON: %v", jobs)
	if len(jobs) != 2 || jobs[0].Name != "testing" || jobs[1].Name != "testing" {
		t.Errorf("Run jobs produced bad result: %v", jobs)
	}
}

func TestDiff(t *testing.T) {
	i1 := &pb.Intent{Spec: &pbs.JobSpec{Name: "testing"}, Count: 1}
	c1 := &pb.Config{Intents: []*pb.Intent{i1}}
	c2 := &pb.Config{Intents: []*pb.Intent{}}

	diff := configDiff(c1, c2)
	if len(diff.Intents) != 1 || diff.Intents[0].Spec.Name != "testing" {
		t.Errorf("Error in diff: %v", diff)
	}
}

func TestDiffWhenMatch(t *testing.T) {
	i1 := &pb.Intent{Spec: &pbs.JobSpec{Name: "testing"}, Count: 1}
	c1 := &pb.Config{Intents: []*pb.Intent{i1}}
	c2 := &pb.Config{Intents: []*pb.Intent{i1}}

	diff := configDiff(c1, c2)
	if len(diff.Intents) != 0 && diff.Intents[0].Count != 0 {
		t.Errorf("Error in diff: %v", diff)
	}
}

func TestLoadOntoDiskMachine(t *testing.T) {
	conf := &pbs.JobSpec{Name: "needsdisk", Disk: 1024}

	machine1 := &pbs.Config{Disk: 100}
	machine2 := &pbs.Config{Disk: 2000}

	server := chooseServer(conf, testChecker{machines: []*pbs.Config{machine1, machine2}})
	if server != "server2" {
		t.Errorf("Failed to select correct server: %v", server)
	}
}

func TestDoubleLoadServer(t *testing.T) {
	conf := &pbs.JobSpec{Name: "test1"}
	machine1 := &pbs.Config{Disk: 100}
	machine2 := &pbs.Config{Disk: 100}
	server := chooseServer(conf, testChecker{machines: []*pbs.Config{machine1, machine2}})

	if server == "server1" {
		t.Errorf("Loaded on server1 even though job was running there: %v", server)
	}
}

func TestMissServer(t *testing.T) {
	conf := &pbs.JobSpec{Name: "needsdisk", Disk: 1024}

	machine1 := &pbs.Config{Disk: 100}
	machine2 := &pbs.Config{Disk: 100}

	server := chooseServer(conf, testChecker{machines: []*pbs.Config{machine1, machine2}})
	if server != "" {
		t.Errorf("Found a server even though one is not there: %v", server)
	}
}
