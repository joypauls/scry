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
	".env": {
		Data:    []byte("hello, world"),
		Mode:    0775,
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
	"model.py": {
		Data:    []byte("hello, world"),
		Mode:    0775,
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

	// no dot files
	dirProcessed := processDirectory(dirRaw, false)
	res := len(dirProcessed)
	exp := 3
	if res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
	// show dot files
	dirProcessed = processDirectory(dirRaw, true)
	res = len(dirProcessed)
	exp = 4
	if res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
}

func TestDirectory(t *testing.T) {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	d := NewDirectoryFromSlice(dirRaw, false)

	if res, exp := d.Size(), 3; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	if res, exp := d.IsEmpty(), false; res != exp {
		t.Errorf("Result: %t, Expected: %t", res, exp)
	}

	if res, exp := d.IsProblem(), false; res != exp {
		t.Errorf("Result: %t, Expected: %t", res, exp)
	}

	// test index-based file selection
	f := d.File(0)
	if res, exp := f.Name, "hello.txt"; res != exp {
		t.Errorf("Result: %s, Expected: %s", res, exp)
	}

	// test grabbing all files
	files := d.Files()
	if res, exp := files[0].Name, "hello.txt"; res != exp {
		t.Errorf("Result: %s, Expected: %s", res, exp)
	}
}

func TestDirectorySorting(t *testing.T) {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	d := new(Directory)
	d.files = processDirectory(dirRaw, false)
	d.size = len(d.files)
	d.SortNameDesc()

	if res, exp := d.File(0).Name, "hello.txt"; res != exp {
		t.Errorf("Result: %s, Expected: %s", res, exp)
	}

	if res, exp := d.File(2).Name, "sub"; res != exp {
		t.Errorf("Result: %s, Expected: %s", res, exp)
	}
}
