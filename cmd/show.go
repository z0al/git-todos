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
	"github.com/ahmed-taj/git-todos/lib/todos"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Args:  cobra.MaximumNArgs(0),
	Short: "Show Todo details",
	Run: func(cmd *cobra.Command, args []string) {
		// Get marked todo (if --marked) or ask the user to select one
		todo, err := todos.GetMarkedOrSelected(marked)

		if err == nil {
			todos.FormatAndPrint(todo)
		}
	},
}

func init() {
	appCmd.AddCommand(showCmd)
}
