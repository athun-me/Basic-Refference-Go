package main

import (
	"context"
	"log"
	"net"

	pb "athun/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	response := pb.TestResponse{
		Testreturn: req.Testname,
	}
	return &response, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server implementation
	pb.RegisterAuthServiceServer(s, &server{})
	// Start the server
	log.Println("Server started listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
