// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: bgpping/bgpping.proto

package bgpping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BgpPingClient is the client API for BgpPing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BgpPingClient interface {
	// A simple RPC where the client sends a request to the server
	// using the stub and waits for a response to come back,
	// just like a normal function call.
	PingBgpPeer(ctx context.Context, in *BgpPingRequest, opts ...grpc.CallOption) (*BgpPingResponse, error)
	// A server-side streaming RPC where the client sends a request to the server
	// and gets a stream to read a sequence of messages back.
	PingBgpPeerContinuously(ctx context.Context, in *BgpPingRequest, opts ...grpc.CallOption) (*BgpPingResponse, error)
}

type bgpPingClient struct {
	cc grpc.ClientConnInterface
}

func NewBgpPingClient(cc grpc.ClientConnInterface) BgpPingClient {
	return &bgpPingClient{cc}
}

func (c *bgpPingClient) PingBgpPeer(ctx context.Context, in *BgpPingRequest, opts ...grpc.CallOption) (*BgpPingResponse, error) {
	out := new(BgpPingResponse)
	err := c.cc.Invoke(ctx, "/bgpping.BgpPing/PingBgpPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bgpPingClient) PingBgpPeerContinuously(ctx context.Context, in *BgpPingRequest, opts ...grpc.CallOption) (*BgpPingResponse, error) {
	out := new(BgpPingResponse)
	err := c.cc.Invoke(ctx, "/bgpping.BgpPing/PingBgpPeerContinuously", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BgpPingServer is the server API for BgpPing service.
// All implementations must embed UnimplementedBgpPingServer
// for forward compatibility
type BgpPingServer interface {
	// A simple RPC where the client sends a request to the server
	// using the stub and waits for a response to come back,
	// just like a normal function call.
	PingBgpPeer(context.Context, *BgpPingRequest) (*BgpPingResponse, error)
	// A server-side streaming RPC where the client sends a request to the server
	// and gets a stream to read a sequence of messages back.
	PingBgpPeerContinuously(context.Context, *BgpPingRequest) (*BgpPingResponse, error)
	mustEmbedUnimplementedBgpPingServer()
}

// UnimplementedBgpPingServer must be embedded to have forward compatible implementations.
type UnimplementedBgpPingServer struct {
}

func (UnimplementedBgpPingServer) PingBgpPeer(context.Context, *BgpPingRequest) (*BgpPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingBgpPeer not implemented")
}
func (UnimplementedBgpPingServer) PingBgpPeerContinuously(context.Context, *BgpPingRequest) (*BgpPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingBgpPeerContinuously not implemented")
}
func (UnimplementedBgpPingServer) mustEmbedUnimplementedBgpPingServer() {}

// UnsafeBgpPingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BgpPingServer will
// result in compilation errors.
type UnsafeBgpPingServer interface {
	mustEmbedUnimplementedBgpPingServer()
}

func RegisterBgpPingServer(s grpc.ServiceRegistrar, srv BgpPingServer) {
	s.RegisterService(&BgpPing_ServiceDesc, srv)
}

func _BgpPing_PingBgpPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BgpPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpPingServer).PingBgpPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpping.BgpPing/PingBgpPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpPingServer).PingBgpPeer(ctx, req.(*BgpPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BgpPing_PingBgpPeerContinuously_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BgpPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BgpPingServer).PingBgpPeerContinuously(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bgpping.BgpPing/PingBgpPeerContinuously",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BgpPingServer).PingBgpPeerContinuously(ctx, req.(*BgpPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BgpPing_ServiceDesc is the grpc.ServiceDesc for BgpPing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BgpPing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bgpping.BgpPing",
	HandlerType: (*BgpPingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingBgpPeer",
			Handler:    _BgpPing_PingBgpPeer_Handler,
		},
		{
			MethodName: "PingBgpPeerContinuously",
			Handler:    _BgpPing_PingBgpPeerContinuously_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bgpping/bgpping.proto",
}
