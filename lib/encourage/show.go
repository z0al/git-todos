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

package encourage

import (
	// Native
	"math/rand"
	"time"

	// Packages
	"github.com/ahmed-taj/git-todos/lib/log"
)

// Show displays little encouragements while you work
func Show() {
	// Random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Encouragement List
	messages := []string{
		"Bravo! 👏",
		"Coding win! 🍸",
		"FTW! ⚡️",
		"Genius work! 🍩",
		"I see what you did there! 🙏",
		"Nice Job! 🎇",
		"Nnnnnnnailed it! ✌",
		"People like you! 💞",
		"So good! 💖",
		"Thumbs up! 👍",
		"Way to go! ✨",
		"Well done! 🎉",
		"Wow, nice change! 💗",
		"Yep! 🙆",
		"You got this 👍",
		"You rock! 🚀",
		"You're good enough! 😎",
		"You're smart enough! 💫",
	}
	// Random index
	log.Info(messages[r.Intn(len(messages))])
}
