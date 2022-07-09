package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	testing           bool
	decisions         map[string]string
	claimed           string
}

func (s *Server) alertOnMissingJob(ctx context.Context) error {
	for _, nin := range s.config.Nintents {
		_, err := utils.ResolveV3(nin.Job.Name)
		if err != nil && !nin.GetNoMaster() {
			if _, ok := s.lastSeen[nin.Job.Name]; !ok {
				s.lastSeen[nin.Job.Name] = time.Now()
			}

			if time.Now().Sub(s.lastSeen[nin.Job.Name]) > time.Hour*2 {
				if nin.Job.Name == "githubcard" {
					fmt.Printf("Unable to locate githubcard\n")
					os.Exit(1)
				}

				// Discovery does not show up in discovery
				if nin.Job.Name != "discovery" {
					s.RaiseIssue("Missing Job", fmt.Sprintf("%v is missing - last seen %v (%v)", nin.Job.Name, time.Now().Sub(s.lastSeen[nin.Job.Name]), err))
				}
			}
		} else {
			s.lastSeen[nin.Job.Name] = time.Now()
		}
	}

	return nil
}

type prodGetter struct {
	dial func(entry string) (*grpc.ClientConn, error)
}

func (g *prodGetter) getJobs(ctx context.Context, server *pbd.RegistryEntry) ([]*pbs.JobAssignment, error) {
	conn, err := g.dial(fmt.Sprintf("%v:%v", server.GetIdentifier(), server.GetPort()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	slave := pbs.NewBuildSlaveClient(conn)

	// Set a tighter rpc deadline for listing jobs.
	ctx, cancel := utils.ManualContext("getJobs", time.Minute)
	defer cancel()
	r, err := slave.ListJobs(ctx, &pbs.ListRequest{})
	if err != nil {
		return nil, err
	}
	return r.Jobs, err
}

func (g *prodGetter) getConfig(ctx context.Context, server *pbd.RegistryEntry) ([]*pbs.Requirement, error) {
	conn, err := g.dial(fmt.Sprintf("%v:%v", server.GetIdentifier(), server.GetPort()))
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

func (g *prodGetter) getSlaves(ctx context.Context) (*pbd.ServiceList, error) {
	ret := &pbd.ServiceList{}

	conn, err := g.dial("127.0.0.1:" + strconv.Itoa(utils.RegistryPort))
	if err != nil {
		return ret, err
	}
	defer conn.Close()

	registry := pbd.NewDiscoveryServiceV2Client(conn)
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

func (s *Server) runJob(ctx context.Context, job *pbs.Job, localSlave *pbd.RegistryEntry, bits int) error {
	if s.testing {
		return nil
	}
	conn, err := s.FDial(fmt.Sprintf("%v:%v", localSlave.GetIdentifier(), localSlave.GetPort()))
	if err == nil {
		defer conn.Close()

		slave := pbs.NewBuildSlaveClient(conn)
		s.Log(fmt.Sprintf("Attempting to run %v", job))
		_, err = slave.RunJob(ctx, &pbs.RunRequest{Job: job, Bits: int32(bits)})
	}
	return err
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

//GetState gets the state of the server
func (s Server) GetState() []*pbg.State {
	return []*pbg.State{}
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

func (s *Server) GetDecisions(ctx context.Context, _ *pb.GetDecisionsRequest) (*pb.GetDecisionsResponse, error) {
	resp := &pb.GetDecisionsResponse{}
	for job, dec := range s.decisions {
		resp.Decisions = append(resp.Decisions, &pb.Decision{
			JobName: job,
			Running: len(dec) == 0,
			Reason:  dec,
		})
	}
	return resp, nil
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
		false,
		make(map[string]string),
		"",
	}
	s.getter = &prodGetter{s.FDial}

	return s
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

func (s *Server) registerJob(ctx context.Context, int *pb.NIntent) error {
	conn, err := s.FDialServer(ctx, "githubcard")
	if err != nil {
		return nil
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	_, err = client.RegisterJob(ctx, &pbgh.RegisterRequest{Job: int.GetJob().GetName()})

	return err
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
	s.PrepServer("gobuildmaster")

	err = s.RegisterServerV2(false)
	if err != nil {
		if c := status.Convert(err); c.Code() == codes.FailedPrecondition || c.Code() == codes.Unavailable {
			// this is expected if disc is not ready
			return
		}
		log.Fatalf("Unable to register: %v", err)
	}

	//We need to register ourselves
	go func() {
		for true {
			time.Sleep(time.Minute)

			ctx, cancel := utils.ManualContext("gbm-register", time.Minute)
			defer cancel()

			conn, err := s.FDialServer(ctx, "githubcard")
			if err != nil {
				s.CtxLog(ctx, fmt.Sprintf("Cannot dial: %v", err))
				continue
			}

			client := pbgh.NewGithubClient(conn)
			_, err = client.RegisterJob(ctx, &pbgh.RegisterRequest{Job: "gobuildmaster"})
			if err != nil {
				s.CtxLog(ctx, fmt.Sprintf("Unable to register: %v", err))
			}
			break
		}
	}()

	go func() {
		for !s.LameDuck {
			t1 := time.Now()
			ctx, cancel := utils.ManualContext("gobuildmaster", time.Minute*10)
			s.CtxLog(ctx, "Beginning rebuild run")
			err = s.adjustWorld(ctx)
			if err != nil {
				s.CtxLog(ctx, fmt.Sprintf("Cannot run jobs: %v", err))
			}
			cancel()
			s.CtxLog(ctx, fmt.Sprintf("Took %v to adjust the world", time.Since(t1)))
			rebuildTime.Set(float64(time.Since(t1).Seconds()))
			time.Sleep(time.Minute * 10)
		}
	}()
	err = s.Serve()
	if err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
