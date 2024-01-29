package main

import (
	"context"
	pb "github.com/briannqc/hello-grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial %v, err: %v\n", address, err)
	}
	log.Printf("Dialed %v\n", address)

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close connection to: %v, err: %v\n", address, err)
		}
	}()

	client := pb.NewCalculatorServiceClient(conn)
	res, err := client.Add(context.Background(), &pb.AddRequest{First: 10, Second: 3})
	if err != nil {
		log.Fatalf("Failed to call CalculatorService.Add(), err: %v\n", err)
	}
	log.Printf("Result: %v\n", res.Sum)
}
