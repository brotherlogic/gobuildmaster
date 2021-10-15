package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/gobuildmaster/proto"
)

func main() {
	ctx, cancel := utils.ManualContext("gobuildmaster-cli", time.Minute)
	defer cancel()

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: list run\n")
	} else {
		switch os.Args[1] {
		case "list":
			listFlags := flag.NewFlagSet("List", flag.ExitOnError)
			var host = listFlags.String("host", "", "Name of the queue")
			if err := listFlags.Parse(os.Args[2:]); err == nil {
				conn, err := utils.LFDialSpecificServer(ctx, "gobuildmaster", *host)
				if err != nil {
					log.Fatalf("Unable to dial: %v", err)
				}

				client := pb.NewGoBuildMasterClient(conn)
				decisions, err := client.GetDecisions(ctx, &pb.GetDecisionsRequest{})
				if err != nil {
					log.Fatalf("Request failed: %v", err)
				}

				for _, decision := range decisions.GetDecisions() {
					fmt.Printf("%v -> %v / %v\n", decision.GetJobName(), decision.GetRunning(), decision.GetReason())
				}
			}
		}
	}
}
