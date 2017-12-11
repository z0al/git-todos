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

package git

import (
	// Native
	"os/exec"
	"strings"
)

// IsInstalled checks if 'git' command exists in the PATH
func IsInstalled() bool {
	if _, err := exec.LookPath("git"); err != nil {
		return false
	}
	return true
}

// GetRoot returns the absolute path of top-level git directory
func GetRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	root, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(root)), err
}
