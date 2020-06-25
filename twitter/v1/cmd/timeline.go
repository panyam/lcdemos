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
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"leetcoach.com/demos/twitter/gen"
)

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use: "timeline",
}

// timelineSelfCmd represents the list command
var timelineSelfCmd = &cobra.Command{
	Use: "self",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewTimelineServiceClient(conn)
			userid := args[0]
			request := &gen.ListTweetsRequest{Userid: userid}
			response, err := client.ListTweets(ctx, request)
			if err != nil {
				return err
			}
			log.Printf("Self Timeline for (%d): \n", userid, response)
			return nil
		})
	},
}

// timelineHomeCmd represents the list command
var timelineHomeCmd = &cobra.Command{
	Use: "home",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewTimelineServiceClient(conn)
			userid := args[0]
			request := &gen.ListTweetsRequest{Userid: userid}
			response, err := client.GetTimeline(ctx, request)
			if err != nil {
				return err
			}
			log.Printf("Home Timeline for (%d): \n", userid, response)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(timelineCmd)
	timelineCmd.AddCommand(timelineSelfCmd)
	timelineCmd.AddCommand(timelineHomeCmd)
}
