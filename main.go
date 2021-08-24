package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	fp "path/filepath"

	"github.com/joypauls/scry/app"
)

const configFile = ".scry.yaml"

const titleText = "Scry CLI tool"
const helpText = `Usage:
        scry                  (Basic)
        scry [flags] <path>   (Advanced)

Path:
        <path> is a single optional argument that scry will try to resolve 
        to a valid starting directory. Default is the current directory.

Flags:
`

func customUsageText() {
	fmt.Fprintf(os.Stderr, "%s\n\n", titleText)
	fmt.Fprintln(os.Stderr, helpText)
	flag.PrintDefaults()
}

func main() {
	// parse the config file if present
	configPath := ""
	home, err := os.UserHomeDir()
	if err == nil {
		configPath = fp.Join(home, configFile)
	}
	// should check if this exists
	config := app.NewConfig(configPath)

	// custom usage output
	flag.Usage = customUsageText

	// parse args
	useEmojiFlag := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	showHiddenFlag := flag.Bool("a", false, "Show dotfiles/directories")
	flag.Parse()

	// intended behavior is <=1, which is a path or resolve to path
	args := flag.Args()
	if len(args) == 1 {
		parsed, err := fp.Abs(args[0])
		if err != nil {
			log.Fatalf("Couldn't parse the path: %s", args[0])
		}
		fi, err := os.Stat(parsed)
		if os.IsNotExist(err) {
			log.Fatalf("No such file or directory: %s", args[0])
		} else if !fi.IsDir() {
			parsed = fp.Dir(parsed)
		}
		config.StartPath = parsed
		fmt.Printf("Arg: %s\n", config.StartPath)
	} else if len(args) > 1 {
		log.Fatal("Too many arguments supplied - zero(0) or one(1) required")
	}

	if *useEmojiFlag {
		config.UseEmoji = *useEmojiFlag
	} // else ignore because it wasnt supplied right?
	if *showHiddenFlag {
		config.ShowHidden = *showHiddenFlag
	} // else ignore because it wasnt supplied right?

	// dev log messages, should remove for release
	log.Print("Starting scry")
	defer func() {
		log.Print("Exiting properly")
		os.Exit(0)
	}()

	// start the render loop
	app.Run(config)
}
