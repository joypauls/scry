package app

import (
	"testing"

	"github.com/joypauls/scry/fst"
)

// func TestFormatPathRegular(t *testing.T) {
// 	// table of test cases
// 	tables := []struct {
// 		pathString string
// 		maxLen     int
// 		expected   string // after formatting
// 	}{
// 		{"/", 100, "/"},
// 		{"/test", 100, "/test/"},
// 		{"/test/", 100, "/test/"},
// 		{"/test/more/stuff", 100, "/test/more/stuff/"},
// 	}

// 	// iterate over test tables
// 	for _, table := range tables {
// 		result := formatPath(fst.NewPath(table.pathString), table.maxLen)
// 		if result != table.expected {
// 			t.Errorf("Result: %s, Expected: %s", result, table.expected)
// 		}
// 	}
// }

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

// func TestFormatFileRegular(t *testing.T) {
// 	testFile := fst.File{
// 		Name:      "test.go",
// 		Size:      fst.BytesSI(12345),
// 		Time:      time.Date(2021, time.October, 1, 12, 0, 0, 0, time.UTC),
// 		IsDir:     false,
// 		IsReg:     true,
// 		IsSymLink: false,
// 		Perm:      0777,
// 	}
// 	testPath := fst.NewPath("/")
// 	result := formatFile(testFile, testPath)
// 	expected := fmt.Sprintf("%s 10-01-21  0777  12.3 KB    test.go ", fileLabel)
// 	if result != expected {
// 		t.Errorf("Result: %s, Expected: %s", result, expected)
// 	}
// }
