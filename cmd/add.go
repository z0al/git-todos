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

	// Packages
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Flag
var simple bool

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "new"},
	Short:   "Add a new Todo item",
	Run: func(cmd *cobra.Command, args []string) {
		// The questions to ask
		questions := []*survey.Question{
			{
				Name:     "title",
				Validate: survey.Required,
				Prompt: &survey.Input{
					Message: "Title",
					Help:    " A friendly, meaningful, single-line description for this Todo",
				},
			},
		}

		if !simple {
			// We only ask for additional if --simple flag is false
			questions = append(questions, &survey.Question{
				Name: "description",
				Prompt: &survey.Input{
					Message: "Description",
					Help:    " A longer description for this Todo. Markdown is supported",
				},
			})
		}

		// The answers will be written to this struct
		answers := struct {
			Title       string
			Description string
		}{}

		// Finaly, Ask!
		survey.Ask(questions, &answers)

		fmt.Println("title is ", answers.Title)
		fmt.Println("Desc is ", answers.Description)
		fmt.Println("DONE")
	},
}

func init() {
	appCmd.AddCommand(addCmd)

	// Flag: -s, --simple, to simplify new Todos creation
	addCmd.Flags().BoolVarP(
		&simple,
		"simple", "s",
		false,
		"Don't ask for long description",
	)
}
