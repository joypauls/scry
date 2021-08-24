package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	fp "path/filepath"

	"github.com/joypauls/scry/app"
	"github.com/joypauls/scry/fst"
)

const titleText = "Scry CLI tool"
const helpText = `Usage:
  scry                   (Basic)
  scry [flags] <path>    (Optional)

Path:
  <path> is a single optional argument that scry will try to resolve 
  to a valid starting directory. Default is the current directory.

Flags:`

func customUsageText() {
	fmt.Fprintf(os.Stderr, "%s\n\n", titleText)
	fmt.Fprintln(os.Stderr, helpText)
	flag.PrintDefaults()
}

func parseArgs(args []string, c *app.Config) {
	if len(args) == 0 {
		c.Home = fst.NewPath("")
	} else if len(args) == 1 {
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
		c.Home = fst.NewPath(parsed)
		fmt.Printf("Arg: %s\n", c.Home)
	} else {
		log.Fatal("Too many arguments supplied - zero(0) or one(1) required")
	}
}

func main() {
	defer os.Exit(0)
	// read config file or set defaults
	config := app.MakeConfig()

	// set custom usage output (-h or --help)
	flag.Usage = customUsageText

	// parse flags
	useEmojiFlag := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	showHiddenFlag := flag.Bool("a", false, "Show dotfiles/directories")
	devFlag := flag.Bool("dev", false, "Show debugging messages")
	flag.Parse()
	if *useEmojiFlag {
		config.UseEmoji = *useEmojiFlag
	} // else ignore
	if *showHiddenFlag {
		config.ShowHidden = *showHiddenFlag
	} // else ignore

	// parse remaining args
	parseArgs(flag.Args(), &config)

	if *devFlag {
		// dev log messages
		log.Print("START")
		log.Printf("home -> %s", config.Home)
		defer log.Print("EXIT")
	}

	// start the render loop
	app.Run(config)
}
