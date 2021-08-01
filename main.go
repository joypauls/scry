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
	var line string
	if selected {
		line = fmt.Sprintf("|%-5s|%-5d|%-9s|%c %s\n",
			fs.Label,
			fs.Time.Year(),
			fs.SizePretty,
			arrowRight,
			fs.Name,
		)
	} else {
		line = fmt.Sprintf("|%-5s|%-5d|%-9s|%s\n",
			fs.Label,
			fs.Time.Year(),
			fs.SizePretty,
			fs.Name,
		)
	}
	termboxPrint(x, y, fg, bg, line)
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

// var edit_box EditBox

// const edit_box_width = 30

// const grid_cols = 5 // not character width, grid cell width
// const grid_rows = 5
// const grid_cell_width = 7

// var files [grid_rows]string

// func fillWithJunk(files []string) {
// 	for i := 0; i < grid_rows; i++ {
// 		files[i] = fmt.Sprintf("file%d", i)
// 	}
// }

// // left oriented
// func draw_test_grid(xStart int, yStart int, coldef termbox.Attribute) {
// 	fillWithJunk(files[:])
// 	// termbox.SetCell(xStart, yStart, '│', coldef, coldef)
// 	for i := 0; i < grid_rows; i++ {
// 		for j := 0; j < grid_cols; j++ {
// 			// iterate across x axis
// 			left_side := (grid_cell_width * j) + xStart
// 			// termbox.SetCell(left_side, yStart+i, 'O', coldef, coldef)
// 			formatter := fmt.Sprintf("%%-%ds", grid_cell_width)
// 			if j == grid_cols-1 {
// 				termboxPrint(left_side, yStart+i, coldef, coldef, fmt.Sprintf(formatter, files[i]))
// 			} else {
// 				termboxPrint(left_side, yStart+i, coldef, coldef, fmt.Sprintf(formatter, "0"))
// 			}
// 		}
// 	}
// }

// the current selected index in the list
// needs to be bounded by the current size of array of files
var curIndex = 0

// starting upper left corner of canvas
var xStart int = 0
var yStart int = 0

var xGridStart int = 0
var yGridStart int = 2

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

// // This should move the marker in the *backing data structure*.
// // These coordinates need not reflect the termbox cells displayed.
// func moveMarker(x_change int, y_change int) {
// 	xMarker = minInt(maxInt(xMarker+x_change, 1), grid_cols)
// 	yMarker = minInt(maxInt(yMarker+y_change, 1), grid_rows)
// }

// // pass virtual coordinates, and place in termbox space
// func placeMarker(x int, y int, coldef termbox.Attribute) {
// 	if x == grid_cols {
// 		formatter := fmt.Sprintf("%c %%-%ds", arrowRight, grid_cell_width)
// 		termboxPrint(
// 			(x-1)*grid_cell_width+xGridStart,
// 			(y-1)+yGridStart,
// 			coldef,
// 			coldef,
// 			fmt.Sprintf(formatter, files[y-1]),
// 		)
// 	} else {
// 		termbox.SetCell(
// 			(x-1)*grid_cell_width+xGridStart,
// 			(y-1)+yGridStart,
// 			'X',
// 			coldef,
// 			coldef,
// 		)
// 	}
// }

// Handles drawing on the screen, hydrating grid with current state.
func redraw() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	// setting starting point of main screen object
	_, height := termbox.Size()
	// x_end := width - 1
	y_end := height - 1

	// draw_test_grid(xGridStart, yGridStart, coldef)
	// placeMarker(xMarker, yMarker, coldef)

	var dir Directory
	dir.Read()

	// draw top menu bar
	termboxPrint(xStart, yStart, coldef, coldef, dir.path)

	// draw files
	for i, f := range dir.files {
		displayFile(xGridStart, yGridStart+i, coldef, coldef, f, i == curIndex)
	}

	// draw bottom menu bar
	// coordStr := fmt.Sprintf("(%d,%d)", xMarker, yMarker)
	// termboxPrint(x_end-len(coordStr)+1, y_end, coldef, coldef, coordStr)
	termboxPrint(xGridStart, y_end, coldef, coldef, "Press: ESC/CTRL+c (quit), h (help)")

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

	redraw()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break loop
				// case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				// 	// edit_box.MoveCursorOneRuneBackward()
				// 	moveMarker(-1, 0)
				// case termbox.KeyArrowRight, termbox.KeyCtrlF:
				// 	// edit_box.MoveCursorOneRuneForward()
				// 	moveMarker(1, 0)
				// case termbox.KeyArrowDown:
				// 	moveMarker(0, 1)
				// case termbox.KeyArrowUp:
				// 	moveMarker(0, -1)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw()
	}
}
