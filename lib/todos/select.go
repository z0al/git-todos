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
	"sort"

	// Packages
	"github.com/ttacon/chalk"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// Select asks the user to select a single Todo from all avaiable Todos and
// returns it.
func Select(todosMap map[string]Todo) (Todo, error) {
	// Extract Todo titles
	var titles []string
	for t := range todosMap {
		titles = append(titles, t)
	}
	// Sort them
	sort.Strings(titles)

	var selected string
	prompt := &survey.Select{
		Message:  "Select an item " + chalk.Dim.TextStyle("(Use arrow keys)"),
		Options:  titles,
		PageSize: 10,
	}

	err := survey.AskOne(prompt, &selected, nil)
	if err != nil {
		return Todo{}, err
	}
	return todosMap[selected], nil
}
