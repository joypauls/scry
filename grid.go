// This package is handling the printing, terminal functionality, and user input.
package main

// inspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

import (
	"fmt"
	"unicode/utf8"

	ft "github.com/joypauls/scry/filetools"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// Uses termbox specific functionality to display a string in cells.
func termboxPrint(x, y int, fg, bg termbox.Attribute, s string) {
	for _, c := range s {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func rune_advance_len(r rune, pos int) int {
	if r == '\t' {
		return tabstop_length - pos%tabstop_length
	}
	return runewidth.RuneWidth(r)
}

func voffset_coffset(text []byte, boffset int) (voffset, coffset int) {
	text = text[:boffset]
	for len(text) > 0 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		coffset += 1
		voffset += rune_advance_len(r, voffset)
	}
	return
}

func byte_slice_grow(s []byte, desired_cap int) []byte {
	if cap(s) < desired_cap {
		ns := make([]byte, len(s), desired_cap)
		copy(ns, s)
		return ns
	}
	return s
}

func byte_slice_remove(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func byte_slice_insert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byte_slice_grow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

const preferred_horizontal_threshold = 5
const tabstop_length = 8

type EditBox struct {
	text           []byte
	line_voffset   int
	cursor_boffset int // cursor offset in bytes
	cursor_voffset int // visual cursor offset in termbox cells
	cursor_coffset int // cursor offset in unicode code points
}

// Draws the EditBox in the given location, 'h' is not used at the moment
func (eb *EditBox) Draw(x, y, w, h int) {
	eb.AdjustVOffset(w)

	const coldef = termbox.ColorDefault
	const colred = termbox.ColorRed

	fill(x, y, w, h, termbox.Cell{Ch: ' '})

	t := eb.text
	lx := 0
	tabstop := 0
	for {
		rx := lx - eb.line_voffset
		if len(t) == 0 {
			break
		}

		if lx == tabstop {
			tabstop += tabstop_length
		}

		if rx >= w {
			termbox.SetCell(x+w-1, y, arrowRight,
				colred, coldef)
			break
		}

		r, size := utf8.DecodeRune(t)
		if r == '\t' {
			for ; lx < tabstop; lx++ {
				rx = lx - eb.line_voffset
				if rx >= w {
					goto next
				}

				if rx >= 0 {
					termbox.SetCell(x+rx, y, ' ', coldef, coldef)
				}
			}
		} else {
			if rx >= 0 {
				termbox.SetCell(x+rx, y, r, coldef, coldef)
			}
			lx += runewidth.RuneWidth(r)
		}
	next:
		t = t[size:]
	}

	if eb.line_voffset != 0 {
		termbox.SetCell(x, y, arrowLeft, colred, coldef)
	}
}

// Adjusts line visual offset to a proper value depending on width
func (eb *EditBox) AdjustVOffset(width int) {
	ht := preferred_horizontal_threshold
	max_h_threshold := (width - 1) / 2
	if ht > max_h_threshold {
		ht = max_h_threshold
	}

	threshold := width - 1
	if eb.line_voffset != 0 {
		threshold = width - ht
	}
	if eb.cursor_voffset-eb.line_voffset >= threshold {
		eb.line_voffset = eb.cursor_voffset + (ht - width + 1)
	}

	if eb.line_voffset != 0 && eb.cursor_voffset-eb.line_voffset < ht {
		eb.line_voffset = eb.cursor_voffset - ht
		if eb.line_voffset < 0 {
			eb.line_voffset = 0
		}
	}
}

func (eb *EditBox) MoveCursorTo(boffset int) {
	eb.cursor_boffset = boffset
	eb.cursor_voffset, eb.cursor_coffset = voffset_coffset(eb.text, boffset)
}

func (eb *EditBox) RuneUnderCursor() (rune, int) {
	return utf8.DecodeRune(eb.text[eb.cursor_boffset:])
}

func (eb *EditBox) RuneBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(eb.text[:eb.cursor_boffset])
}

func (eb *EditBox) MoveCursorOneRuneBackward() {
	if eb.cursor_boffset == 0 {
		return
	}
	_, size := eb.RuneBeforeCursor()
	eb.MoveCursorTo(eb.cursor_boffset - size)
}

func (eb *EditBox) MoveCursorOneRuneForward() {
	// remove this to do something
	// if eb.cursor_boffset == len(eb.text) {
	// 	return
	// }
	_, size := eb.RuneUnderCursor()
	eb.MoveCursorTo(eb.cursor_boffset + size)
}

func (eb *EditBox) MoveCursorToBeginningOfTheLine() {
	eb.MoveCursorTo(0)
}

func (eb *EditBox) MoveCursorToEndOfTheLine() {
	eb.MoveCursorTo(len(eb.text))
}

func (eb *EditBox) DeleteRuneBackward() {
	if eb.cursor_boffset == 0 {
		return
	}

	eb.MoveCursorOneRuneBackward()
	_, size := eb.RuneUnderCursor()
	eb.text = byte_slice_remove(eb.text, eb.cursor_boffset, eb.cursor_boffset+size)
}

func (eb *EditBox) DeleteRuneForward() {
	if eb.cursor_boffset == len(eb.text) {
		return
	}
	_, size := eb.RuneUnderCursor()
	eb.text = byte_slice_remove(eb.text, eb.cursor_boffset, eb.cursor_boffset+size)
}

func (eb *EditBox) DeleteTheRestOfTheLine() {
	eb.text = eb.text[:eb.cursor_boffset]
}

func (eb *EditBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	eb.text = byte_slice_insert(eb.text, eb.cursor_boffset, buf[:n])
	eb.MoveCursorOneRuneForward()
}

// Please, keep in mind that cursor depends on the value of line_voffset, which
// is being set on Draw() call, so.. call this method after Draw() one.
func (eb *EditBox) CursorX() int {
	return eb.cursor_voffset - eb.line_voffset
}

////////////////////
// State Tracking //
////////////////////

//////////////////
// Main Program //
//////////////////

var edit_box EditBox

const edit_box_width = 30

const grid_cols = 5 // not character width, grid cell width
const grid_rows = 5
const grid_cell_width = 7

var files [grid_rows]string

func fillWithJunk(files []string) {
	for i := 0; i < grid_rows; i++ {
		files[i] = fmt.Sprintf("file%d", i)
	}
}

// left oriented
func draw_test_grid(x_start int, y_start int, coldef termbox.Attribute) {
	fillWithJunk(files[:])
	// termbox.SetCell(x_start, y_start, '│', coldef, coldef)
	for i := 0; i < grid_rows; i++ {
		for j := 0; j < grid_cols; j++ {
			// iterate across x axis
			left_side := (grid_cell_width * j) + x_start
			// termbox.SetCell(left_side, y_start+i, 'O', coldef, coldef)
			formatter := fmt.Sprintf("%%-%ds", grid_cell_width)
			if j == grid_cols-1 {
				termboxPrint(left_side, y_start+i, coldef, coldef, fmt.Sprintf(formatter, files[i]))
			} else {
				termboxPrint(left_side, y_start+i, coldef, coldef, fmt.Sprintf(formatter, "0"))
			}
		}
	}
}

// state of selection
var x_marker int = 1
var y_marker int = 1

// starting upper left corner of canvas
var x_start int = 0
var y_start int = 0

var xGridStart int = 0
var yGridStart int = 2

var arrowLeft = '←'
var arrowRight = '→'

func init() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

func min_int(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This should move the marker in the *backing data structure*.
// These coordinates need not reflect the termbox cells displayed.
func move_marker(x_change int, y_change int) {
	x_marker = min_int(max_int(x_marker+x_change, 1), grid_cols)
	y_marker = min_int(max_int(y_marker+y_change, 1), grid_rows)
}

// pass virtual coordinates, and place in termbox space
func place_marker(x int, y int, coldef termbox.Attribute) {
	if x == grid_cols {
		formatter := fmt.Sprintf("%c %%-%ds", arrowRight, grid_cell_width)
		termboxPrint(
			(x-1)*grid_cell_width+xGridStart,
			(y-1)+yGridStart,
			coldef,
			coldef,
			fmt.Sprintf(formatter, files[y-1]),
		)
	} else {
		termbox.SetCell(
			(x-1)*grid_cell_width+xGridStart,
			(y-1)+yGridStart,
			'X',
			coldef,
			coldef,
		)
	}
}

// Handles drawing on the screen, hydrating grid with current state.
func redraw() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	// setting starting point of main screen object
	// width and height
	width, height := termbox.Size()
	x_end := width - 1
	y_end := height - 1

	// unicode box drawing chars around the edit box
	draw_test_grid(xGridStart, yGridStart, coldef)

	// draw the dynamic content dependent on user input
	// edit_box.Draw(x_start, y_start, edit_box_width, 1)
	place_marker(x_marker, y_marker, coldef)

	// draw top menu bar
	curDir := ft.GetCurDir()
	termboxPrint(x_start, y_start, coldef, coldef, curDir)

	// draw bottom menu bar
	coordStr := fmt.Sprintf("(%d,%d)", x_marker, y_marker)
	termboxPrint(x_end-len(coordStr)+1, y_end, coldef, coldef, coordStr)
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
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break mainloop
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				// edit_box.MoveCursorOneRuneBackward()
				move_marker(-1, 0)
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				// edit_box.MoveCursorOneRuneForward()
				move_marker(1, 0)
			case termbox.KeyArrowDown:
				move_marker(0, 1)
			case termbox.KeyArrowUp:
				move_marker(0, -1)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw()
	}
}
