package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "wiloon.com/golang-x/grpcx/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
)

const (
	address = ":7000"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("say hello:", in.ParaFoo)
	return &pb.HelloReply{Message: "hello from server"}, nil
}
func main() {
	log.Println("server starting...")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listen:" + address)

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
