package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

func InitTestServer() *Server {
	s := Init(&pb.Config{})
	s.PrepServer()
	s.SkipLog = true
	return s
}

func TestBuildWorld(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld(context.Background())
	if _, ok := s.world["testjob"]["testserver"]; !ok {
		t.Fatalf("World has not been built correctly: %v", s.world)
	}
}

func TestBuildWorldFailSlaves(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{failGetSlaves: true, running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld(context.Background())
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}

func TestBuildWorldFailJobs(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{failGetJobs: true, running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld(context.Background())
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}

func TestAlertOnMissingServer(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	g.running["server1"] = []*pbs.JobAssignment{&pbs.JobAssignment{Server: "server1", Job: &pbs.Job{Name: "testjob"}}, &pbs.JobAssignment{Server: "server2", Job: &pbs.Job{Name: "testjob2"}}}
	g.running["server2"] = []*pbs.JobAssignment{&pbs.JobAssignment{Server: "server1", Job: &pbs.Job{Name: "testjob"}}, &pbs.JobAssignment{Server: "server2", Job: &pbs.Job{Name: "testjob2"}}}

	s.getter = g
	s.buildWorld(context.Background())

	g = &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	g.running["server1"] = []*pbs.JobAssignment{&pbs.JobAssignment{Server: "server1", Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g
	s.buildWorld(context.Background())

	if s.AlertsFired != 1 {
		t.Errorf("No alert on missing server: %v", s.AlertsFired)
	}
}
