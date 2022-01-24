package app

import (
	"testing"

	"github.com/joypauls/scry/fst"
)

func TestGetGitPath(t *testing.T) {
	// cur, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// table of test cases
	tables := []struct {
		path     string
		expected string // after formatting
	}{
		{".", "/Users/joypauls/Documents/code/file-scry/.git\n"},
		{"/Users/joypauls/Documents/code/file-scry", "/Users/joypauls/Documents/code/file-scry/.git\n"},
		// {"/test/more/stuff", "/test/more/stuff/"},
	}

	// iterate over test tables
	for _, table := range tables {
		result, _ := getGitPath(fst.NewPath(table.path))
		if result != table.expected {
			t.Errorf("Result: %s, Expected: %s", result, table.expected)
		}
	}
}
