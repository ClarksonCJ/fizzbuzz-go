// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"sync"

	"github.com/ClarksonCJ/fizzbuzz/pkg/numbers"

	"github.com/spf13/cobra"
)

var inputNumbers []int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the FizzBuzz Processor over a list of numbers",
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		intChan := make(chan int, len(inputNumbers))
		r := &numbers.Reporter{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range intChan {
				fmt.Printf("%s ", r.Report(i))
			}
		}()
		for _, i := range inputNumbers {
			intChan <- i
		}
		close(intChan)
		wg.Wait()
	},
}

func init() {
	defIntSlice := make([]int, 100)
	for i := 0; i < 100; i++ {
		defIntSlice[i] = i
	}
	runCmd.PersistentFlags().IntSliceVar(&inputNumbers, "Input", defIntSlice, "Initial Set of numbers")
}
