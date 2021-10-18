package fst

import (
	"os"
	"testing"
)

// main testing of Path constructor
func TestPathConstructor(t *testing.T) {
	// setup
	wd, _ := os.Getwd()
	// test
	got := NewPath("")
	want := NewPath(wd)
	if *got != *want {
		t.Errorf("got %q, wanted %q", *got, *want)
	}

	// setup
	s := "/usr/src/app/"
	// test
	got1 := NewPath(s)
	want1 := Path{path: "/usr/src/app"}
	if *got1 != want1 {
		t.Errorf("got %q, wanted %q", *got1, want1)
	}
}

func TestPathSet(t *testing.T) {
	// setup
	s := "/usr/src/app"
	// test
	got := NewPath("/")
	got.Set(s)
	want := NewPath(s)
	if *got != *want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPathString(t *testing.T) {
	// test
	want := "/usr/src/app"
	got := NewPath(want).String()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPathCopy(t *testing.T) {
	// setup
	s := "/usr/src/app"
	// test
	got := NewPath(s)
	want := got.Copy()
	if *got != *want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPathParent(t *testing.T) {
	// setup
	s := "/usr/src/app"
	p := NewPath(s)
	// test
	got := p.Parent()
	want := "/usr/src"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPathToParent(t *testing.T) {
	// setup
	s := "/usr/src/app"
	p := NewPath(s)
	// test
	p.ToParent()
	got := p.String()
	want := "/usr/src"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
