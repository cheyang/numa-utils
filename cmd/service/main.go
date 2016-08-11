package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
	numa "github.com/cheyang/numa-utils/proto"
	"github.com/cheyang/numa-utils/service"
	"google.golang.org/grpc"
)

var bind string

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)
	interrupts := make(chan os.Signal)
	signal.Notify(interrupts, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-interrupts
		log.Infof("Exit.")
		os.Exit(0)
	}()

	listen, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := service.Server{}
	s := grpc.NewServer()
	numa.RegisterNumaServer(s, server)
	log.Infof("listen on %v", listen)
	panic(s.Serve(listen))

}

func init() {
	flag.StringVar(&bind, "bind", ":30000", "Service bind address")
}
