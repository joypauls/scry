package main

import "github.com/joypauls/scry/fst"

// Implementing Item interface from bubbles
// https://pkg.go.dev/github.com/charmbracelet/bubbles@v0.13.0/list#Item
// --------------------------------------------------------
// type Item interface {
// 	// Filter value is the value we use when filtering against this item when
// 	// we're filtering the list.
// 	FilterValue() string
// }
// --------------------------------------------------------
//
// Should this use the composition pattern?
type FileItem struct {
	fst.File
}

// All these accessors are defined by default delegate
// https://pkg.go.dev/github.com/charmbracelet/bubbles@v0.13.0/list#NewDefaultDelegate
func (fi FileItem) FilterValue() string { return fi.Name() }

func MakeFileItem(f fst.File) FileItem {
	fi := FileItem{f}
	return fi
}
