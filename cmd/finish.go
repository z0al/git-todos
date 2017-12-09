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
	"os"

	// Packages
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	survey "gopkg.in/AlecAivazis/survey.v1"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/git"
	"github.com/ahmed-taj/git-todos/lib/log"
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var finishCmd = &cobra.Command{
	Use:     "finish",
	Aliases: []string{"done"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Finish a Todo and commit staged changes",
	Run: func(cmd *cobra.Command, args []string) {
		// Ask the user to select a Todo
		todo, err := todos.Select()

		if err == nil {
			// Make a field as optional
			optionalField := func(str string) string {
				return str + " " + chalk.Dim.TextStyle("(Optional)")
			}

			// The questions to ask
			questions := []*survey.Question{
				{
					Name:     "type",
					Validate: survey.Required,
					Prompt: &survey.Input{
						Message: "Type of change that you're committing",
						Help:    " E.g., a fix, feat, chore ..etc",
					},
				},
				{
					Name: "scope",
					Prompt: &survey.Input{
						Message: optionalField("Denote the scope of this change"),
						Help:    " Additional contextual information to commit's type",
					},
				},
				{
					Name: "subject",
					Validate: survey.ComposeValidators(
						survey.MinLength(1),
						survey.MaxLength(70),
					),
					Prompt: &survey.Input{
						Message: "Short description",
						Help:    " A short, meaningful, description for your change",
						Default: todo.Title,
					},
				},
				{
					Name: "body",
					Prompt: &survey.Input{
						Message: optionalField("Longer description of the change"),
						Help:    " More detailed description for your chagnes",
					},
				},
				{
					Name: "close",
					Prompt: &survey.Input{
						Message: optionalField("List any issues closed by this change"),
						Help:    " Comma-separated list of issue numbers to be closed",
					},
				},
			}

			// The answers will be written to this struct
			answers := git.CommitMessage{}

			// Finally, Ask!
			if err := survey.Ask(questions, &answers); err == nil {
				_, err := git.Commit(answers)
				if err != nil {
					log.Error("Failed to commit. Did you forget to stage your changes?")
					os.Exit(1)
				}
				log.Info("Your work has been commited")
				// The todo should no longer exist; git log to see your history
				todos.Remove(todo)
			}
		}
	},
}

func init() {
	appCmd.AddCommand(finishCmd)
}