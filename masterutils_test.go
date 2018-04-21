package main

import (
	"testing"

	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

func InitTestServer() *Server {
	s := Init(&pb.Config{})
	s.SkipLog = true
	return s
}

func TestBuildWorld(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld()
	if _, ok := s.world["testjob"]["testserver"]; !ok {
		t.Fatalf("World has not been built correctly: %v", s.world)
	}
}

func TestBuildWorldFailSlaves(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{failGetSlaves: true, running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld()
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}

func TestBuildWorldFailJobs(t *testing.T) {
	s := InitTestServer()
	g := &testGetter{failGetJobs: true, running: make(map[string][]*pbs.JobAssignment)}
	g.running["testserver"] = []*pbs.JobAssignment{&pbs.JobAssignment{Job: &pbs.Job{Name: "testjob"}}}
	s.getter = g

	s.buildWorld()
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}
