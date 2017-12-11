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

package todos

import (
	// Native
	"fmt"

	// Packages
	"github.com/ttacon/chalk"
)

// List all todos
func List() map[string]Todo {
	dict := make(map[string]Todo)
	for _, t := range store.Todos {
		key := t.Title + formatID(t.ID)

		dict[key] = t
	}
	return dict
}

// FormatAndPrint a Todo item to the console
func FormatAndPrint(t Todo) {
	fmt.Println()

	// Title
	fmt.Printf("%s", chalk.Underline.TextStyle(t.Title))

	// Issue ID
	fmt.Println(formatID(t.ID))
	fmt.Println()

	// Description
	if t.Description != "" {
		fmt.Println(t.Description)
	} else {
		fmt.Println("(No description)")
	}
}

func formatID(id int) string {
	if id != 0 {
		str := fmt.Sprintf("#%d", id)
		return fmt.Sprintf("%s (%s)", chalk.Reset, chalk.Magenta.Color(str))
	}
	return ""
}
