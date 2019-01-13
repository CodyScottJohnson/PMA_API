package cmd

import (
	"context"
	"flag"
	"fmt"

	// mysql driver
	"Product-Management-API/pkg/database"
	"Product-Management-API/pkg/protocol/grpc"
	"Product-Management-API/pkg/protocol/grpcWeb"
	"Product-Management-API/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort    string
	GrpcWebPort string
	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.GrpcWebPort, "grpc-web-port", "", "gRPC  web port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		cfg.GRPCPort = "9000"
		//return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	if len(cfg.GrpcWebPort) == 0 {
		cfg.GrpcWebPort = "9001"
		//return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database

	database, err := db.ConnectDB("")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer database.Close()

	v1API := v1.NewProjectServiceServer(database)
	go func() {
		err := grpcWeb.RunServer(ctx, v1API, cfg.GrpcWebPort)
		if err != nil {
			fmt.Println("failed to start Grpc Web: ", err)
		}
	}()
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
