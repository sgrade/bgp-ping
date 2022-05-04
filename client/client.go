package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sgrade/bgp-ping/bgpping"
)

const (
	defaultPeerIp = "10.0.12.11"
)

var (
	serverAddr    = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	peerIpAddress = flag.String("ip", defaultPeerIp, "The BGP peer IP address")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewBgpPingClient(conn)

	r, err := client.PingBgpPeer(context.Background(), &pb.BgpPingRequest{PeerIp: *peerIpAddress})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	fmt.Println(r.ProbesSent, r.ProbesSuccessful, r.ProbesFailed)
}
