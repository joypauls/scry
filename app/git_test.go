package app

import (
	"testing"

	"github.com/joypauls/scry/fst"
)

// func TestGetGitPath(t *testing.T) {
// 	// cur, err := os.Getwd()
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// table of test cases
// 	tables := []struct {
// 		path     string
// 		expected string // after formatting
// 	}{
// 		{".", "/Users/joypauls/Documents/code/file-scry/.git"},
// 		{"/Users/joypauls/Documents/code/file-scry", "/Users/joypauls/Documents/code/file-scry/.git"},
// 	}

// 	// iterate over test tables
// 	for _, table := range tables {
// 		result, _ := getGitPath(fst.NewPath(table.path))
// 		if result.String() != table.expected {
// 			t.Errorf("Result: %s, Expected: %s", result, table.expected)
// 		}
// 	}
// }

func TestGetGitBranch(t *testing.T) {
	// table of test cases
	tables := []struct {
		path     string
		expected string // after formatting
	}{
		{".", "main"}, // this is dumb and will break lol
	}

	// iterate over test tables
	for _, table := range tables {
		result, _ := getGitBranchName(fst.NewPath(table.path))
		// this is a shitty check but no way to know correct branch?
		if result == "" {
			t.Errorf("Result: %s, Expected: %s", result, table.expected)
		}
	}
}
