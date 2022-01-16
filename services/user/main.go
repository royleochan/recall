package main

import (
	"log"
	"net"

	"github.com/royleochan/recall/services/user/handlers"
	"github.com/royleochan/recall/services/user/userpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("User Service")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &handlers.Server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
