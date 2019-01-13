package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"Product-Management-API/pkg/api/v1"
)

//RunServer - Start GRPC Server
func RunServer(ctx context.Context, API v1.ProjectServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	//Register Service
	server := grpc.NewServer()
	//wrappedGrpc := grpcweb.WrapServer(server)

	v1.RegisterProjectServiceServer(server, API)

	//Gracefull Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()
	log.Println("starting gRPC server...")
	return server.Serve(listen)

}
