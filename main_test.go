package main

import (
	"os"
	"strings"
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
	got := config.InitDir
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
	got = config.InitDir
	want = fst.NewPath(wd)
	if *got != *want {
		t.Errorf("got %q, wanted %q", *got, *want)
	}
}

func TestUsageText(t *testing.T) {
	numLines := strings.Count(formatUsageText(), "\n")
	if res, exp := numLines, 10; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
}
