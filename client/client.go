package main

import (
	"context"
	"fmt"
	pb "gateway/gen/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn , err := grpc.Dial("localhost:8080" , grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while starting server : %v" , err)
	}
	client := pb.NewTaskClient(conn)
	resp , err := client.Echo(context.Background() , &pb.HelloRequest{
		Msg: "Hello World!",
	})
	if err != nil {
		log.Fatalf("Erorr while serving message on the server side : %v" , err)
	}
	fmt.Println(resp.Msg)
}
