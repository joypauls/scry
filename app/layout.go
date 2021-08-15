package app

import (
	"github.com/nsf/termbox-go"
)

// Managing the UI layout
type Layout struct {
	width        int
	height       int
	xStart       int
	xEnd         int
	yStart       int
	yEnd         int
	windowHeight int // equals the # of available lines
}

// constructor for Layout
func MakeLayout() Layout {
	padding := 2 // min for this is 2
	f := Layout{}
	f.width, f.height = termbox.Size()
	f.xStart = 0
	f.xEnd = f.width - 1
	f.yStart = 0 + padding
	f.yEnd = f.height - 1 - padding
	f.windowHeight = f.yEnd - f.yStart + 1
	return f
}
