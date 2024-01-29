package main

import (
	"context"
	pb "github.com/briannqc/hello-grpc-go/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received AddRequest: %v\n", req)
	res := &pb.AddResponse{
		Sum: req.First + req.Second,
	}
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on: %v, err: %v\n", address, err)
	}
	log.Printf("Listening on: %v\n", address)

	defer func() {
		if err := listener.Close(); err != nil {
			log.Fatalf("Failed to close listener on: %v, err: %v\n", address, err)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve listener: %v, err: %v\n", address, err)
	}
}
