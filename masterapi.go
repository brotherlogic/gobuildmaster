package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
	pbs "github.com/brotherlogic/gobuildslave/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
)

const (
	intentWait = time.Second
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	config     *pb.Config
	serving    bool
	lastIntent time.Time
}

type mainChecker struct {
	prev []string
}

func getIP(servertype, servername string) (string, int) {
	conn, err := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	log.Printf("Error? %v", err)
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.Empty{}, grpc.FailFast(false))
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

func (t *mainChecker) getprev() []string {
	return t.prev
}
func (t *mainChecker) setprev(v []string) {
	t.prev = v
}

func (t *mainChecker) assess(server string) (*pbs.JobList, *pbs.Config) {
	list := &pbs.JobList{}
	conf := &pbs.Config{}

	ip, port := getIP("gobuildslave", server)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return list, conf
	}
	defer conn.Close()

	slave := pbs.NewGoBuildSlaveClient(conn)
	r, err := slave.List(ctx, &pbs.Empty{}, grpc.FailFast(false))
	if err != nil {
		return list, conf
	}

	r2, err := slave.GetConfig(ctx, &pbs.Empty{}, grpc.FailFast(false))
	if err != nil {
		return list, conf
	}

	return r, r2
}

func (t *mainChecker) master(entry *pbd.RegistryEntry) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, _ := grpc.DialContext(ctx, entry.GetIp()+":"+strconv.Itoa(int(entry.GetPort())), grpc.WithInsecure())
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	log.Printf("SETTING MASTER: %v", entry)
	_, err := server.Mote(ctx, &pbg.MoteRequest{Master: entry.GetMaster()}, grpc.FailFast(false))
	if err != nil {
		log.Printf("RESPONSE: %v", err)
	}
}

func runJob(job *pbs.JobSpec, server string) {
	log.Printf("Run %v on %v", job.Name, server)
	if server != "" {
		ip, port := getIP("gobuildslave", server)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, _ := grpc.DialContext(ctx, ip+":"+strconv.Itoa(port), grpc.WithInsecure())
		defer conn.Close()

		slave := pbs.NewGoBuildSlaveClient(conn)
		job.Server = server
		slave.Run(ctx, job, grpc.FailFast(false))
	}
}

func (t *mainChecker) discover() *pbd.ServiceList {
	ret := &pbd.ServiceList{}

	conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.Empty{}, grpc.FailFast(false))
	if err == nil {
		for _, s := range r.Services {
			ret.Services = append(ret.Services, s)
		}
	}

	return ret
}

// DoRegister Registers this server
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterGoBuildMasterServer(server, &s)
}

// ReportHealth determines if the server is healthy
func (s Server) ReportHealth() bool {
	return true
}

// Mote promotes/demotes this server
func (s Server) Mote(master bool) error {
	return nil
}

//GetState gets the state of the server
func (s Server) GetState() []*pbg.State {
	return []*pbg.State{&pbg.State{Key: "last_intent", TimeValue: s.lastIntent.Unix()}}
}

//Compare compares current state to desired state
func (s Server) Compare(ctx context.Context, in *pb.Empty) (*pb.CompareResponse, error) {
	resp := &pb.CompareResponse{}
	list, _ := getFleetStatus(&mainChecker{})
	cc := &pb.Config{}
	for _, jlist := range list {
		for _, job := range jlist.GetDetails() {
			cc.Intents = append(cc.Intents, &pb.Intent{Spec: job.GetSpec()})
		}
	}
	resp.Current = cc
	resp.Desired = s.config

	return resp, nil
}

func getConfig(c checker) *pb.Config {
	list, _ := getFleetStatus(c)
	config := &pb.Config{}

	for _, jlist := range list {
		for _, job := range jlist.Details {
			found := false
			for _, ij := range config.Intents {
				if job.Spec.Name == ij.Spec.Name {
					ij.Count++
					found = true
				}
			}

			if !found {
				config.Intents = append(config.Intents, &pb.Intent{Spec: &pbs.JobSpec{Name: job.Spec.Name}, Count: 1})
			}
		}
	}

	return config
}

// MatchIntent tries to match the intent with the state of production
func (s *Server) MatchIntent() {
	checker := &mainChecker{}
	for s.serving {
		time.Sleep(intentWait)
		s.lastIntent = time.Now()

		state := getConfig(checker)
		diff := configDiff(s.config, state)
		joblist := runJobs(diff)
		for _, job := range joblist {
			runJob(job, chooseServer(job, checker))
		}
	}
}

// SetMaster sets up the master settings
func (s Server) SetMaster() {
	checker := &mainChecker{}
	for s.serving {
		time.Sleep(intentWait)

		fleet := checker.discover()
		matcher := make(map[string]*pbd.RegistryEntry)
		for _, entry := range fleet.GetServices() {
			if _, ok := matcher[entry.GetName()]; !ok {
				matcher[entry.GetName()] = entry
			} else {
				if entry.GetMaster() {
					matcher[entry.GetName()] = entry
				}
			}
		}

		for _, entry := range matcher {
			if !entry.GetMaster() {
				entry.Master = true
				checker.master(entry)
			}
		}
	}
}

func main() {
	config, err := loadConfig("config.pb")
	if err != nil {
		log.Fatalf("Fatal loading of config: %v", err)
	}

	var sync = flag.Bool("once", false, "One pass intent match")
	s := Server{&goserver.GoServer{}, config, true, time.Now()}

	var quiet = flag.Bool("quiet", true, "Show all output")
	flag.Parse()

	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	if *sync {
		s.MatchIntent()
	} else {
		s.Register = s
		s.PrepServer()
		s.GoServer.Killme = false
		s.RegisterServer("gobuildmaster", false)
		s.RegisterServingTask(s.MatchIntent)
		s.RegisterServingTask(s.SetMaster)
		s.Serve()
	}
}
