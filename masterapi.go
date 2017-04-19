package main

import (
	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
}

// DoRegister Registers this server
func (s Server) DoRegister(server *grpc.Server) {
	// Do nothing
}

// ReportHealth determines if the server is healthy
func (s Server) ReportHealth() bool {
	return true
}

func main() {
	s := Server{&goserver.GoServer{}}
	s.Register = s
	s.PrepServer()
	s.RegisterServer("gobuildmaster", false)
	s.Serve()
}
