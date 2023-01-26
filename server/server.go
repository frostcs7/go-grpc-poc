// Package main implements a server for HelloWorld service.
package main

import (
	"fmt"
	pb "github.com/frostcs7/go-grpc-poc/gen/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = ":50051"
)

// Server is used to implement gen.protos.HelloWorldServer
type Server struct {
	pb.UnsafeHelloWorldServer
}

// SayHello implements gen.protos.HelloWorldServer
func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Recieve the message from client: %s", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &s)
	reflection.Register(grpcServer)
	log.Println("Starting gRPC Server!")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc at %v %v", grpcPort, err)
	}
}
