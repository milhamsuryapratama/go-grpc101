package main

import (
	"context"
	"grpc-101/grpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greetingService struct {
	hello.UnimplementedUserServiceServer
}

func (s *greetingService) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := &hello.HelloResponse{
		Message: "Hello " + req.Name,
	}
	return response, nil
}

func (s *greetingService) SayGoodbye(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := &hello.HelloResponse{
		Message: "Goodbye " + req.Name,
	}
	return response, nil
}

func main() {
	// start gRPC server disini
	port_server := ":8081"
	listener, err := net.Listen("tcp", port_server)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", port_server, err)
	}

	s := grpc.NewServer()

	hello.RegisterUserServiceServer(s, &greetingService{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
