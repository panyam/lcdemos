package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"leetcoach.com/demos/twitter/gen"
	"leetcoach.com/demos/twitter/services"
	"log"
	"net"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	httpPort = flag.Int("http_port", 8080, "The HTTP server port")
	grpcPort = flag.Int("grpc_port", 10000, "The GRPC server port")
)

func main() {
	flag.Parse()
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	gen.RegisterTimelineServiceServer(grpcServer, services.NewTimelineService())
	gen.RegisterTweetServiceServer(grpcServer, services.NewTweetService())
	gen.RegisterFollowServiceServer(grpcServer, services.NewFollowService())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
