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
	"io/ioutil"
	"os"
	"path"

	// Packages
	yaml "gopkg.in/yaml.v2"

	// Ours
	"github.com/ahmed-taj/git-todos/lib/git"
	"github.com/ahmed-taj/git-todos/lib/log"
)

// Todo represents a single Todo
type Todo struct {
	Title       string
	Description string
	Marked      bool
}

// TodoCollection represents an array of Todos
type TodoCollection struct {
	Todos []Todo
}

// Globals
var (
	// Absolute path of .todos.yml
	filename string

	// In memory Todos store
	store TodoCollection
)

func init() {
	// We don't expect an error here since we already check this on the top-level
	// PreRun hook
	dir, _ := git.GetRoot()
	filename = path.Join(dir, ".todos.yml")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Empty todos list
		store = TodoCollection{}
	} else {
		// Attempt to read todos file
		src, err := ioutil.ReadFile(filename)
		format := "An error occured when trying to read Todos from '%s'"

		if err != nil {
			log.Error(fmt.Sprintf(format, filename))
			os.Exit(1)
		}

		// Attempt to read existing todos
		err = yaml.Unmarshal(src, &store)
		if err != nil {
			log.Error(fmt.Sprintf(format, filename))
			os.Exit(1)
		}
	}
}

// Write current todos to disk
func saveTodos() bool {
	output, _ := yaml.Marshal(&store)
	if err := ioutil.WriteFile(filename, output, 0777); err != nil {
		return false
	}
	return true
}
