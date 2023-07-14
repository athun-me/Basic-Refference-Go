package main

import (
	"athun/pb"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)
	request := &pb.TestRequest{
		Testname: "This is the return",
	}
	res, err := client.Test(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call GetData: %v", err)
	}
	log.Printf("Response: %s", res.Testreturn)

}
