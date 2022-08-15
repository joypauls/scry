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

func drawFile(s tcell.Screen, x1, x2, y int, selected bool, f fst.File, p *fst.Path) {
	style := theme.Default
	if selected {
		style = theme.Selected
	}
	draw(s, x1, y, style, formatFile(f, p, x2-x1))
}

// draw stuff that is not directory contents
func drawFrame(s tcell.Screen, app *App) {
	maxHeaderLen := (7 * app.width) / 10 // 70% of width
	// top bar content
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

	branch, _ := getGitBranchName(app.Path)
	if branch != "" {
		fmtStr = "%" + strconv.Itoa(app.width-maxHeaderLen) + "s"
		draw(s, maxHeaderLen, 0, theme.Highlight, fmt.Sprintf(fmtStr, "git:("+branch+")"))
	}

	// // bottom line
	// coordStr := fmt.Sprintf("(%d)", app.index)
	// draw(s, app.xEnd-len(coordStr)+1, app.height-1, defStyle, coordStr)
	fmtStr = "%" + strconv.Itoa(app.width) + "s"
	draw(s, 0, app.height-1, theme.Highlight, fmt.Sprintf(fmtStr, "[esc]quit  [h]home  [b]initial"))
}

func drawDirectoryContents(s tcell.Screen, app *App) {
	limit := minInt(app.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			s,
			app.xStart,
			app.middle,
			app.yStart+i,
			app.index == i,
			app.File(i+app.offset),
			app.Path,
		)
	}
}

func drawDivider(s tcell.Screen, x, y1, y2 int, style tcell.Style) {
	for i := y1; i <= y2; i++ {
		draw(s, x, i, style, fmt.Sprintf("%c", vLineRune))
	}
}

func drawSelectionDetails(s tcell.Screen, x, y int, f fst.File) {
	draw(s, x, y, theme.DefaultEmph, fmt.Sprintf(
		"%-20s",
		formatFileName(f),
	))
	draw(s, x, y+2, theme.Default, fmt.Sprintf(
		"Size          : %s",
		f.Size.String(),
	))
	draw(s, x, y+3, theme.Default, fmt.Sprintf(
		"Last Modified : %s",
		fmt.Sprintf("%2d/%02d/%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100),
	))
	draw(s, x, y+4, theme.Default, fmt.Sprintf(
		"Permissions   : %#-4o",
		f.Perm,
	))
	draw(s, x, y+5, theme.Default, fmt.Sprintf(
		"                %s",
		f.Perm,
	))
}

// Actual file contents
func drawWindow(s tcell.Screen, app *App) {
	// handle case when we need overflow indicators
	if !app.IsEmpty() {
		if app.offset > 0 {
			draw(s, (app.middle-app.innerPadding)/2, 1, theme.Default, fmt.Sprintf("%c", arrowUp))
		}
		if app.maxIndex+app.offset+1 < app.Size() {
			draw(s, (app.middle-app.innerPadding)/2, app.height-2, theme.Default, fmt.Sprintf("%c", arrowDown))
		}
	}
	// draw main content
	drawDirectoryContents(s, app)
	drawDivider(s, app.middle, 1, app.height-2, theme.Default)
	f := app.File(app.Index() + app.offset)
	drawSelectionDetails(s, app.middle+app.innerPadding, app.yStart, f)
}
