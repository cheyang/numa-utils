package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	numa "github.com/cheyang/numa-utils/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	server string
	period time.Duration
)

func main() {
	flag.Parse()

	ctx := context.Background()
	conn, err := grpc.Dial(server, grpc.WithInsecure())

	if err != nil {
		fmt.Printf("error is %v", err)
		return
	}
	client := numa.NewNumaClient(conn)

	signals := make(chan os.Signal)

	//Exit on system signal -3 or -15
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signals:
			fmt.Println("exited successfully.")
			os.Exit(0)
		case <-time.After(period):
		}
		printInfo(client, ctx)
		printMetrics(client, ctx)

	}

	// client :=

}

func init() {
	flag.StringVar(&server, "server", "localhost:30000", "the server url")
	flag.DurationVar(&period, "period", 5*time.Second, "the period")
}

func printInfo(client numa.NumaClient, ctx context.Context) {

	resp, err := client.GetInfo(ctx, &numa.Empty{})

	if err != nil {
		fmt.Printf("error is %v", err)
	}

	nodes := resp.GetNodes()

	for _, node := range nodes {
		fmt.Printf("node %d cpus: %+v\n", node.Id, node.CpuSet)
		fmt.Printf("node %d size: %d MB\n", node.Id, node.Total)
		fmt.Printf("node %d free: %d MB\n", node.Id, node.Free)
	}

	fmt.Println("")
}

func printMetrics(client numa.NumaClient, ctx context.Context) {
	resp, err := client.GetMetrics(ctx, &numa.Empty{})

	if err != nil {
		fmt.Printf("error is %v", err)
	}

	fmt.Printf("node distances:\n")

	for _, distance := range resp.GetDistances() {

		fmt.Printf("node: %d and %d 's distance is % 3d \n", distance.Start, distance.End, distance.Length)

	}

	fmt.Println("=============================")
}
