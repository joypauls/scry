package main

import (
	"flag"

	"github.com/joypauls/scry/app"
)

const configPath = "./config.yaml"

func main() {
	config := app.NewConfig(configPath)
	// pointer to a bool
	useEmojiFlag := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	flag.Parse()
	if *useEmojiFlag {
		config.UseEmoji = *useEmojiFlag
	} // else ignore because it wasnt supplied right?
	app.Run(config)
}
