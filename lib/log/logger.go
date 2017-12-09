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

package log

import (
	// Native
	"fmt"

	// Packages
	"github.com/ttacon/chalk"
)

// Info prints info messages
func Info(msg string) {
	fmt.Print(chalk.Green, "✔ ", chalk.Reset, " ")
	fmt.Println(msg)
}

// Warn prints warning messages
func Warn(msg string) {
	fmt.Print(chalk.Yellow, "⚠ ", chalk.Reset, " ")
	fmt.Println(msg)
}

// Error prints error messages
func Error(msg string) {
	fmt.Print(chalk.Red, "✖ ", chalk.Reset, " ")
	fmt.Println(msg)
}
