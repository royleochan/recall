package main

import (
	"log"

	"github.com/royleochan/recall/services/user/src/userpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := userpb.NewUserServiceClient(conn)

	response, err := c.GetUser(context.Background(), &userpb.GetUserRequest{UserId: 2})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}
	log.Printf("Response from server: %d", response.UserId)
}
