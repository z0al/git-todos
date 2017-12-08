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

	// Ours
	"github.com/ahmed-taj/git-todos/lib/git"
)

// appCmd represents the base command when called without any subcommands
var appCmd = &cobra.Command{
	Use:   "git-todos [command]",
	Short: "A Git based Todos App for Developers",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 'help' and 'version' commands don't require git
		if cmd.Name() == "help" || cmd.Name() == "version" {
			return
		}

		// Do we have git?
		if !git.IsInstalled() {
			fmt.Println("Command not found: git")
			fmt.Println("Make sure Git installed and available in the PATH")
			os.Exit(1)
		}

		// Are we inside a Git repository?
		if _, err := git.GetRoot(); err != nil {
			fmt.Println("Not a Git repository (or any of the parent directories)")
			fmt.Println("You must be inside a Git repository to run commands")
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the app command and sets flags
// appropriately. This is called by main.main().
func Execute() {
	if err := appCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
