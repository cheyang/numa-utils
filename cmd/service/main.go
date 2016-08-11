package main

import (
	"flag"
	"net"
	"os"
	"os/signal"

	log "github.com/Sirupsen/logrus"
	numa "github.com/cheyang/numa-utils/proto"
	"github.com/cheyang/numa-utils/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var bind string

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	interrupts := make(chan os.Signal)
	signal.Notify(interrupts, os.Interrupt)
	go func() {
		<-interrupts
		cancel()
		os.Exit(0)
	}()

	listen, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := service.Server{}
	s := grpc.NewServer()
	numa.RegisterNumaServer(s, server)
	panic(s.Serve(listen))

}

func init() {
	flag.StringVar(&bind, "bind", ":20000", "Service bind address")
}
