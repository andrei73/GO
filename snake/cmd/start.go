// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/andrei73/GO/snake/pkg/random"
	"github.com/spf13/cobra"
	"time"
	"github.com/spf13/viper"
)
var file string
// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the snake application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		var n, x int
		n = 10
		x = 2
		file := cmd.Flag("file").Value.String()
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
		for {
			time.Sleep(time.Duration(x) * time.Second)
			fmt.Println(random.Generate(n))
			fmt.Println(file)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("file", "f", "", "Add a config file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
