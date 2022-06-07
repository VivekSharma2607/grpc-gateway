package main

import (
	"context"
	"fmt"
	pb "gateway/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTaskServer
}

func (s *server) User(ctx context.Context , req *pb.UserRequest) (*pb.UserResponse , error) {
	return &pb.UserResponse{} , nil
}

func (s *server) Echo(ctx context.Context , req *pb.HelloRequest) (*pb.HelloRequest , error) {
	return req , nil
}

func main() {
	fmt.Print("Connected")
	listen , err := net.Listen("tcp" , "localhost:8080")
	if err != nil {
		log.Fatalf("Error while connecting to the server : %v" , err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s , &server{})
	err = s.Serve(listen) 
	if err != nil {
		log.Fatalf("Error  while serving : %v" , err)
	}
}