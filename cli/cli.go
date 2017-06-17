package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/gobuildmaster/proto"
)

func findServer(name string) (string, int) {
	conn, _ := grpc.Dial("192.168.86.64:50055", grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	rs, _ := registry.ListAllServices(context.Background(), &pbdi.Empty{})

	for _, r := range rs.Services {
		if r.Name == name {
			log.Printf("%v -> %v", name, r)
			return r.Ip, int(r.Port)
		}
	}

	return "", -1
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: list run\n")
	} else {
		switch os.Args[1] {
		case "list":
			host, port := findServer("gobuildmaster")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			registry := pb.NewGoBuildMasterClient(conn)
			res, err := registry.Compare(context.Background(), &pb.Empty{})
			if err != nil {
				log.Printf("Error building job: %v", err)
			}
			log.Printf("Actual: %v", res.Current)
			log.Printf("Desire: %v", res.Desired)
		}
	}
}
