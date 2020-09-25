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
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jhoffe/2048_go/pkg/montecarlo"
	"github.com/spf13/cobra"
)

var SampleSize int
var Step bool

var ItersA int
var IterStepSize int

var DepthA int
var DepthStepSize int

var Result string

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Create(Result)
		if err != nil {
			panic(err)
		}

		writer := csv.NewWriter(file)
		var mu sync.Mutex
		defer func() {
			writer.Flush()
			file.Close()
		}()
		writer.Write([]string{"N", "R", "D", "Score", "HighestBrick", "Time"})

		var wg sync.WaitGroup

		for n := 0; n < SampleSize; n++ {
			if Step {
				for r := IterStepSize; r <= ItersA; r = r + IterStepSize {
					for d := DepthStepSize; d <= DepthA; d = d + DepthStepSize {
						wg.Add(1)
						go func(wg *sync.WaitGroup, r, d, n int) {
							fmt.Printf("Beginning R=%d and D=%d and N=%d \n", r, d, n)
							g, duration := montecarlo.Run(r, d, montecarlo.ScoreRewardFunction)

							fmt.Printf("Finished in %dms with a score of: %d and highest brick: %d. For R=%d and D=%d and N=%d \n",
								duration, g.Score, g.GetHighestBrickValue(), r, d, n)

							mu.Lock()
							writer.Write([]string{strconv.Itoa(n), strconv.Itoa(r), strconv.Itoa(d), strconv.Itoa(int(g.Score)),
								strconv.Itoa(g.GetHighestBrickValue()), strconv.Itoa(int(duration))})
							mu.Unlock()

							wg.Done()
						}(&wg, r, d, n)
					}
				}
			} else {
				wg.Add(1)
				go func(wg *sync.WaitGroup, r, d, n int) {
					fmt.Printf("Beginning R=%d and D=%d and N=%d \n", ItersA, DepthA, SampleSize)

					g, duration := montecarlo.Run(ItersA, DepthA, montecarlo.ScoreRewardFunction)
					fmt.Printf("Finished in %dms with a score of: %d and highest brick: %d. For R=%d and D=%d and N=%d \n",
						duration, g.Score, g.GetHighestBrickValue(), ItersA, DepthA, SampleSize)

					mu.Lock()
					writer.Write([]string{strconv.Itoa(n), strconv.Itoa(ItersA), strconv.Itoa(DepthA), strconv.Itoa(int(g.Score)),
						strconv.Itoa(g.GetHighestBrickValue()), strconv.Itoa(int(duration))})
					mu.Unlock()

					wg.Done()
				}(&wg, ItersA, DepthA, n)
			}
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().IntVarP(&SampleSize, "size", "n", 30, "")
	analyzeCmd.Flags().BoolVarP(&Step, "step", "s", false, "")

	analyzeCmd.Flags().IntVarP(&ItersA, "iters", "i", 50, "The amount of iterations to take over each random tree")
	analyzeCmd.Flags().IntVarP(&IterStepSize, "itersteps", "I", 1, "The step size between each run")

	analyzeCmd.Flags().IntVarP(&DepthA, "depth", "d", 200, "The depth of the tree to crawl. If 0, until the game is done")
	analyzeCmd.Flags().IntVarP(&DepthStepSize, "depthsteps", "D", 1, "The step size between each run")

	analyzeCmd.Flags().StringVarP(&Result, "result", "r", "results.csv", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
