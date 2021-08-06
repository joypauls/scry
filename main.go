// This package is handling the printing, terminal functionality, and user input.
package main

// partially nspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

import (
	"fmt"

	ft "github.com/joypauls/file-scry/filetools"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

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

func displayFile(x, y int, fg, bg termbox.Attribute, fs ft.FileStats, selected bool) {
	formatter := "%-3s %-10s %-9s"
	var line string
	if selected {
		line = fmt.Sprintf(formatter+"  %c  %s",
			fs.Label,
			fmt.Sprintf("%d-%d-%d", fs.Time.Month(), fs.Time.Day(), fs.Time.Year()),
			fs.SizePretty,
			arrowRight,
			fs.Name,
		)
	} else {
		line = fmt.Sprintf(formatter+" %s",
			fs.Label,
			fmt.Sprintf("%d-%d-%d", fs.Time.Month(), fs.Time.Day(), fs.Time.Year()),
			fs.SizePretty,
			fs.Name,
		)
	}
	termboxPrint(x, y, fg, bg, line)
}

////////
// UI //
////////

// Here we're managing the UI framing
type Frame struct {
	width     int
	height    int
	xEnd      int
	yEnd      int
	topPad    int
	bottomPad int
}

func (f *Frame) Init() {
	f.width, f.height = termbox.Size()
	f.xEnd = f.width - 1
	f.yEnd = f.height - 1
	f.topPad = 2
	f.bottomPad = 2
}

// generator func for Frame
func NewFrame() *Frame {
	f := new(Frame)
	f.Init()
	return f
}

/////////////////////
// Directory State //
/////////////////////

type Directory struct {
	path  string
	files []ft.FileStats
}

func (d *Directory) Read() {
	d.path = ft.GetCurDir() // this should be less limited
	d.files = ft.GetFiles(".")
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
func refresh(frame *Frame) {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	var dir Directory // do we need to redeclare??
	dir.Read()
	maxIndex = len(dir.files) - 1 // update

	// draw top menu bar
	termboxPrint(0, 0, coldef, coldef, dir.path)
	termboxPrint(0, 1, coldef, coldef, "-------------------------------------------------------")

	// draw files
	for i, f := range dir.files {
		displayFile(0, 0+frame.topPad+i, coldef, coldef, f, i == curIndex)
	}

	// draw bottom menu bar
	termboxPrint(0, frame.yEnd-1, coldef, coldef, "-------------------------------------------------------")
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

	frame := NewFrame()

	refresh(frame)

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
		refresh(frame)
	}
}
