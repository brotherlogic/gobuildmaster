package main

import (
	"errors"
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbgs "github.com/brotherlogic/gobuildslave/proto"
)

func InitTestServer() *Server {
	s := Init(&pb.Config{})
	s.SkipLog = true
	return s
}

type testGetter struct {
	failGetSlaves bool
	failGetJobs   bool
}

func (g *testGetter) getSlaves() (*pbd.ServiceList, error) {
	if g.failGetSlaves {
		return &pbd.ServiceList{}, errors.New("Built to fail")
	}
	return &pbd.ServiceList{Services: []*pbd.RegistryEntry{
		&pbd.RegistryEntry{Name: "gobuildslave", Identifier: "testserver"},
	}}, nil
}

func (g *testGetter) getJobs(r *pbd.RegistryEntry) ([]*pbgs.JobAssignment, error) {
	if g.failGetJobs {
		return []*pbgs.JobAssignment{}, errors.New("Built to fail")
	}
	return []*pbgs.JobAssignment{&pbgs.JobAssignment{Server: "testserver", Job: &pbgs.Job{Name: "testjob"}}}, nil

}

func TestBuildWorld(t *testing.T) {
	s := InitTestServer()
	s.getter = &testGetter{}

	s.buildWorld()
	if _, ok := s.world["testjob"]["testserver"]; !ok {
		t.Fatalf("World has not been built correctly: %v", s.world)
	}
}

func TestBuildWorldFailSlaves(t *testing.T) {
	s := InitTestServer()
	s.getter = &testGetter{failGetSlaves: true}

	s.buildWorld()
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}

func TestBuildWorldFailJobs(t *testing.T) {
	s := InitTestServer()
	s.getter = &testGetter{failGetJobs: true}

	s.buildWorld()
	if len(s.world) != 0 {
		t.Errorf("World has been built: %v", s.world)
	}
}
