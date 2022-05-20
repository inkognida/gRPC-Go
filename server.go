package main

import (
	"context"
	"flag"
	"fmt"
	"grpcadder/api/proto/adderpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	adderpb.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, req *adderpb.AddRequest) (*adderpb.AddResponse, error) {
	log.Printf("Received: %v %v", req.GetX(), req.GetY())
	return &adderpb.AddResponse{R: req.GetX() + req.GetY()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	adderpb.RegisterAdderServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failde to serve: %v", err)
	}
}
