// should store formatted strings so they don't need to be processed each refresh
package app

import (
	"fmt"

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

func drawFile(s tcell.Screen, x, y int, selected bool, f fst.File, p *fst.Path) {
	style := defStyle
	if selected {
		style = selStyle
	}
	draw(s, x, y, style, formatFile(f, p))
}

// draw stuff that is not directory contents
func drawFrame(s tcell.Screen, app *App) {
	// top bar content
	maxHeaderLen := (7 * app.width) / 10 // 70% of width
	header := formatPath(app.Path, maxHeaderLen)
	if app.UseEmoji {
		header = "🔮 " + header
	}
	draw(s, 0, 0, defStyle, header)

	if app.offset > 0 {
		draw(s, 0, 1, defStyle, fmt.Sprintf("%c", arrowUp))
	}
	if app.maxIndex+app.offset+1 < app.Size() {
		draw(s, 0, app.height-2, defStyle, fmt.Sprintf("%c", arrowDown))
	}
	// // bottom line
	// coordStr := fmt.Sprintf("(%d)", app.index)
	// draw(s, app.xEnd-len(coordStr)+1, app.height-1, defStyle, coordStr)
	draw(s, 0, app.height-1, defStyle, "[esc] quit [h] home [b] initial")
}

// Actual file contents
func drawWindow(s tcell.Screen, app *App) {
	limit := minInt(app.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			s,
			app.xStart,
			app.yStart+i,
			i == app.index,
			*app.File(i + app.offset),
			app.Path,
		)
	}
}
