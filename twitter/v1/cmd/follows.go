/*
Copyright © 2020 Sriram Panyam <sri.panyam@gmail.com>

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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"leetcoach.com/demos/twitter/gen"
	"log"
)

// followsCmd represents the follow command
var followsCmd = &cobra.Command{
	Use: "follows",
}

// createCmd represents the create command
var followsCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewFollowServiceClient(conn)
			for i := 0; i < len(args); i += 2 {
				log.Printf("Follow: (%s -> %s): ", args[i], args[i+1])
				request := &gen.CreateFollowRequest{Followerid: args[i], Leaderid: args[i+1]}
				_, err := client.CreateFollow(ctx, request)
				if err != nil {
					return err
				}
				log.Println("Deleted")
			}
			return nil
		})
	},
}

// followsDeleteCmd represents the delete command
var followsDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewFollowServiceClient(conn)
			for i := 0; i < len(args); i += 2 {
				request := gen.DeleteFollowRequest{Followerid: args[i], Leaderid: args[i+1]}
				_, err := client.DeleteFollow(ctx, &request)
				log.Printf("Follow: (%s -> %s): ", args[i], args[i+1])
				if err != nil {
					return err
				}
				log.Println("Created")
			}
			return nil
		})
	},
}

// followsListCmd represents the list command
var followsListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewFollowServiceClient(conn)
			request := gen.GetFollowRequest{Userid: args[0], PageSize: int32(pageSize)}
			response, err := client.GetFollowers(ctx, &request)
			log.Printf("Followers of %s: ", args[0])
			if err != nil {
				return err
			}
			log.Println(response)
			return nil
		})
	},
}

// followsListUpCmd represents the list command
var followsListUpCmd = &cobra.Command{
	Use: "listup",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewFollowServiceClient(conn)
			request := gen.GetFollowRequest{Userid: args[0], PageSize: int32(pageSize)}
			response, err := client.GetFollowees(ctx, &request)
			log.Printf("Followees of %s: ", args[0])
			if err != nil {
				return err
			}
			log.Println(response)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(followsCmd)
	followsCmd.AddCommand(followsCreateCmd)
	followsCmd.AddCommand(followsDeleteCmd)
	followsCmd.AddCommand(followsListCmd)
	followsCmd.AddCommand(followsListUpCmd)
}
