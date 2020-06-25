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
	"fmt"

	"github.com/spf13/cobra"
)

// tweetsCmd represents the tweets command
var tweetsCmd = &cobra.Command{
	Use: "tweets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tweets called")
	},
}

// createCmd represents the create command
var tweetsCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

// tweetsDeleteCmd represents the delete command
var tweetsDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

// tweetsListCmd represents the list command
var tweetsListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(tweetsCmd)
	tweetsCmd.AddCommand(tweetsCreateCmd)
	tweetsCmd.AddCommand(tweetsDeleteCmd)
	tweetsCmd.AddCommand(tweetsListCmd)
}
