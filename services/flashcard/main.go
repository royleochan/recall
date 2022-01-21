package main

import (
	"log"
	"net"

	"github.com/royleochan/recall/services/flashcard/database"
	"github.com/royleochan/recall/services/flashcard/flashcardpb"
	"github.com/royleochan/recall/services/flashcard/handlers"
	"google.golang.org/grpc"
)

func main() {
	// Initialize DB connection
	database.InitDb()
	database.Migrate()
	db, _ := database.DBCon.DB()
	defer db.Close()

	// Setup TCP connection and GRPC server with handlers
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	flashcardpb.RegisterFlashcardServiceServer(grpcServer, &handlers.Server{Db: database.DBCon, FlashcardServiceServer: flashcardpb.UnimplementedFlashcardServiceServer{}})

	log.Printf("Starting flashcard service at %v", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
