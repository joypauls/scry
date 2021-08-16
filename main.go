package main

import (
	"flag"

	"github.com/joypauls/scry/app"
)

func main() {
	// pointer to a bool
	useEmoji := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	flag.Parse()
	app.Run(*useEmoji)
}
