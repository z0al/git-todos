// Copyright © 2017 Ahmed T. Ali <ah.tajelsir@gmail.com>
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
	// Native
	"fmt"
	"sort"

	// Packages
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "List available Todos",
	Run: func(cmd *cobra.Command, args []string) {
		// Get all Todos from store
		todos := todos.List()
		if len(todos) == 0 {
			fmt.Println("All done. You rock! 🚀")
		} else {
			fmt.Println()

			// Extract Todo titles
			var titles []string
			for t := range todos {
				titles = append(titles, t)
			}
			// Sort them
			sort.Strings(titles)

			count := 0
			for _, t := range titles {
				count++
				// e.g 1)
				n := fmt.Sprint(count) + ")"
				fmt.Printf("%s %s\n", chalk.Bold.TextStyle(n), t)
			}
			fmt.Println("\nYou've got this 😎")
		}
	},
}

func init() {
	appCmd.AddCommand(listCmd)
}
