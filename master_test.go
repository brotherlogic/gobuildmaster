package main

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

type testChecker struct {
	machines []*pbs.Config
	running  bool
}

func (t testChecker) assess(ctx context.Context, server string) (*pbs.JobList, *pbs.Config) {
	log.Printf("ASSESS %v", server)
	if server == "server1" {
		return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test1"}}}}, t.machines[0]
	}
	return &pbs.JobList{Details: []*pbs.JobDetails{&pbs.JobDetails{Spec: &pbs.JobSpec{Name: "test1"}}}}, t.machines[1]
}

func (t testChecker) discover() *pbd.ServiceList {
	return &pbd.ServiceList{Services: []*pbd.RegistryEntry{&pbd.RegistryEntry{Identifier: "server1", Name: "gobuildslave"}, &pbd.RegistryEntry{Identifier: "server2", Name: "gobuildslave"}}}
}

func (t testChecker) getprev() []string {
	return make([]string, 0)
}

func (t testChecker) setprev(v []string) {
	// Do nothing
}

func (t testChecker) master(ctx context.Context, entry *pbd.RegistryEntry, master bool) (bool, error) {
	// Do nothing
	return true, nil
}

func TestPullData(t *testing.T) {
	status, _ := getFleetStatus(context.Background(), &testChecker{machines: []*pbs.Config{&pbs.Config{}, &pbs.Config{}}})
	if val, ok := status["server1"]; !ok || len(val.Details) != 1 {
		t.Errorf("Status has come back bad: %v", status)
	}
}

func TestFleetCount(t *testing.T) {
	status, _ := getFleetStatus(context.Background(), &testChecker{machines: []*pbs.Config{&pbs.Config{}, &pbs.Config{}}, running: false})
	if val, ok := status["server2"]; !ok || len(val.Details) != 1 {
		t.Errorf("Status has come back good when not running: %v", val)
	}
}

func TestLoadMainConfig(t *testing.T) {
	f := "config.pb"
	c, err := loadConfig(f)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}

	found := false
	for _, i := range c.Nintents {
		if i.Job.Name == "buildserver" {
			found = true
		}
	}
	if !found {
		t.Errorf("Cannot find buildserver: %v", c.Nintents)
	}
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

	server := chooseServer(context.Background(), conf, testChecker{machines: []*pbs.Config{machine1, machine2}})
	if server != "server2" {
		t.Errorf("Failed to select correct server: %v", server)
	}
}

func TestLoadOntoAlreadyRunning(t *testing.T) {
	conf := &pbs.JobSpec{Name: "test1"}

	server := chooseServer(context.Background(), conf, testChecker{machines: []*pbs.Config{&pbs.Config{Disk: 100}, &pbs.Config{Disk: 1000}}})
	if server != "" {
		t.Errorf("Failed to select correct server: %v", server)
	}
}

func TestLoadOntoExternalMachine(t *testing.T) {
	conf := &pbs.JobSpec{Name: "needsexternal", External: true}

	machine1 := &pbs.Config{Disk: 100, External: false}
	machine2 := &pbs.Config{Disk: 100, External: true}

	server := chooseServer(context.Background(), conf, testChecker{machines: []*pbs.Config{machine1, machine2}})
	if server != "server2" {
		t.Errorf("Failed to select correct server: %v", server)
	}
}

func TestDoubleLoadServer(t *testing.T) {
	conf := &pbs.JobSpec{Name: "test1"}
	machine1 := &pbs.Config{Disk: 100}
	machine2 := &pbs.Config{Disk: 100}
	server := chooseServer(context.Background(), conf, testChecker{machines: []*pbs.Config{machine1, machine2}})

	if server == "server1" {
		t.Errorf("Loaded on server1 even though job was running there: %v", server)
	}
}

func TestMissServer(t *testing.T) {
	conf := &pbs.JobSpec{Name: "needsdisk", Disk: 1024}

	machine1 := &pbs.Config{Disk: 100}
	machine2 := &pbs.Config{Disk: 100}

	server := chooseServer(context.Background(), conf, testChecker{machines: []*pbs.Config{machine1, machine2}})
	if server != "" {
		t.Errorf("Found a server even though one is not there: %v", server)
	}
}

type testGetter struct {
	failGetSlaves bool
	failGetJobs   bool
	running       map[string][]*pbs.JobAssignment
	config        map[string][]*pbs.Requirement
}

func (t *testGetter) getSlaves() (*pbd.ServiceList, error) {
	if t.failGetSlaves {
		return &pbd.ServiceList{}, errors.New("Built to fail")
	}

	list := &pbd.ServiceList{}
	for key := range t.running {
		list.Services = append(list.Services, &pbd.RegistryEntry{Identifier: key})
	}
	return list, nil
}

func (t *testGetter) getJobs(ctx context.Context, e *pbd.RegistryEntry) ([]*pbs.JobAssignment, error) {
	if t.failGetJobs {
		return []*pbs.JobAssignment{}, errors.New("Built to fail")
	}

	if val, ok := t.running[e.Identifier]; ok {
		return val, nil
	}
	return make([]*pbs.JobAssignment, 0), nil
}

func (t *testGetter) getConfig(ctx context.Context, e *pbd.RegistryEntry) ([]*pbs.Requirement, error) {
	if val, ok := t.config[e.Identifier]; ok {
		return val, nil
	}
	return make([]*pbs.Requirement, 0), nil
}

func TestFirstSelect(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	tg.running["badserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "runner"}}}
	tg.running["goodserver"] = []*pbs.JobAssignment{}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner"}, tg)
	if server != "goodserver" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestDiskSelect(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	tg.running["badserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "runner"}}}
	tg.running["goodserver"] = []*pbs.JobAssignment{}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner"}, tg)
	if server != "goodserver" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestReqSelectWithLimits(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment), config: make(map[string][]*pbs.Requirement)}
	tg.running["badserver"] = []*pbs.JobAssignment{}
	tg.running["goodserver"] = []*pbs.JobAssignment{}
	tg.config["goodserver"] = []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "backup"}}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner", Requirements: []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "maindisk"}}}, tg)
	if server != "" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestReqSelect(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment), config: make(map[string][]*pbs.Requirement)}
	tg.running["goodserver"] = []*pbs.JobAssignment{}
	tg.config["goodserver"] = []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "maindisk"}}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner", Requirements: []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "maindisk"}}}, tg)
	if server != "goodserver" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestReqSelectExternal(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment), config: make(map[string][]*pbs.Requirement)}
	tg.running["discover"] = []*pbs.JobAssignment{}
	tg.config["discover"] = []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_EXTERNAL, Properties: "external_ready"}}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner", Requirements: []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_EXTERNAL, Properties: "external_ready"}}}, tg)
	if server != "discover" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestReqSelectFail(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment), config: make(map[string][]*pbs.Requirement)}
	tg.running["goodserver"] = []*pbs.JobAssignment{}
	tg.config["goodserver"] = []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "maindisker"}, &pbs.Requirement{Category: pbs.RequirementCategory_ACCESS_POINT, Properties: "70:3A:CB:17:CF:BB"}}

	server := s.selectServer(context.Background(), &pbs.Job{Name: "runner", Requirements: []*pbs.Requirement{&pbs.Requirement{Category: pbs.RequirementCategory_DISK, Properties: "maindisk"}}}, tg)
	if server != "" {
		t.Errorf("Wrong server selected: %v", server)
	}
}

func TestAddAccessPoints(t *testing.T) {
	s := InitTestServer()
	s.accessPoints["blah"] = time.Now().Add(-time.Hour * 2)
	for _, str := range []string{"70:3A:CB:17:CF:BB", "70:3A:CB:17:CC:D3", "70:3A:CB:17:CE:E3", "70:3A:CB:17:CF:BF", "blah"} {
		s.addAccessPoint(context.Background(), str)
	}
}
