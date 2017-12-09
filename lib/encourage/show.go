package encourage

import (
	// Native
	"math/rand"
	"time"

	"github.com/ahmed-taj/git-todos/lib/log"
	// Packages
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
