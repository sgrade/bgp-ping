package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sgrade/bgp-ping/bgpping"
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
