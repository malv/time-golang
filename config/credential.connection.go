package config

import (
	"context"
	"log"

	pb "time-project/model"

	"google.golang.org/grpc"
)

// Host ...
var Host = "camskoleksi.com:8091"

// Ctx ...
var Ctx = context.Background()

// Client ...
var Client pb.UserServiceClient

// Connected ...
func Connected() {
	conn, err := grpc.Dial(Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	Client = pb.NewUserServiceClient(conn)
}
