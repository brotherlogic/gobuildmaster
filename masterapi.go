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
	pbgh "github.com/brotherlogic/githubcard/proto"
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
	config            *pb.Config
	serving           bool
	LastIntent        time.Time
	LastMaster        time.Time
	worldMutex        *sync.Mutex
	world             map[string]map[string]struct{}
	getter            getter
	mapString         string
	lastWorldRun      int64
	lastMasterSatisfy map[string]time.Time
	serverMap         map[string]time.Time
	lastSeen          map[string]time.Time
	timeChange        time.Duration
	registerAttempts  int64
}

func (s *Server) alertOnMissingJob(ctx context.Context) {
	for _, nin := range s.config.Nintents {
		_, _, err := utils.Resolve(nin.Job.Name)
		if err != nil {
			if _, ok := s.lastSeen[nin.Job.Name]; !ok {
				s.lastSeen[nin.Job.Name] = time.Now()
			}

			if time.Now().Sub(s.lastSeen[nin.Job.Name]) > time.Hour {
				if nin.Job.Name == "githubcard" {
					panic(fmt.Errorf("Unable to locate githubcard"))
				}
				s.RaiseIssue(ctx, "Missing Job", fmt.Sprintf("%v is missing - last seen %v (%v)", nin.Job.Name, time.Now().Sub(s.lastSeen[nin.Job.Name]), err), false)
			}
		} else {
			s.lastSeen[nin.Job.Name] = time.Now()
		}
	}
}

type prodGetter struct{}

func (g *prodGetter) getJobs(server *pbd.RegistryEntry) ([]*pbs.JobAssignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.Dial(server.GetIp()+":"+strconv.Itoa(int(server.GetPort())), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	slave := pbs.NewBuildSlaveClient(conn)
	r, err := slave.ListJobs(ctx, &pbs.ListRequest{})
	if err != nil {
		return nil, err
	}
	return r.Jobs, err
}

func (g *prodGetter) getConfig(server *pbd.RegistryEntry) ([]*pbs.Requirement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.Dial(server.GetIp()+":"+strconv.Itoa(int(server.GetPort())), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	slave := pbs.NewBuildSlaveClient(conn)
	r, err := slave.SlaveConfig(ctx, &pbs.ConfigRequest{})
	if err != nil {
		return nil, err
	}
	return r.Config.Requirements, err
}

func (s *Server) checkerThread(i *pb.NIntent) {
	lastRun := int64(10)
	for true {
		time.Sleep(time.Minute)
		if s.Registry.Master && lastRun < s.lastWorldRun {
			if i.Redundancy == pb.Redundancy_GLOBAL {
				s.runJob(i.GetJob())
			}
			s.worldMutex.Lock()
			if i.Redundancy == pb.Redundancy_REDUNDANT {
				if len(s.world[i.GetJob().GetName()]) < 3 {
					s.runJob(i.GetJob())
				}
			}
			if len(s.world[i.GetJob().GetName()]) != int(i.Count) {
				if len(s.world[i.GetJob().GetName()]) < int(i.Count) {
					s.runJob(i.GetJob())
				}
			}
			s.worldMutex.Unlock()
			lastRun = time.Now().Unix()
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
	r, err := registry.ListAllServices(ctx, &pbd.ListRequest{}, grpc.FailFast(false))
	if err != nil {
		return ret, err
	}

	for _, s := range r.GetServices().Services {
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
	r, err := registry.ListAllServices(ctx, &pbd.ListRequest{}, grpc.FailFast(false))
	if err != nil {
		return "", -1
	}
	for _, s := range r.GetServices().Services {
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

func (t *mainChecker) master(entry *pbd.RegistryEntry, master bool) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, _ := grpc.Dial(entry.GetIp()+":"+strconv.Itoa(int(entry.GetPort())), grpc.WithInsecure())
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	_, err := server.Mote(ctx, &pbg.MoteRequest{Master: master}, grpc.FailFast(false))

	return err == nil, err
}

func (s *Server) runJob(job *pbs.Job) {
	server := s.selectServer(job, s.getter)
	if server != "" {
		s.Log(fmt.Sprintf("Running %v on %v", job.Name, server))
		ip, port := getIP("gobuildslave", server)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
		defer conn.Close()

		slave := pbs.NewBuildSlaveClient(conn)
		slave.RunJob(ctx, &pbs.RunRequest{Job: job}, grpc.FailFast(false))
	} else {
		s.Log(fmt.Sprintf("Unable to find server for %v", job.Name))
	}
}

func (t *mainChecker) discover() *pbd.ServiceList {
	ret := &pbd.ServiceList{}

	conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.ListRequest{}, grpc.FailFast(false))
	if err == nil {
		for _, s := range r.GetServices().Services {
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

// Shutdown does the shutdown
func (s Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s Server) Mote(ctx context.Context, master bool) error {
	return nil
}

//GetState gets the state of the server
func (s Server) GetState() []*pbg.State {
	return []*pbg.State{&pbg.State{Key: "last_intent", TimeValue: s.LastIntent.Unix()},
		&pbg.State{Key: "last_master", TimeValue: s.LastMaster.Unix()},
		&pbg.State{Key: "world", Text: fmt.Sprintf("%v", s.world)},
		&pbg.State{Key: "master", Text: s.mapString},
		&pbg.State{Key: "seen", Text: fmt.Sprintf("%v", s.lastMasterSatisfy)},
		&pbg.State{Key: "servers", Text: fmt.Sprintf("%v", s.serverMap)},
		&pbg.State{Key: "seen_map", Text: fmt.Sprintf("%v", s.lastSeen)},
		&pbg.State{Key: "register_attempts", Value: s.registerAttempts},
	}
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

// SetMaster sets up the master settings
func (s *Server) SetMaster(ctx context.Context) {
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
					masterMap[entry.GetName()] = entry.GetIdentifier()
				}
				matcher[entry.GetName()] = append(matcher[entry.GetName()], entry)
			}
		}
	}

	for key, entries := range matcher {
		seen := hasMaster[key] == 1
		if hasMaster[key] > 1 {
			hasMaster[key] = 1
			for _, entry := range entries {
				if seen && entry.GetMaster() {
					s.Log(fmt.Sprintf("Setting %v master for %v", entry.Identifier, entry.Name))
					checker.master(entry, false)
				} else if entry.GetMaster() {
					seen = true
				}
			}
		}

		if hasMaster[key] == 0 {
			if len(entries) == 0 {
				masterMap[key] = "NONE_AVAILABLE"
			}

			for _, entry := range entries {
				val, err := checker.master(entry, true)
				if val {
					masterMap[entry.GetName()] = entry.GetIdentifier()
					entry.Master = true
					seen = true
					break
				} else {
					masterMap[entry.GetName()] = fmt.Sprintf("%v", err)
				}
			}

		}

		_, ok := s.lastMasterSatisfy[key]
		if ok && seen {
			delete(s.lastMasterSatisfy, key)
		}
		if !ok && !seen {
			s.lastMasterSatisfy[key] = time.Now()
		}

	}
	s.mapString = fmt.Sprintf("%v", masterMap)
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
		0,
		make(map[string]time.Time),
		make(map[string]time.Time),
		make(map[string]time.Time),
		time.Hour,
		int64(0),
	}
	return s
}

func (s *Server) becomeMaster(ctx context.Context) {
	for true {
		time.Sleep(time.Second * 5)
		if !s.Registry.Master {
			_, _, err := utils.Resolve("gobuildmaster")
			if err != nil {
				s.registerAttempts++
				s.Registry.Master = true
			}
		}
	}
}

func (s *Server) raiseIssue(ctx context.Context) {
	for key, val := range s.lastMasterSatisfy {
		if time.Now().Sub(val) > time.Hour {
			ip, port := s.GetIP("githubcard")
			if port > 0 {
				conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
				if err == nil {
					defer conn.Close()
					client := pbgh.NewGithubClient(conn)
					client.AddIssue(ctx, &pbgh.Issue{Service: key, Title: fmt.Sprintf("No Master Found - %v", key), Body: ""}, grpc.FailFast(false))
				}
			}

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
	err = s.RegisterServer("gobuildmaster", false)
	if err != nil {
		log.Fatalf("Unable to register: %v", err)
	}
	s.RegisterRepeatingTask(s.buildWorld, "build_world", time.Minute)
	s.RegisterServingTask(s.becomeMaster, "become_master")
	s.RegisterRepeatingTask(s.SetMaster, "set_master", time.Minute)
	s.RegisterRepeatingTask(s.alertOnMissingJob, "alert_on_missing_job", time.Minute*5)

	for i := 0; i < len(s.config.GetNintents()); i++ {
		go s.checkerThread(s.config.GetNintents()[i])
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
