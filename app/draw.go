// should store formatted strings so they don't need to be processed each refresh
package app

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
)

// SI defined base for multiple byte units
const baseSI = 1000
const prefixSI = "kMGTPE"

// symbols for display
const dirLabel = "📁"
const fileLabel = "  "

// Converts an integer number of bytes to SI units.
func humanizeBytes(bytes int64) string {
	if bytes < baseSI {
		return fmt.Sprintf("%d B", bytes) // < 1kB
	}
	magnitude := int64(baseSI)
	maxExp := 0
	for i := bytes / baseSI; i >= baseSI; i /= baseSI {
		magnitude *= baseSI
		maxExp++
	}
	return fmt.Sprintf(
		"%.1f %cB",
		float64(bytes)/float64(magnitude), // want quotient to be float
		prefixSI[maxExp],
	)
}

// Uses tcell specific functionality to display a string in cells.
func draw(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for _, r := range []rune(text) {
		s.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
}

func drawFile(s tcell.Screen, x, y int, selected bool, f *fst.File) {
	style := defStyle
	if selected {
		style = selStyle
	}
	label := fileLabel
	if f.IsDir {
		label = dirLabel
	}
	line := fmt.Sprintf("%s  %-4s  %#-4o  %-9s  %s ",
		label,
		fmt.Sprintf("%02d-%02d-%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100),
		f.Perm,
		humanizeBytes(f.Size),
		f.Name,
	)
	draw(s, x, y, style, line)
}
