package grpcWeb

import (
	"Product-Management-API/pkg/api/v1"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, API v1.ProjectServiceServer, port string) error {
	/*listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}*/

	//Register Service
	server := grpc.NewServer()
	wrappedGrpc := grpcweb.WrapServer(server)

	v1.RegisterProjectServiceServer(server, API)
	HttpServer := http.Server{}
	//Gracefull Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC Web server...")

			HttpServer.Shutdown(ctx)

			<-ctx.Done()
		}
	}()
	HttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			fmt.Println("GRPC Web Request")
			wrappedGrpc.ServeHTTP(resp, req)
		} else {
			message := req.URL.Path
			//message = strings.TrimPrefix(message, "/")
			message = "Page " + message + " Not Found"
			resp.Write([]byte(message))
		}
		// Fall back to other servers.
		//http.DefaultServeMux.ServeHTTP(resp, req)
	})
	HttpServer.Addr = ":" + port
	log.Println("starting Grpc-Web gateway...")
	return HttpServer.ListenAndServeTLS("./certs/localhost.pem", "./certs/localhost-key.pem") //HttpServer.Serve(listen)

}
