package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	slaveMap          map[string][]string
	getter            getter
	mapString         string
	lastWorldRun      int64
	lastMasterSatisfy map[string]time.Time
	serverMap         map[string]time.Time
	lastSeen          map[string]time.Time
	timeChange        time.Duration
	registerAttempts  int64
	lastMasterRunTime time.Duration
	lastJob           string
	lastTrack         string
	accessPoints      map[string]time.Time
	accessPointsMutex *sync.Mutex
}

func (s *Server) alertOnMissingJob(ctx context.Context) error {
	for _, nin := range s.config.Nintents {
		_, _, err := utils.Resolve(nin.Job.Name, "gobuildmaster-alertonjob")
		if err != nil && !nin.GetNoMaster() {
			if _, ok := s.lastSeen[nin.Job.Name]; !ok {
				s.lastSeen[nin.Job.Name] = time.Now()
			}

			if time.Now().Sub(s.lastSeen[nin.Job.Name]) > time.Hour*2 {
				if nin.Job.Name == "githubcard" {
					panic(fmt.Errorf("Unable to locate githubcard"))
				}
				s.RaiseIssue(ctx, "Missing Job", fmt.Sprintf("%v is missing - last seen %v (%v)", nin.Job.Name, time.Now().Sub(s.lastSeen[nin.Job.Name]), err), false)
			}
		} else {
			s.lastSeen[nin.Job.Name] = time.Now()
		}
	}

	return nil
}

type prodGetter struct {
	dial func(entry *pbd.RegistryEntry) (*grpc.ClientConn, error)
}

func (g *prodGetter) getJobs(ctx context.Context, server *pbd.RegistryEntry) ([]*pbs.JobAssignment, error) {
	conn, err := g.dial(server)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	slave := pbs.NewBuildSlaveClient(conn)

	// Set a tighter rpc deadline for listing jobs.
	ctx, cancel := utils.ManualContext("getJobs", "gobuildmaster", time.Minute)
	defer cancel()
	r, err := slave.ListJobs(ctx, &pbs.ListRequest{})
	if err != nil {
		return nil, err
	}
	return r.Jobs, err
}

func (g *prodGetter) getConfig(ctx context.Context, server *pbd.RegistryEntry) ([]*pbs.Requirement, error) {
	conn, err := g.dial(server)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

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
		ctx, cancel := utils.BuildContext("gobuildmaster", "gobuildmaster")
		defer cancel()
		if s.Registry.Master && lastRun < s.lastWorldRun {
			if i.Redundancy == pb.Redundancy_GLOBAL {
				s.runJob(ctx, i.GetJob())
			}
			s.worldMutex.Lock()
			if i.Redundancy == pb.Redundancy_REDUNDANT {
				if len(s.world[i.GetJob().GetName()]) < 3 {
					s.runJob(ctx, i.GetJob())
				}
			}
			if len(s.world[i.GetJob().GetName()]) != int(i.Count) {
				if len(s.world[i.GetJob().GetName()]) < int(i.Count) {
					s.runJob(ctx, i.GetJob())
				}
			}
			s.worldMutex.Unlock()
			lastRun = time.Now().Unix()
		}
	}
}

func (g *prodGetter) getSlaves() (*pbd.ServiceList, error) {
	ret := &pbd.ServiceList{}

	conn, err := grpc.Dial("127.0.0.1:"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	if err != nil {
		return ret, err
	}
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceV2Client(conn)
	ctx, cancel := utils.ManualContext("getSlaves", "gobuildmaster", time.Minute)
	defer cancel()
	r, err := registry.Get(ctx, &pbd.GetRequest{Job: "gobuildslave"})
	if err != nil {
		return ret, err
	}

	for _, s := range r.GetServices() {
		if s.GetName() == "gobuildslave" {
			ret.Services = append(ret.Services, s)
		}
	}

	return ret, nil
}

type mainChecker struct {
	prev      []string
	logger    func(string)
	dial      func(server, host string) (*grpc.ClientConn, error)
	dialEntry func(*pbd.RegistryEntry) (*grpc.ClientConn, error)
}

func getIP(servertype, servername string) (string, int) {
	conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.ListRequest{})
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

func (t *mainChecker) assess(ctx context.Context, server string) (*pbs.JobList, *pbs.Config) {
	conn, err := t.dial("gobuildslave", server)
	if err != nil {
		return nil, nil
	}
	defer conn.Close()

	slave := pbs.NewGoBuildSlaveClient(conn)
	r, err := slave.List(ctx, &pbs.Empty{})
	if err != nil {
		return nil, nil
	}

	r2, err := slave.GetConfig(ctx, &pbs.Empty{})
	if err != nil {
		return nil, nil
	}

	return r, r2
}

func (t *mainChecker) master(entry *pbd.RegistryEntry, master bool) (bool, error) {
	conn, err := t.dialEntry(entry)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	ctx, cancel := utils.ManualContext("mastermaster", "mastermaster", time.Minute*5)
	defer cancel()

	server := pbg.NewGoserverServiceClient(conn)
	_, err = server.Mote(ctx, &pbg.MoteRequest{Master: master})

	return err == nil, err
}

func (s *Server) runJob(ctx context.Context, job *pbs.Job) {
	server := s.selectServer(ctx, job, s.getter)
	if server != "" {
		conn, err := s.DialServer("gobuildslave", server)
		if err == nil {
			defer conn.Close()

			slave := pbs.NewBuildSlaveClient(conn)
			slave.RunJob(ctx, &pbs.RunRequest{Job: job})
		}
	}
}

func (t *mainChecker) discover() *pbd.ServiceList {
	ret := &pbd.ServiceList{}

	conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := registry.ListAllServices(ctx, &pbd.ListRequest{})
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
	s.worldMutex.Lock()
	defer s.worldMutex.Unlock()
	aps := make([]string, 0)
	s.accessPointsMutex.Lock()
	defer s.accessPointsMutex.Unlock()
	for ap := range s.accessPoints {
		aps = append(aps, ap)
	}
	sort.Strings(aps)
	return []*pbg.State{
		&pbg.State{Key: "num_jobs", Value: int64(len(s.config.Nintents))},
		&pbg.State{Key: "access_points", Text: fmt.Sprintf("%v", aps)},
		&pbg.State{Key: "last_intent", TimeValue: s.LastIntent.Unix()},
		&pbg.State{Key: "last_job", Text: s.lastJob},
		&pbg.State{Key: "last_track", Text: s.lastTrack},
		&pbg.State{Key: "last_state", TimeValue: s.LastMaster.Unix()},
		&pbg.State{Key: "last_master", TimeValue: s.LastMaster.Unix()},
		&pbg.State{Key: "last_master_time", TimeDuration: s.lastMasterRunTime.Nanoseconds()},
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
	list, _ := getFleetStatus(ctx, &mainChecker{logger: s.Log, dial: s.DialServer, dialEntry: s.DoDial})
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

func getConfig(ctx context.Context, c checker) *pb.Config {
	list, _ := getFleetStatus(ctx, c)
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
func (s *Server) SetMaster(ctx context.Context) error {
	checker := &mainChecker{logger: s.Log, dial: s.DialServer, dialEntry: s.DoDial}
	s.LastMaster = time.Now()
	masterMap := make(map[string]string)

	fleet := checker.discover()
	matcher := make(map[string][]*pbd.RegistryEntry)
	hasMaster := make(map[string]int)
	s.lastTrack = "Building Mapping"
	for _, entry := range fleet.GetServices() {
		if !entry.GetIgnoresMaster() && entry.GetVersion() == pbd.RegistryEntry_V1 {
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

	s.lastTrack = "Processing Mapping"
	for key, entries := range matcher {
		s.lastJob = key
		seen := hasMaster[key] == 1
		if hasMaster[key] > 1 {
			hasMaster[key] = 1
			for _, entry := range entries {
				if seen && entry.GetMaster() {
					s.lastTrack = fmt.Sprintf("%v master for %v", entry.Identifier, entry.Name)
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
				s.lastTrack = fmt.Sprintf("%v master for %v", entry.Identifier, entry.Name)
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
	s.lastMasterRunTime = time.Now().Sub(s.LastMaster)

	return nil
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
		make(map[string][]string),
		&prodGetter{},
		"",
		0,
		make(map[string]time.Time),
		make(map[string]time.Time),
		make(map[string]time.Time),
		time.Hour, // time.Change
		int64(0),
		0,
		"",
		"",
		make(map[string]time.Time),
		&sync.Mutex{},
	}
	s.getter = &prodGetter{s.DoDial}

	return s
}

func (s *Server) becomeMaster(ctx context.Context) error {
	for true {
		time.Sleep(time.Second * 5)
		if !s.Registry.Master {
			_, _, err := utils.Resolve("gobuildmaster", "gobuildmaster-become")
			if err != nil {
				s.registerAttempts++
				s.Registry.Master = true
			}
		}
	}

	return nil
}

func (s *Server) raiseIssue(ctx context.Context) error {
	for key, val := range s.lastMasterSatisfy {
		if time.Now().Sub(val) > time.Hour {
			conn, err := s.DialMaster("githubcard")
			if err == nil {
				defer conn.Close()
				client := pbgh.NewGithubClient(conn)
				client.AddIssue(ctx, &pbgh.Issue{Service: key, Title: fmt.Sprintf("No Master Found - %v", key), Body: ""})
			}

		}
	}

	return nil
}

func (s *Server) registerJobs(ctx context.Context) error {
	conn, err := s.DialMaster("githubcard")
	if err != nil {
		return nil
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	for _, job := range s.config.Nintents {
		_, err := client.RegisterJob(ctx, &pbgh.RegisterRequest{Job: job.GetJob().GetName()})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	config, err := loadConfig()
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

	err = s.RegisterServerV2("gobuildmaster", false, false)
	if err != nil {
		if c := status.Convert(err); c.Code() == codes.FailedPrecondition {
			// this is expected if disc is not ready
			return
		}
		log.Fatalf("Unable to register: %v", err)
	}
	s.RegisterRepeatingTask(s.buildWorld, "build_world", time.Minute)
	s.RegisterServingTask(s.becomeMaster, "become_master")

	// Don't trace out master requests - they can take a while
	s.RegisterRepeatingTask(s.alertOnMissingJob, "alert_on_missing_job", time.Minute*5)
	s.RegisterRepeatingTask(s.registerJobs, "register_jobs", time.Minute)

	for i := 0; i < len(s.config.GetNintents()); i++ {
		go s.checkerThread(s.config.GetNintents()[i])
	}

	err = s.Serve()
	if err != nil {

		log.Fatalf("Serve error: %v", err)
	}
}
