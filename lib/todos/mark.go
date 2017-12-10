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
	// Ours
	"github.com/ahmed-taj/git-todos/lib/log"
)

// Mark a single todo
func Mark(todo Todo) {
	for i, t := range store.Todos {
		if t.Title == todo.Title {
			store.Todos[i].Marked = true
		} else {
			store.Todos[i].Marked = false
		}
	}
	// Write to .todos.yml
	if saveTodos() {
		log.Info("A Todo has been marked")
	}
}

// GetMarkedOrSelected todo item
func GetMarkedOrSelected(marked bool) (Todo, error) {
	if marked {
		// Try to get the marked Todo
		for _, t := range store.Todos {
			if t.Marked {
				return t, nil
			}
		}
	}
	// Failed? Ask the user to select from available Todos
	return Select(List())
}
