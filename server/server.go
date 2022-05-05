package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"

	pb "github.com/sgrade/bgp-ping/bgpping"
	"github.com/sgrade/bgp-ping/common/ping"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type bgpPingServer struct {
	pb.UnimplementedBgpPingServer

	peerIpAddress string
}

func newBgpPingServer() *bgpPingServer {
	s := &bgpPingServer{peerIpAddress: ""}

	return s
}

var (
	// showVersion bool
	// version     string
	// gitCommit   string
	counter int // ping counter
	// timeout  string
	interval string // interval in time.Duration format, e.g. "1s"
	sigs     chan os.Signal
)

func (s *bgpPingServer) PingBgpPeer(ctx context.Context, in *pb.BgpPingRequest) (*pb.BgpPingResponse, error) {

	interval = "1s"
	counter = 4

	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var timeoutDuration time.Duration
	var intervalDuration time.Duration

	if res, err := strconv.Atoi(interval); err == nil {
		intervalDuration = time.Duration(res) * time.Millisecond
	} else {
		intervalDuration, err = time.ParseDuration(interval)
		if err != nil {
			fmt.Println("parse interval failed:\n", err)
			// return
			// TODO: handle an error properly
		}
	}

	var protocol ping.Protocol

	host := "10.0.12.11"
	port := 179

	parseHost, _ := ping.FormatIP(host)
	target := ping.Target{
		Timeout:  timeoutDuration,
		Interval: intervalDuration,
		Host:     parseHost,
		Port:     port,
		Counter:  counter,
		Protocol: protocol,
	}

	pinger := ping.NewTCPing()

	pinger.SetTarget(&target)
	pingerDone := pinger.Start()
	select {
	case <-pingerDone:
		break
	case <-sigs:
		break
	}

	fmt.Println(pinger.Result())

	return &pb.BgpPingResponse{
		ProbesSent:       int32(pinger.Result().Counter),
		ProbesSuccessful: int32(pinger.Result().SuccessCounter),
		ProbesFailed:     int32(pinger.Result().Failed()),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBgpPingServer(grpcServer, newBgpPingServer())
	grpcServer.Serve(lis)
}
