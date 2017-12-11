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
	"os"

	// Packages
	"github.com/spf13/cobra"
	config "github.com/tcnksm/go-gitconfig"
	"github.com/ttacon/chalk"
	survey "gopkg.in/AlecAivazis/survey.v1"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/helpers"
	"github.com/ahmed-taj/git-todos/lib/log"
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"pull", "get"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Import an issue from remote Provider (ie. GitHub) as Todo",
	Run: func(cmd *cobra.Command, args []string) {
		// Search term
		var term string

		prompt := &survey.Input{Message: "Search query"}

		if err := survey.AskOne(prompt, &term, nil); err == nil {
			// @todo support different remote names other than origin
			remote := "origin"
			url, err := config.Local(fmt.Sprintf("remote.%s.url", remote))

			if err != nil {
				log.Error(
					fmt.Sprintf("Remote (%s) is not set", chalk.Cyan.Color(remote)),
				)
				os.Exit(1)
			}

			// @todo support custom To-Dos provider!
			provider := todos.GitHubProvider{URL: url}

			log.Wait(
				fmt.Sprintf(
					"Fetching items from %s (%s)",
					chalk.Cyan.Color(provider.Name()),
					chalk.Yellow.Color(remote),
				),
			)

			// List issues from provider
			todo, err := todos.Select(todos.ImportList(term, provider))

			if err != nil {
				return
			}

			// Add the Todo
			if simple {
				// We only import issue description if --simple is not set
				todos.Add(todo.Title, "", todo.ID)
			} else {
				todos.Add(todo.Title, todo.Description, todo.ID)
			}

			// Mark it
			if marked {
				todos.Mark(todo)
			}
			// Good job!
			helpers.Encourage()
		}
	},
}

func init() {
	importCmd.Flags().BoolVarP(
		&simple,
		"simple", "s",
		false,
		"Don't import remote issue description",
	)
	importCmd.Flags().BoolVarP(
		&marked,
		"marked", "m",
		false,
		"Automatically mark the imported Todo",
	)
	appCmd.AddCommand(importCmd)
}
