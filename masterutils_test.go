package main

import (
	pb "github.com/brotherlogic/gobuildmaster/proto"
)

func InitTestServer() *Server {
	s := Init(&pb.Config{})
	s.PrepServer()
	s.SkipLog = true
	s.SkipIssue = true
	return s
}
