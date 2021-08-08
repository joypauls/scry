// This package is handling the printing, terminal functionality, and user input.
package main

// partially nspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

import (
	"fmt"

	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault // termbox.Attribute

////////////////////////////////
// Termbox Printing Utilities //
////////////////////////////////

// Uses termbox specific functionality to display a string in cells.
func termboxPrint(x, y int, fg, bg termbox.Attribute, s string) {
	for _, c := range s {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func displayFile(x, y int, selected bool, f *fst.File) {
	fg := coldef
	bg := coldef
	if selected {
		fg = termbox.ColorBlack
		bg = termbox.ColorCyan
	}
	line := fmt.Sprintf("%s  %-10s %-9s %s",
		f.Label,
		fmt.Sprintf("%02d-%02d-%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100),
		f.SizePretty,
		f.Name,
	)
	termboxPrint(x, y, fg, bg, line)
}

////////
// UI //
////////

// Managing the UI layout
type Frame struct {
	width     int
	height    int
	xEnd      int
	yEnd      int
	topPad    int
	bottomPad int
}

// generator func for Frame
func NewFrame() *Frame {
	f := new(Frame)
	f.width, f.height = termbox.Size()
	f.xEnd = f.width - 1
	f.yEnd = f.height - 1
	f.topPad = 2
	f.bottomPad = 2
	return f
}

//////////////////
// Main Program //
//////////////////

// the current selected index in the list
// needs to be bounded by the current size of array of files
var curIndex = 0
var maxIndex = 0

var arrowLeft = '←'
var arrowRight = '→'

// initialize one time configs at program start
func init() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

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

// Handles drawing on the screen, hydrating grid with current state.
func refresh(frame *Frame, d *fst.Directory) {
	termbox.Clear(coldef, coldef)

	maxIndex = len(d.Files) - 1 // update

	// draw top menu bar
	termboxPrint(0, 0, coldef, coldef, d.Path)
	// termboxPrint(0, 1, coldef, coldef, "-------------------------------------------------------")

	// draw files
	for i, f := range d.Files {
		displayFile(0, 0+frame.topPad+i, i == curIndex, f)
	}

	// draw bottom menu bar
	// termboxPrint(0, frame.yEnd-1, coldef, coldef, "-------------------------------------------------------")
	coordStr := fmt.Sprintf("(%d)", curIndex)
	termboxPrint(frame.xEnd-len(coordStr)+1, frame.yEnd, coldef, coldef, coordStr)
	termboxPrint(0, frame.yEnd, coldef, coldef, "[ESC] quit, [h] help")

	// cleanup
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// set the frame
	frame := NewFrame()
	// init in current directory
	curDir := fst.GetCurDir()
	d := fst.NewDirectory(curDir)

	// draw the UI for the first time
	refresh(frame, d)

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
			panic(ev.Err)
		}
		refresh(frame, d)
	}
}
