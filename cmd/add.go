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
	survey "gopkg.in/AlecAivazis/survey.v1"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/helpers"
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "new"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Add a new Todo",
	Run: func(cmd *cobra.Command, args []string) {
		// The questions to ask
		questions := []*survey.Question{
			{
				Name: "title",
				Validate: survey.ComposeValidators(
					survey.MinLength(1),
					survey.MaxLength(70),
				),
				Prompt: &survey.Input{
					Message: helpers.RequiredField("Title"),
					Help:    " A friendly, meaningful, single-line description",
				},
			},
		}

		if !simple {
			// We only ask for description if --simple flag is not set
			questions = append(questions, &survey.Question{
				Name: "description",
				Prompt: &survey.Input{
					Message: "Description",
					Help:    " A longer description for this Todo",
				},
			})
		}

		// The answers will be written to this struct
		answers := todos.Todo{}

		// Finally, Ask!
		if err := survey.Ask(questions, &answers); err == nil {
			// Add the Todo
			todos.Add(answers.Title, answers.Description, 0)

			// --marked?
			if marked {
				todos.Mark(answers)
			}
			helpers.Encourage()
		}
	},
}

func init() {
	addCmd.Flags().BoolVarP(
		&simple,
		"simple", "s",
		false,
		"Don't ask for long description",
	)
	addCmd.Flags().BoolVarP(
		&marked,
		"marked", "m",
		false,
		"Automatically mark the new Todo",
	)

	appCmd.AddCommand(addCmd)
}
