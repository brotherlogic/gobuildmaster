package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
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
	LastIntent time.Time
	LastMaster time.Time
	worldMutex *sync.Mutex
	world      map[string]map[string]struct{}
	getter     getter
	mapString  string
}

type prodGetter struct{}

func (g *prodGetter) getJobs(server *pbd.RegistryEntry) (*pbs.JobList, error) {
	list := &pbs.JobList{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.Dial(server.GetIp()+":"+strconv.Itoa(int(server.GetPort())), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return list, err
	}

	slave := pbs.NewGoBuildSlaveClient(conn)
	r, err := slave.List(ctx, &pbs.Empty{}, grpc.FailFast(false))
	return r, err
}

func (s *Server) checkerThread(i *pb.Intent) {
	for true {
		time.Sleep(time.Minute)

		if len(s.world[i.GetSpec().GetName()]) != int(i.Count) {
			s.Log(fmt.Sprintf("MISMATCH: %v, %v", i, s.world[i.GetSpec().GetName()]))
			if len(s.world[i.GetSpec().GetName()]) < int(i.Count) {
				s.Log(fmt.Sprintf("UP : %v", i.GetSpec().GetName()))
			}
		}
	}
}

func (g *prodGetter) getSlaves() (*pbd.ServiceList, error) {
	ret := &pbd.ServiceList{}

	conn, err := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	if err != nil {
		return ret, err
	}
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.Empty{}, grpc.FailFast(false))
	if err != nil {
		return ret, err
	}

	for _, s := range r.Services {
		if s.GetName() == "gobuildslave" {
			ret.Services = append(ret.Services, s)
		}
	}

	return ret, nil
}

type mainChecker struct {
	prev   []string
	logger func(string)
}

func getIP(servertype, servername string) (string, int) {
	conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
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
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return list, conf
	}

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

func (t *mainChecker) master(entry *pbd.RegistryEntry, master bool) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, _ := grpc.Dial(entry.GetIp()+":"+strconv.Itoa(int(entry.GetPort())), grpc.WithInsecure())
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	_, err := server.Mote(ctx, &pbg.MoteRequest{Master: master}, grpc.FailFast(false))
	if err != nil {
		t.logger(fmt.Sprintf("Master REJECT(%v): %v", entry, err))
	}

	return err == nil
}

func runJob(job *pbs.JobSpec, server string) {
	if server != "" {
		ip, port := getIP("gobuildslave", server)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
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
	return []*pbg.State{&pbg.State{Key: "last_intent", TimeValue: s.LastIntent.Unix()},
		&pbg.State{Key: "last_master", TimeValue: s.LastMaster.Unix()},
		&pbg.State{Key: "world", Text: fmt.Sprintf("%v", s.world)},
		&pbg.State{Key: "master", Text: s.mapString}}
}

//Compare compares current state to desired state
func (s Server) Compare(ctx context.Context, in *pb.Empty) (*pb.CompareResponse, error) {
	resp := &pb.CompareResponse{}
	list, _ := getFleetStatus(&mainChecker{logger: s.Log})
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
	checker := &mainChecker{logger: s.Log}
	for s.serving {
		time.Sleep(intentWait)
		s.LastIntent = time.Now()

		state := getConfig(checker)
		diff := configDiff(s.config, state)
		joblist := runJobs(diff)
		for _, job := range joblist {
			runJob(job, chooseServer(job, checker))
		}
	}
}

// SetMaster sets up the master settings
func (s *Server) SetMaster() {
	t := time.Now()
	checker := &mainChecker{logger: s.Log}
	s.LastMaster = time.Now()
	masterMap := make(map[string]string)

	fleet := checker.discover()
	matcher := make(map[string][]*pbd.RegistryEntry)
	hasMaster := make(map[string]int)
	for _, entry := range fleet.GetServices() {
		if !entry.GetIgnoresMaster() {
			if _, ok := matcher[entry.GetName()]; !ok {
				if entry.GetMaster() {
					hasMaster[entry.GetName()]++
					masterMap[entry.GetName()] = entry.GetIdentifier()
				}
				matcher[entry.GetName()] = []*pbd.RegistryEntry{entry}
			} else {
				if entry.GetMaster() {
					hasMaster[entry.GetName()] = 1
					matcher[entry.GetName()] = append(matcher[entry.GetName()], entry)
					masterMap[entry.GetName()] = entry.GetIdentifier()
				}
			}
		}
	}

	for key, entries := range matcher {
		if hasMaster[key] > 1 {
			hasMaster[key] = 1
			seen := false
			for _, entry := range entries {
				if seen && entry.GetMaster() {
					checker.master(entry, false)
				} else if entry.GetMaster() {
					seen = true
				}
			}
		}

		if hasMaster[key] == 0 {
			for _, entry := range entries {
				if checker.master(entry, true) {
					masterMap[entry.GetName()] = entry.GetIdentifier()
					entry.Master = true
					break
				}
			}
		}
	}
	s.mapString = fmt.Sprintf("%v", masterMap)
	s.LogFunction("SetMasterRun", t)
}

//Init builds up the server
func Init(config *pb.Config) *Server {
	s := &Server{
		&goserver.GoServer{},
		config,
		true,
		time.Now(),
		time.Now(),
		&sync.Mutex{},
		make(map[string]map[string]struct{}),
		&prodGetter{},
		"",
	}
	return s
}

func (s *Server) becomeMaster() {
	for true {
		time.Sleep(time.Second * 5)
		_, _, err := utils.Resolve("gobuildmaster")
		if err != nil {
			s.Registry.Master = true
		}
	}
}

func main() {
	config, err := loadConfig("config.pb")
	if err != nil {
		log.Fatalf("Fatal loading of config: %v", err)
	}

	s := Init(config)

	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	s.Register = s
	s.PrepServer()
	s.GoServer.Killme = false
	s.RegisterServer("gobuildmaster", false)
	s.RegisterServingTask(s.MatchIntent)
	s.RegisterServingTask(s.becomeMaster)
	s.RegisterRepeatingTask(s.SetMaster, time.Second)
	s.RegisterRepeatingTask(s.buildWorld, time.Minute)

	for i := 0; i < len(s.config.GetIntents()); i++ {
		go s.checkerThread(s.config.GetIntents()[i])
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
