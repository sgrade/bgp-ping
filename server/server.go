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
	port                          = flag.Int("port", 50051, "The BGP ping server port")
	stopAfterSendingCountRequests = flag.Int64("count", 3, "Stop after sending count pings")
)

type bgpPingServer struct {
	pb.UnimplementedBgpPingServer

	peerIpAddress                 string
	stopAfterSendingCountRequests int64
}

func newBgpPingServer() *bgpPingServer {
	s := &bgpPingServer{peerIpAddress: "", stopAfterSendingCountRequests: int64(*stopAfterSendingCountRequests)}

	return s
}

var (
	// showVersion bool
	// version     string
	// timeout  string
	interval string // interval in time.Duration format, e.g. "1s"
	sigs     chan os.Signal
)

func (s *bgpPingServer) PingBgpPeer(ctx context.Context, in *pb.BgpPingRequest) (*pb.BgpPingResponse, error) {

	interval = "1s"

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

	peerIpAddress := "10.0.12.11"
	peerTcpPort := 179

	parseHost, _ := ping.FormatIP(peerIpAddress)
	target := ping.Target{
		Timeout:  timeoutDuration,
		Interval: intervalDuration,
		Host:     parseHost,
		Port:     peerTcpPort,
		Counter:  s.stopAfterSendingCountRequests,
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
		ProbesSent:       int64(pinger.Result().Counter),
		ProbesSuccessful: int64(pinger.Result().SuccessCounter),
		ProbesFailed:     int64(pinger.Result().Failed()),
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

	// https://stackoverflow.com/questions/55797865/behavior-of-server-gracefulstop-in-golang
	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	// bind OS events to the signal channel
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	// run blocking call in a separate goroutine, report errors via channel
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	// terminate your environment gracefully before leaving main function
	defer func() {
		grpcServer.GracefulStop()
		// closeDbConnections()
	}()

	// block until either OS signal, or server fatal error
	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case s := <-stopChan:
		log.Printf("\nGot signal %v, attempting graceful shutdown", s)
	}
}
