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

	"github.com/jhoffe/2048_go/pkg/montecarlo"
	"github.com/spf13/cobra"
)

var Iters int
var Depth int

// mcCmd represents the mc command
var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		g, duration := montecarlo.Run(Iters, Depth, montecarlo.ScoreRewardFunction)

		fmt.Printf("Finished in %dms with a score of: %d and highest brick: %d. For R=%d and D=%d \n", duration, g.Score, g.GetHighestBrickValue(), Iters, Depth)
	},
}

func init() {
	rootCmd.AddCommand(mcCmd)

	mcCmd.Flags().IntVarP(&Iters, "iters", "i", 50, "The amount of iterations to take over each random tree")
	mcCmd.Flags().IntVarP(&Depth, "depth", "d", 200, "The depth of the tree to crawl. If 0, until the game is done")
}
