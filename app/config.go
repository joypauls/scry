package app

import (
	"io/ioutil"
	"os"
	fp "path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/joypauls/scry/fst"
)

const configFile = ".scry.yaml"

type Config struct {
	ShowHidden bool `yaml:"show-hidden"`
	UseEmoji   bool `yaml:"use-emoji"`
	Home       *fst.Path
}

func MakeConfig() Config {
	config := Config{}
	// parse the config file if present
	path := ""
	home, err := os.UserHomeDir()
	if err == nil {
		path = fp.Join(home, configFile)
	}
	// should this be checked if it exists?
	f, err := ioutil.ReadFile(path)
	if err != nil {
		// can't find it? ignore?
	} else {
		err = yaml.Unmarshal(f, &config)
		if err != nil {
			// problem with the file, ignore?
		}
	}
	return config
}
