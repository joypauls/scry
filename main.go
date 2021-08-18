package main

import (
	"flag"
	"os"
	fp "path/filepath"

	"github.com/joypauls/scry/app"
)

const configFile = ".scry.yaml"

func main() {
	// parse the config file if present
	configPath := ""
	home, err := os.UserHomeDir()
	if err == nil {
		configPath = fp.Join(home, configFile)
	}
	// should check if this exists
	config := app.NewConfig(configPath)
	// pointer to a bool
	useEmojiFlag := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	flag.Parse()
	if *useEmojiFlag {
		config.UseEmoji = *useEmojiFlag
	} // else ignore because it wasnt supplied right?
	app.Run(config)
}
