package main

import "github.com/joypauls/scry/fst"

// Implementing Item interface from bubbles
type fileItem struct {
	title, size string
}

func (i fileItem) Title() string       { return i.title }
func (i fileItem) Description() string { return i.size }
func (i fileItem) FilterValue() string { return i.title }

func MakeFileItem(f fst.File) fileItem {
	var i fileItem
	i.title = f.Name
	i.size = f.Size.String()

	return i
}
