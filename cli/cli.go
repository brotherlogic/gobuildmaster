package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
)

func findServer(name string) (string, int) {
	conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	rs, _ := registry.ListAllServices(context.Background(), &pbdi.ListRequest{})

	for _, r := range rs.Services.Services {
		if r.Name == name {
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

			conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Cannot dial master: %v", err)
			}
			defer conn.Close()

		}
	}
}
