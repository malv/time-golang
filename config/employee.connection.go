package config

import (
	"context"
	"log"
	proto "time-project/proto/model"

	"google.golang.org/grpc"
)

// EmployeeServer ...
var EmployeeServer = "localhost:8888"

// CtxEmployee ...
var CtxEmployee = context.Background()

// EmployeeClient ...
var EmployeeClient proto.EmployeeServiceClient

// EmployeeHostConnection ...
func EmployeeHostConnection() {
	conn, err := grpc.Dial(EmployeeServer, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("Not connected err => ", err)
	}

	EmployeeClient = proto.NewEmployeeServiceClient(conn)
}
