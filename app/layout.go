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
	topPad       int // min val for this is 2
	bottomPad    int // min val for this is 2
	windowHeight int // equals the # of available lines
}

// generator func for Layout
func NewLayout() *Layout {
	f := new(Layout)
	f.width, f.height = termbox.Size()
	f.xStart = 0
	f.xEnd = f.width - 1
	f.topPad = 2
	f.bottomPad = 2
	f.yStart = 0 + f.topPad
	f.yEnd = f.height - 1 - f.bottomPad
	f.windowHeight = f.yEnd - f.yStart + 1
	return f
}
