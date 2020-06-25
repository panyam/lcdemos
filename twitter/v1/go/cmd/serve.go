/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	// "encoding/json"
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
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Run,
}

var httpPort int
var grpcPort int

func init() {
	rootCmd.AddCommand(serveCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().IntVarP(&httpPort, "httpPort", "p", 8080, "Port to run HTTP server on")
	serveCmd.Flags().IntVarP(&grpcPort, "grpcPort", "g", 10000, "Port to run GRPC server on")
}

func Run(cmd *cobra.Command, args []string) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go StartGrpcServer(grpcPort, func(server *grpc.Server) error {
		gen.RegisterTimelineServiceServer(server, services.NewTimelineService())
		gen.RegisterTweetServiceServer(server, services.NewTweetService())
		gen.RegisterFollowServiceServer(server, services.NewFollowService())
		return nil
	})

	go StartHttpServer(httpPort, func(mux *http.ServeMux, gmux *runtime.ServeMux) error {
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
	fmt.Printf("Starting Http Server on port %d ...\n", httpPort)
	server.ListenAndServe()
}

func StartGrpcServer(grpcPort int, callback func(*grpc.Server) error) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	callback(server)
	fmt.Printf("Starting GRPC Server on port %d ...\n", grpcPort)
	server.Serve(lis)
	return server
}
