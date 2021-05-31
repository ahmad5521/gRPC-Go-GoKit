package main

import (
	"context"
	"log"
	"time"

	"elm.sa/ahmasiri/grpc-client/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMathServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Add(ctx, &pb.MathRequest{NumA: 5, NumB: 3})
	if err != nil {
		log.Fatalf("could not get result: %v", err)
	}
	log.Printf("Result: %s", r)

	r2, err2 := c.Add2(ctx, &pb.MathRequest{NumA: 3, NumB: 3})
	if err2 != nil {
		log.Fatalf("could not get result: %v", err2)
	}
	log.Printf("Result: %s", r2)
}
