// This package is handling the printing, terminal functionality, and user input.
package main

// inspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

import (
	"strconv"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// Uses termbox specific functionality to display a string in cells.
func termbox_print(x, y int, fg, bg termbox.Attribute, s string) {
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
const grid_rows = 3
const grid_cell_width = 2

// func draw_box_outline(x_start int, y_start int, coldef termbox.Attribute) {
// 	termbox.SetCell(x_start-1, y_start, '│', coldef, coldef)
// 	termbox.SetCell(x_start+edit_box_width, y_start, '│', coldef, coldef)
// 	termbox.SetCell(x_start-1, y_start-1, '┌', coldef, coldef)
// 	termbox.SetCell(x_start-1, y_start+1, '└', coldef, coldef)
// 	termbox.SetCell(x_start+edit_box_width, y_start-1, '┐', coldef, coldef)
// 	termbox.SetCell(x_start+edit_box_width, y_start+1, '┘', coldef, coldef)
// 	fill(x_start, y_start-1, edit_box_width, 1, termbox.Cell{Ch: '─'})
// 	fill(x_start, y_start+1, edit_box_width, 1, termbox.Cell{Ch: '─'})
// }

// left oriented
func draw_test_grid(x_start int, y_start int, coldef termbox.Attribute) {
	// termbox.SetCell(x_start, y_start, '│', coldef, coldef)
	for i := 0; i < grid_rows; i++ {
		for j := 0; j < grid_cols; j++ {
			// iterate across x axis
			left_side := (grid_cell_width * j) + x_start
			termbox.SetCell(left_side, y_start+i, 'O', coldef, coldef)
			termbox.SetCell(left_side+1, y_start+i, ' ', coldef, coldef)
		}
	}
}

var x_marker int = 1
var y_marker int = 1

var x_start int = 1
var y_start int = 1

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
	x_marker = min_int(max_int(x_marker+x_change, x_start), grid_cols)
	y_marker = min_int(max_int(y_marker+y_change, y_start), grid_rows)
}

// pass virtual coordinates, and place in termbox space
func place_marker(x int, y int, coldef termbox.Attribute) {
	termbox.SetCell(
		(x-1)*grid_cell_width+x_start,
		(y-1)+y_start,
		'X',
		coldef,
		coldef,
	)
}

func redraw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	// setting starting point of main screen object
	// width and height
	_, h := termbox.Size()
	// using offset +1 at corner (0+1) for niceness
	// midy := h / 2
	y_end := h - 1

	// set marker initial state

	// unicode box drawing chars around the edit box
	// draw_box_outline(x_start, y_start, coldef)
	draw_test_grid(x_start, y_start, coldef)

	// if runewidth.EastAsianWidth {
	// 	termbox.SetCell(midx-1, midy, '|', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy, '|', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy-1, '+', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy+1, '+', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy-1, '+', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy+1, '+', coldef, coldef)
	// 	fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '-'})
	// 	fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '-'})
	// } else {
	// 	termbox.SetCell(midx-1, midy, '│', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy, '│', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy-1, '┐', coldef, coldef)
	// 	termbox.SetCell(midx+edit_box_width, midy+1, '┘', coldef, coldef)
	// 	fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '─'})
	// 	fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '─'})
	// }

	// draw the dynamic content dependent on user input
	// edit_box.Draw(x_start, y_start, edit_box_width, 1)
	place_marker(x_marker, y_marker, coldef)
	// termbox.SetCursor(x_start+edit_box.CursorX(), y_start)

	// print at bottom of box
	coords_str := "(" + strconv.Itoa(x_marker) + "," + strconv.Itoa(y_marker) + ")"
	termbox_print(x_start-1, y_end-5, coldef, coldef, coords_str)
	termbox_print(x_start-1, y_end, coldef, coldef, "Press: ESC/CTRL+c (quit), h (help)")
	termbox.Flush()
}

var arrowLeft = '←'
var arrowRight = '→'

func init() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	redraw_all()
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
			// case termbox.KeyBackspace, termbox.KeyBackspace2:
			// 	edit_box.DeleteRuneBackward()
			// case termbox.KeyDelete, termbox.KeyCtrlD:
			// 	edit_box.DeleteRuneForward()
			// case termbox.KeyTab:
			// 	edit_box.InsertRune('\t')
			// case termbox.KeySpace:
			// 	edit_box.InsertRune(' ')
			// case termbox.KeyCtrlK:
			// 	edit_box.DeleteTheRestOfTheLine()
			// case termbox.KeyHome, termbox.KeyCtrlA:
			// 	edit_box.MoveCursorToBeginningOfTheLine()
			// case termbox.KeyEnd, termbox.KeyCtrlE:
			// 	edit_box.MoveCursorToEndOfTheLine()
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}
