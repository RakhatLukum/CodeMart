package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/RakhatLukum/CodeMart/proto" // /Import of generated protobuf files
)

func main() {
	// Listening on port 50053
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creating a new rpc server
	grpcServer := grpc.NewServer()

	// Registering a Carservice
	pb.RegisterCartServiceServer(grpcServer, &CartServer{
		carts: make(map[int32][]int32),
	})

	log.Println("🛒 CartService is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
