package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
)

const (
	intentWait = 1000
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	config  *pb.Config
	serving bool
}

type mainChecker struct{}

func getIP(servertype, servername string) (string, int) {
	conn, _ := grpc.Dial("192.168.86.34:50055", grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	r, err := registry.ListAllServices(context.Background(), &pbd.Empty{})
	if err != nil {
		return "", -1
	}
	for _, s := range r.Services {
		if s.Name == servertype && s.Identifier == servername {
			return s.Ip, int(s.Port)
		}
	}

	return "", -1
}

func (t *mainChecker) assess(server string) *pbs.JobList {
	list := &pbs.JobList{}

	ip, port := getIP("gobuildslave", server)
	conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	defer conn.Close()

	slave := pbs.NewGoBuildSlaveClient(conn)
	r, err := slave.List(context.Background(), &pbs.Empty{})
	if err == nil {
		return r
	}

	return list
}

func runJob(job *pbs.JobSpec) {
	ip, port := getIP("gobuildslave", job.Server)
	conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	defer conn.Close()

	slave := pbs.NewGoBuildSlaveClient(conn)
	slave.Run(context.Background(), job)
}

func (t *mainChecker) discover() *pbd.ServiceList {
	ret := &pbd.ServiceList{}

	conn, _ := grpc.Dial("192.168.86.34:50055", grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	r, err := registry.ListAllServices(context.Background(), &pbd.Empty{})
	if err == nil {
		for _, s := range r.Services {
			ret.Services = append(ret.Services, s)
		}
	}

	return ret
}

// DoRegister Registers this server
func (s Server) DoRegister(server *grpc.Server) {
	// Do nothing
}

// ReportHealth determines if the server is healthy
func (s Server) ReportHealth() bool {
	return true
}

func getConfig(c checker) *pb.Config {
	list := getFleetStatus(c)
	config := &pb.Config{}

	for _, jlist := range list {
		for _, job := range jlist.Details {
			found := false
			for _, ij := range config.Intents {
				if job.Spec.Name == ij.Spec.Name {
					ij.Masters++
					found = true
				}
			}

			if !found {
				config.Intents = append(config.Intents, &pb.Intent{Spec: &pbs.JobSpec{Name: job.Spec.Name}, Masters: 1})
			}
		}
	}

	return config
}

// MatchIntent tries to match the intent with the state of production
func (s Server) MatchIntent() {
	//for s.serving {
	time.Sleep(intentWait)

	state := getConfig(&mainChecker{})
	diff := configDiff(s.config, state)
	joblist := runJobs(diff)
	for _, job := range joblist {
		runJob(job)
	}
	//}
}

func main() {
	config, err := loadConfig("config.pb")
	if err != nil {
		log.Fatalf("Fatal loading of config: %v", err)
	}

	var sync = flag.Bool("once", true, "One pass intent match")
	s := Server{&goserver.GoServer{}, config, true}
	if *sync {
		s.MatchIntent()
	} else {
		s.Register = s
		s.PrepServer()
		s.RegisterServer("gobuildmaster", false)
		s.Serve()
	}
}
