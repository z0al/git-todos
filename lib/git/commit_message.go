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
	"strings"

	// Packages
	wordwrap "github.com/mitchellh/go-wordwrap"
)

// CommitMessage type
type CommitMessage struct {
	Type    string
	Scope   string
	Subject string
	Body    string
	Close   string
}

// Format outputs Conventional Commits spec friendly commit message
// To keep it simple; I ignored the breaking changes segment. I think using the
// CLI to write breaking changes is hard, so, you will propably write the commit
// Message using an editor instead.
// Ref: https://conventionalcommits.org
func (c CommitMessage) Format() string {
	// HEADER
	var msg = c.Type
	if c.Scope != "" {
		msg += "(" + c.Scope + ")"
	}
	msg += ": " + c.Subject + "\n\n"

	// BODY
	if c.Body != "" {
		msg += wordwrap.WrapString(c.Body, 79) + "\n\n"
	}
	// FOOTER
	if c.Close != "" {
		msg += c.issuesList()
	}
	return strings.TrimSpace(msg)
}

// 1,2,3 => Closes #1, #2, #3
func (c CommitMessage) issuesList() string {
	var numbers []string
	for _, item := range strings.Split(c.Close, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			numbers = append(numbers, "#"+item)
		}
	}
	if len(numbers) > 0 {
		return "Closes " + strings.Join(numbers, ", ")
	}
	return ""
}
