package main

import "github.com/joypauls/scry/fst"

// Implementing Item interface from bubbles
type item struct {
	title, size string
}

// All these accessors are defined by default delegate
// https://pkg.go.dev/github.com/charmbracelet/bubbles@v0.13.0/list#NewDefaultDelegate
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.size }
func (i item) FilterValue() string { return i.title }

func MakeItem(f fst.File) item {
	var i item
	i.title = f.Name
	i.size = f.Size.String()

	return i
}
