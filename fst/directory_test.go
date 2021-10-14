package fst

import (
	"log"
	"testing"
	"testing/fstest"
	"time"
)

var sysValue int

// Emulating https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/io/fs/readdir_test.go
var testFs = fstest.MapFS{
	"hello.txt": {
		Data:    []byte("hello, world"),
		Mode:    0456,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
	// MapFS implicitly adds the directory file "sub" for us
	"sub/goodbye.txt": {
		Data:    []byte("goodbye, world"),
		Mode:    0456,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
}

// func createDirEntrySlice(fs fstest.MapFS, path string) []os.DirEntry {
// 	fi, err := Stat(testFs, test.path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dirEntry := FileInfoToDirEntry(fi)
// }

// testFs := fstest.MapFS{
// 	"notadir.txt": {
// 		Data:    []byte("hello, world"),
// 		Mode:    0,
// 		ModTime: time.Now(),
// 		Sys:     &sysValue,
// 	},
// 	"adir": {
// 		Data:    nil,
// 		Mode:    os.ModeDir,
// 		ModTime: time.Now(),
// 		Sys:     &sysValue,
// 	},
// }

func TestProcessDirectory(t *testing.T) {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	dirProcessed := processDirectory(dirRaw, true)

	result := len(dirProcessed)
	expected := 2
	if result != expected {
		t.Errorf("Result: %d, Wanted: %d", result, expected)
	}
}

func TestDirectory(t *testing.T) {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	dirProcessed := processDirectory(dirRaw, true)
	// kinda hacky, should refactor d.Read() so that it works on FS interface to use MapFS
	d := Directory{
		files: dirProcessed,
		size:  len(dirProcessed),
	}
	d.SortNameDesc() // standardize

	if res, exp := d.Size(), 2; res != exp {
		t.Errorf("Result: %d, Wanted: %d", res, exp)
	}

	if res, exp := d.IsEmpty(), false; res != exp {
		t.Errorf("Result: %t, Wanted: %t", res, exp)
	}

	// test index-based file selection
	f := d.File(0)
	if res, exp := f.Name, "hello.txt"; res != exp {
		t.Errorf("Result: %s, Wanted: %s", res, exp)
	}

	// test alternative for previous
	files := d.Files()
	if res, exp := files[1].Name, "sub"; res != exp {
		t.Errorf("Result: %s, Wanted: %s", res, exp)
	}
}
