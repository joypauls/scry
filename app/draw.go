// should store formatted strings so they don't need to be processed each refresh
package app

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
)

var vLineRune = '\u2502'

// Uses tcell specific functionality to display a string in cells.
func draw(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for _, r := range []rune(text) {
		s.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
}

func drawDivider(s tcell.Screen, x, y1, y2 int, style tcell.Style) {
	for i := y1; i <= y2; i++ {
		draw(s, x, i, style, fmt.Sprintf("%c", vLineRune))
	}
}

func drawFile(s tcell.Screen, x, y int, selected bool, f fst.File, p *fst.Path) {
	style := theme.Default
	if selected {
		style = theme.Selected
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
	// user, err := user.Current()
	// if err != nil {
	// 	header = fmt.Sprintf(header, "?")
	// } else {
	// 	header = fmt.Sprintf(header, user.Username)
	// }
	fmtStr := "%-" + strconv.Itoa(app.width) + "s"
	draw(s, 0, 0, theme.Highlight, fmt.Sprintf(fmtStr, header))
	// draw(s, 0, 0, selStyle, fmt.Sprintf(fmtStr, "[esc] quit [h] home [b] initial"))

	// // bottom line
	// coordStr := fmt.Sprintf("(%d)", app.index)
	// draw(s, app.xEnd-len(coordStr)+1, app.height-1, defStyle, coordStr)
	fmtStr = "%" + strconv.Itoa(app.width) + "s"
	draw(s, 0, app.height-1, theme.Highlight, fmt.Sprintf(fmtStr, "[esc] quit [h] home [b] initial"))
}

// Actual file contents
func drawWindow(s tcell.Screen, app *App) {
	if !app.IsEmpty() {
		if app.offset > 0 {
			draw(s, 24, 1, theme.Default, fmt.Sprintf("%c", arrowUp))
		}
		if app.maxIndex+app.offset+1 < app.Size() {
			draw(s, 24, app.height-2, theme.Default, fmt.Sprintf("%c", arrowDown))
		}
	}
	limit := minInt(app.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			s,
			// app.xStart,
			0,
			app.yStart+i,
			i == app.index,
			app.File(i+app.offset),
			app.Path,
		)
	}
	drawDivider(s, 54, 1, app.height-2, theme.Default)

	f := app.File(app.Index() + app.offset)
	draw(s, 57, app.yStart, theme.Default, fmt.Sprintf("%-20s", formatFileName(f)))
	draw(s, 57, app.yStart+2, theme.Default, fmt.Sprintf("Size           %s", f.Size.String()))
	draw(s, 57, app.yStart+3, theme.Default, fmt.Sprintf("Last Modified  %s", fmt.Sprintf("%2d/%02d/%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100)))
	draw(s, 57, app.yStart+4, theme.Default, fmt.Sprintf("Permissions    %#-4o", f.Perm))
	draw(s, 57, app.yStart+5, theme.Default, fmt.Sprintf("               %s", f.Perm))
}
