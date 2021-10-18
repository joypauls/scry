package app

import (
	"io/ioutil"
	"os"
	fp "path/filepath"

	"github.com/joypauls/scry/fst"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ShowHidden bool      `yaml:"show-hidden"`
	UseEmoji   bool      `yaml:"use-emoji"`
	InitDir    *fst.Path // where scry is initialized
	Home       *fst.Path // actual user home directory
	// Add in special character handling here, charsets, emoji swaps, etc.
}

func (c Config) Parse(file string) Config {
	// refactor to take a reader?
	// should this be checked if it exists? should just error?
	f, err := ioutil.ReadFile(fp.Join(c.Home.String(), file))
	if err != nil {
		// can't find it? ignore?
	} else {
		err = yaml.Unmarshal(f, &c)
		if err != nil {
			// problem with the file, ignore?
		}
	}
	return c
}

// Initialize with sensible defaults.
func MakeConfig() Config {
	config := Config{}
	home, err := os.UserHomeDir()
	if err == nil {
		config.Home = fst.NewPath(home)
	} else {
		config.Home = fst.NewPath("")
	}
	return config
}
