package main

import (
	"google.golang.org/grpc"
	"log"
	pb "wiloon.com/golang-x/grpcx/hello"
	"context"
)

const address = "localhost:7000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{ParaFoo: "foo"})
	if err != nil {
		log.Fatalf("failed to say hello %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
