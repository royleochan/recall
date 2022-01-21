package main

import (
	"log"
	"net"

	"github.com/royleochan/recall/services/flashcard/database"
	"google.golang.org/grpc"
)

func main() {
	// Initialize DB connection
	database.InitDb()
	database.Migrate()
	db, _ := database.DBCon.DB()
	defer db.Close()

	// Setup TCP connection and GRPC server with handlers
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// userpb.RegisterUserServiceServer(grpcServer, &handlers.Server{Collection: usersCollection, UserServiceServer: userpb.UnimplementedUserServiceServer{}})

	log.Printf("Starting flashcard service at %v", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
