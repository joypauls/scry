// should store formatted strings so they don't need to be processed each refresh
package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
)

// Uses tcell specific functionality to display a string in cells.
func draw(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for _, r := range []rune(text) {
		s.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
}

func drawFile(s tcell.Screen, x, y int, selected bool, f fst.File, p fst.Path) {
	style := defStyle
	if selected {
		style = selStyle
	}
	draw(s, x, y, style, formatFile(f, p))
}
