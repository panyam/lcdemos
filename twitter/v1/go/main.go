package main

import (
	// "encoding/json"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/testdata"
	"leetcoach.com/demos/twitter/gen"
	"leetcoach.com/demos/twitter/services"
	"log"
	"net"
	"net/http"
	// "strings"
)

type GrpcServer struct {
	port   Int
	ipConn *net.IPConn
	server *grpc.Server
}

type HttpServer struct {
	port    Int
	addr    string
	ipConn  *net.IPConn
	httpMux *ServeMux
	grpcMux *ServeMux
	server  *http.Server
}

func NewHttpServer(httpPort Int) {
	mux := http.NewServeMux()
	gmux := runtime.NewServeMux()
	addr := fmt.Sprintf(":%d", httpPort)
	ipConn, err := net.Listen("tcp", fmt.Sprintf(":%d", httpPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &http.Server{
		Addr:      addr,
		Handler:   handlerFunc,
		TLSConfig: nil,
	}

	mux.Handle("/", gmux)

	httpServer = &HttpServer{
		httpMux: mux,
		grpcMux: gmux,
		port:    httpPort,
		addr:    addr,
		ipConn:  ipConn,
		server:  server,
	}
	return httpServer
}

func NewGrpcServer(grpcPort Int) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	grpcServer = &GrpcServer{
		ipConn: lis,
		server: server,
		port:   grpcPort,
	}
	return grpcServer
}

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	httpPort = flag.Int("http_port", 8080, "The HTTP server port")
	grpcPort = flag.Int("grpc_port", 10000, "The GRPC server port")
)

func main() {
	flag.Parse()
	grpcServer := NewGrpcServer(*grpcPort)
	gen.RegisterTimelineServiceServer(grpcServer.server, services.NewTimelineService())
	gen.RegisterTweetServiceServer(grpcServer.server, services.NewTweetService())
	gen.RegisterFollowServiceServer(grpcServer.server, services.NewFollowService())

	httpServer := NewHttpServer(*httpPort)
	dopts := []grpc.DialOption{}
	err := gen.RegisterFollowServiceHandlerFromEndpoint(httpServer.ctx, httpServer.gmux, httpServer.addr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return nil
	}
	err := gen.RegisterTweetServiceHandlerFromEndpoint(httpServer.ctx, httpServer.gmux, httpServer.addr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return nil
	}
	err := gen.RegisterTimelineServiceHandlerFromEndpoint(httpServer.ctx, httpServer.gmux, httpServer.addr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return nil
	}

	/*
		httpServer.httpMuxes["/swagger/tweets.json"].HandleFunc("/swagger/tweets.json", func(w http.ResponseWriter, req *http.Request) {
			io.Copy(w, strings.NewReader(pb.Swagger))
		})
		httpServer.httpMuxes["/swagger/timeline.json"].HandleFunc("/swagger/timeline.json", func(w http.ResponseWriter, req *http.Request) {
			io.Copy(w, strings.NewReader(pb.Swagger))
		})
		httpServer.httpMuxes["/swagger/follows.json"].HandleFunc("/swagger/follows.json", func(w http.ResponseWriter, req *http.Request) {
			io.Copy(w, strings.NewReader(pb.Swagger))
		})
	*/
}
