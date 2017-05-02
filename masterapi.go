package main

import (
	"log"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/gobuildmaster/proto"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	config *pb.Config
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
	config, err := loadConfig("config.pb")
	if err != nil {
		log.Fatalf("Fatal loading of config: %v", err)
	}
	s := Server{&goserver.GoServer{}, config}
	s.Register = s
	s.PrepServer()
	s.RegisterServer("gobuildmaster", false)
	s.Serve()
}
