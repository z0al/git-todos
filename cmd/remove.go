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
	// Packages
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	survey "gopkg.in/AlecAivazis/survey.v1"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/encourage"
	"github.com/ahmed-taj/git-todos/lib/log"
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "delete"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Remove existing Todo",
	Run: func(cmd *cobra.Command, args []string) {
		// Get marked todo (if --marked) or ask the user to select one
		todo, err := todos.GetMarkedOrSelected(marked)

		if err == nil {
			yes := false
			log.Warn("This Todo will be removed: " + chalk.Cyan.Color(todo.Title))
			prompt := &survey.Confirm{Message: "Are you sure?"}
			survey.AskOne(prompt, &yes, nil)

			if yes {
				todos.Remove(todo)
				// :)
				encourage.Show()
			}
		}
	},
}

func init() {
	removeCmd.Flags().BoolVarP(
		&marked,
		"marked", "m",
		false,
		"Automatically select the marked Todo (if any)",
	)

	appCmd.AddCommand(removeCmd)
}
