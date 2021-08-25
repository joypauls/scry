package main

import (
	"os"
	"testing"

	"github.com/joypauls/scry/app"
	"github.com/joypauls/scry/fst"
)

func TestParseArgs(t *testing.T) {
	// setup
	args := []string{}
	config := app.MakeConfig()
	parseArgs(args, &config)
	// test
	got := config.Home
	want := fst.NewPath("")
	if *got != *want {
		t.Errorf("got %q, wanted %q", *got, *want)
	}

	// setup
	wd, _ := os.Getwd()
	args = []string{wd}
	config = app.MakeConfig()
	parseArgs(args, &config)
	// test
	got = config.Home
	want = fst.NewPath(wd)
	if *got != *want {
		t.Errorf("got %q, wanted %q", *got, *want)
	}
}
