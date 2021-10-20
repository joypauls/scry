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

// This is overwritten at compile time with build flags with the current tag
// See build step in Makefile to get a sense of what happens
var version = "v0.0.0"

// Location to check for config override file
const configFile = ".config/scry/config.yaml"

const titleText = "Scry CLI tool"
const helpText = `Usage:
  scry                   (Basic)
  scry [flags] <path>    (Optional)

Path:
  <path> is a single optional argument that scry will try to resolve 
  to a valid starting directory. Default is the current directory.

Flags:`

func formatUsageText() string {
	return fmt.Sprintf("%s\n\n%s", titleText, helpText)
}

func printUsageText() {
	fmt.Fprintln(os.Stderr, formatUsageText())
	flag.PrintDefaults()
}

func parseArgs(args []string, c *app.Config) {
	if len(args) == 0 {
		c.InitDir = fst.NewPath("")
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
		c.InitDir = fst.NewPath(parsed)
		fmt.Printf("Arg: %s\n", c.InitDir)
	} else {
		log.Fatal("Too many arguments supplied - zero(0) or one(1) required")
	}
}

func main() {
	defer os.Exit(0)
	// read config file or set defaults
	config := app.MakeConfig()
	config = config.Parse(configFile)

	// set custom usage output (-h or --help)
	flag.Usage = printUsageText

	// parse flags
	useEmojiFlag := flag.Bool("e", false, "Use emoji in UI (sparingly)")
	showHiddenFlag := flag.Bool("a", false, "Show dotfiles/directories")
	versionFlag := flag.Bool("v", false, "Show build version")
	devFlag := flag.Bool("d", false, "Show debugging messages")
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}
	if *useEmojiFlag {
		config.UseEmoji = *useEmojiFlag
	} // else ignore
	if *showHiddenFlag {
		config.ShowHidden = *showHiddenFlag
	} // else ignore

	// parse remaining args
	parseArgs(flag.Args(), &config)

	if *devFlag {
		log.Print("START")
		log.Printf("home -> %s", config.InitDir)
		defer log.Print("EXIT")
	}

	// start the render loop
	render(config)
}
