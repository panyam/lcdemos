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

// tweetsCmd represents the tweets command
var tweetsCmd = &cobra.Command{
	Use: "tweets",
}

// createCmd represents the create command
var tweetsCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewTweetServiceClient(conn)
			creatorid := args[0]
			contents := args[1]
			tweet := &gen.Tweet{Ownerid: creatorid, Contents: contents}
			request := &gen.CreateTweetRequest{Creatorid: creatorid, Tweet: tweet}
			tweet, err := client.CreateTweet(ctx, request)
			if err != nil {
				return err
			}
			log.Println("Tweet Created: ", tweet)
			return nil
		})
	},
}

// tweetsDeleteCmd represents the delete command
var tweetsDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewTweetServiceClient(conn)
			tweetid := args[0]
			request := &gen.DeleteTweetRequest{Tweetid: tweetid}
			_, err := client.DeleteTweet(ctx, request)
			if err != nil {
				return err
			}
			log.Println("Tweet Deleted: ", tweetid)
			return nil
		})
	},
}

// tweetsGetCmd represents the delete command
var tweetsGetCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		connect(grpcAddr(), func(conn *grpc.ClientConn) error {
			ctx := context.Background()
			client := gen.NewTweetServiceClient(conn)
			request := &gen.BatchGetTweetsRequest{Tweetids: args}
			response, err := client.BatchGetTweets(ctx, request)
			if err != nil {
				return err
			}
			log.Println("Batch Get Response: ", response)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(tweetsCmd)
	tweetsCmd.AddCommand(tweetsCreateCmd)
	tweetsCmd.AddCommand(tweetsDeleteCmd)
	tweetsCmd.AddCommand(tweetsGetCmd)
}
