// This package is handling the printing, terminal functionality, and user input.
package ui

// partially nspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

import (
	"fmt"
	"log"

	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault // termbox.Attribute
var arrowLeft = '←'
var arrowRight = '→'

// initialize one time display-related configs at program start
func config() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

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

// the current selected index in the list
// needs to be bounded by the current size of array of files
var curIndex = 0
var maxIndex = 0

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This should move the marker in the *backing data structure*.
// These coordinates need not reflect the termbox cells displayed.
func moveIndex(change int) {
	curIndex = minInt(maxInt(curIndex+change, 0), maxIndex)
}

func drawFrame(l *Layout, d *fst.Directory) {
	// top line
	draw(0, 0, coldef, coldef, d.Path)
	// bottom line
	coordStr := fmt.Sprintf("(%d)", curIndex)
	draw(l.xEnd-len(coordStr)+1, l.yEnd, coldef, coldef, coordStr)
	draw(0, l.yEnd, coldef, coldef, "[ESC] quit, [h] help")
}

func drawWindow(l *Layout, d *fst.Directory) {
	for i, f := range d.Files {
		drawFile(0, 0+l.topPad+i, i == curIndex, f)
	}
}

// Handles drawing on the screen, hydrating grid with current state.
func refresh(l *Layout, d *fst.Directory) {
	termbox.Clear(coldef, coldef) // reset

	maxIndex = len(d.Files) - 1 // update num files

	drawFrame(l, d)
	drawWindow(l, d) // main content

	termbox.Flush() // clean
}

// Main program loop and user interactions
func Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	config()

	// set the layout
	layout := NewLayout()
	// init in current directory
	curDir := fst.GetCurDir()
	d := fst.NewDirectory(curDir)

	// draw the UI for the first time
	refresh(layout, d)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break loop
			case termbox.KeyArrowDown:
				moveIndex(1)
			case termbox.KeyArrowUp:
				moveIndex(-1)
			}
		case termbox.EventError:
			log.Fatal(ev.Err) // os.Exit(1) follows
		}
		refresh(layout, d)
	}
}
