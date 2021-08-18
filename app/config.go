package app

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	IgnoreHidden bool `yaml:"ignore-hidden"`
	UseEmoji     bool `yaml:"use-emoji"`
}

func NewConfig(path string) Config {
	config := Config{}
	// should this be checked if it exists
	f, err := ioutil.ReadFile(path)
	if err != nil {
		// hmm
	} else {
		err = yaml.Unmarshal(f, &config)
		if err != nil {
		}
	}
	return config
}
