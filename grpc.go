package main

import (
	"context"
	"fmt"
	"net"

	"github.com/seth-epps/hello-go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

const authorityMeta = ":authority"

type grpcServer struct {
	protos.UnimplementedHelloServer
}

func (s *grpcServer) SayHello(ctx context.Context, req *protos.HelloRequest) (*protos.HelloResponse, error) {
	msg := "Hello From Go!"
	peer, _ := peer.FromContext(ctx)
	ip := peer.Addr.String()

	authority := ""
	if metaAuth := metadata.ValueFromIncomingContext(ctx, authorityMeta); len(metaAuth) != 0 {
		authority = metaAuth[0]
	}

	return &protos.HelloResponse{
		Ip:        &ip,
		Message:   &msg,
		Authority: &authority,
	}, nil

}

func startGrpc() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return fmt.Errorf("failed to listen on port 9090: %w", err)
	}

	s := grpc.NewServer()
	protos.RegisterHelloServer(s, &grpcServer{})
	reflection.Register(s)

	fmt.Printf("gRPC server listening at %v\n", lis.Addr())
	return s.Serve(lis)
}
