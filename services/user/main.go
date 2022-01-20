package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/royleochan/recall/services/user/handlers"
	"github.com/royleochan/recall/services/user/userpb"
	"github.com/royleochan/recall/services/user/utils"
	"google.golang.org/grpc"
)

func main() {
	// Initialise database connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err, db_name := utils.InitDb(ctx)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	usersCollection := client.Database(db_name).Collection("users")

	// Setup TCP connection and GRPC server with handlers
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &handlers.Server{Collection: usersCollection, UserServiceServer: userpb.UnimplementedUserServiceServer{}})

	log.Printf("Starting user service at %v", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
