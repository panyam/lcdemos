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
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var cfgFile string
var grpcPort int
var grpcHost string
var pageOffset int
var pageSize int

// This represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "twitter",
	Short: "Example Twitter client and server running both grpc and rest",
	Long: `To get started run the serve subcommand which will start the following:

a grpc server on localhost:10000
a http rest server on localhost:8080

twitter serve

Then you can hit it with the client:

    twitter <command>

Or over HTTP 1.1 with curl:

    curl -X POST -k https://localhost:8080/v1/twitter -d '{"value": "foo"}'
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&grpcHost, "grpcHost", "", "localhost", "Host to run GRPC server or client on")
	rootCmd.PersistentFlags().IntVarP(&grpcPort, "grpcPort", "g", 10000, "Port to run GRPC server or client on")
	rootCmd.PersistentFlags().IntVarP(&pageOffset, "pageOffset", "", 0, "Offset into the results for a paginated response")
	rootCmd.PersistentFlags().IntVarP(&pageSize, "pageSize", "", 1000, "Size of a page in a paginated response")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lcdemos/twitter.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("twitter")        // name of config file (without extension)
	viper.AddConfigPath("$HOME/.lcdemos") // adding home directory as first search path
	viper.AutomaticEnv()                  // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func connect(addr string, beforeClose func(conn *grpc.ClientConn) error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println("Serving Error: ", err)
	} else {
		err = beforeClose(conn)
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()
	}
}

func grpcAddr() string {
	return fmt.Sprintf("%s:%d", grpcHost, grpcPort)
}
