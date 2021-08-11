package app

import (
	"github.com/nsf/termbox-go"
)

// Managing the UI layout
type Layout struct {
	width     int
	height    int
	xEnd      int
	yEnd      int
	topPad    int
	bottomPad int
}

// generator func for Layout
func NewLayout() *Layout {
	f := new(Layout)
	f.width, f.height = termbox.Size()
	f.xEnd = f.width - 1
	f.yEnd = f.height - 1
	f.topPad = 2
	f.bottomPad = 2
	return f
}
