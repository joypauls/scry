package testutils

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/gdamore/tcell/v2"
)

// used for fake file system
var sysValue int

// only supporting this character set for now
var charset = "UTF-8"

// var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

// Emulating https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/io/fs/readdir_test.go
var testFS = fstest.MapFS{
	"hello.txt": {
		Data:    []byte("hello, world"),
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
	".env": {
		Data:    []byte("hello, world"),
		Mode:    0456,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
	// MapFS implicitly adds the directory file "sub" for us
	"sub/goodbye.txt": {
		Data:    []byte("goodbye, world"),
		Mode:    0775,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
}

func GetTestFS() fstest.MapFS {
	return testFS
}

// func GetEmptyTestFS() fstest.MapFS {
// 	return emptyTestFS
// }

// Lifted from https://github.com/gdamore/tcell/blob/master/sim_test.go
func MakeTestScreen(t *testing.T) tcell.SimulationScreen {
	s := tcell.NewSimulationScreen(charset)
	if s == nil {
		t.Fatalf("Failed to get simulation screen")
	}
	if e := s.Init(); e != nil {
		t.Fatalf("Failed to initialize screen: %v", e)
	}
	return s
}
