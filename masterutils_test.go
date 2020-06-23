package main

import (
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
	"golang.org/x/net/context"
)

func InitTestServer() *Server {
	s := Init(&pb.Config{})
	s.PrepServer()
	s.SkipLog = true
	s.SkipIssue = true
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment)}
	tg.running["server1"] = []*pbs.JobAssignment{}
	s.Registry = &pbd.RegistryEntry{Identifier: "server1"}
	s.getter = tg
	s.testing = true
	return s
}

func TestFailGetJobs(t *testing.T) {
	s := InitTestServer()
	tg := &testGetter{running: make(map[string][]*pbs.JobAssignment),
		failGetJobs: true}
	s.getter = tg

	_, err := s.updateWorld(context.Background(), &pbd.RegistryEntry{})
	if err == nil {
		t.Errorf("Did not fail")
	}
}

func TestAdjust(t *testing.T) {
	s := InitTestServer()
	s.config.Nintents = append(s.config.Nintents, &pb.NIntent{Job: &pbs.Job{Name: "blah"}, Redundancy: pb.Redundancy_GLOBAL})
	err := s.adjustWorld(context.Background())
	if err != nil {
		t.Errorf("Bad adjust on empty config: %v", err)
	}
}
