package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/testdata"
	"leetcoach.com/demos/twitter/gen"
	"leetcoach.com/demos/twitter/services"
	"log"
	"net"
	"net/http"
	"sync"
	// "strings"
)

func StartHttpServer(httpPort int, callback func(mux *http.ServeMux, gmux *runtime.ServeMux) error) {
	mux := http.NewServeMux()
	gmux := runtime.NewServeMux()
	addr := fmt.Sprintf(":%d", httpPort)
	mux.Handle("/", gmux)

	callback(mux, gmux)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}

func StartGrpcServer(grpcPort int, callback func(*grpc.Server) error) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	callback(server)
	server.Serve(lis)
	return server
}

var (
	httpPort = flag.Int("http_port", 8080, "The HTTP server port")
	grpcPort = flag.Int("grpc_port", 10000, "The GRPC server port")
)

func main() {
	flag.Parse()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go StartGrpcServer(*grpcPort, func(server *grpc.Server) error {
		gen.RegisterTimelineServiceServer(server, services.NewTimelineService())
		gen.RegisterTweetServiceServer(server, services.NewTweetService())
		gen.RegisterFollowServiceServer(server, services.NewFollowService())
		return nil
	})

	go StartHttpServer(*httpPort, func(mux *http.ServeMux, gmux *runtime.ServeMux) error {
		ctx := context.Background()
		addr := fmt.Sprintf(":%d", httpPort)
		dopts := []grpc.DialOption{}
		err := gen.RegisterFollowServiceHandlerFromEndpoint(ctx, gmux, addr, dopts)
		if err != nil {
			fmt.Printf("serve: %v\n", err)
			return err
		}
		err = gen.RegisterTweetServiceHandlerFromEndpoint(ctx, gmux, addr, dopts)
		if err != nil {
			fmt.Printf("serve: %v\n", err)
			return err
		}
		err = gen.RegisterTimelineServiceHandlerFromEndpoint(ctx, gmux, addr, dopts)
		if err != nil {
			fmt.Printf("serve: %v\n", err)
			return err
		}
		return nil
	})

	wg.Wait()
}
