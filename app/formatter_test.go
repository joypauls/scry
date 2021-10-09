package app

import (
	"testing"

	"github.com/joypauls/scry/fst"
)

func TestFormatPathRegular(t *testing.T) {
	// table of test cases
	tables := []struct {
		pathString string
		maxLen     int
		expected   string // after formatting
	}{
		{"/", 100, "/"},
		{"/test", 100, "/test/"},
		{"/test/", 100, "/test/"},
		{"/test/more/stuff", 100, "/test/more/stuff/"},
	}

	// iterate over test tables
	for _, table := range tables {
		result := formatPath(fst.NewPath(table.pathString), table.maxLen)
		if result != table.expected {
			t.Errorf("Result: %s, Expected: %s", result, table.expected)
		}
	}
}

func TestFormatPathWeird(t *testing.T) {
	// table of test cases
	tables := []struct {
		pathString string
		maxLen     int
		expected   string // after formatting
	}{
		{"/test/cutoff/last/part", 10, "/.../part/"},
	}

	// iterate over test tables
	for _, table := range tables {
		result := formatPath(fst.NewPath(table.pathString), table.maxLen)
		if result != table.expected {
			t.Errorf("Result: %s, Expected: %s", result, table.expected)
		}
	}
}
