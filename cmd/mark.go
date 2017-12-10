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

	// Ours
	"github.com/ahmed-taj/git-todos/lib/helpers"
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var markCmd = &cobra.Command{
	Use:     "mark",
	Aliases: []string{"pick"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Mark a single Todo",
	Run: func(cmd *cobra.Command, args []string) {
		// Ask the user to select a Todo
		todo, err := todos.Select(todos.List())

		if err == nil {
			todos.Mark(todo)
			helpers.Encourage()
		}
	},
}

func init() {
	appCmd.AddCommand(markCmd)
}
