/* Dealing with complex formatting for display */
package main

import (
	"fmt"
	fp "path/filepath"
	"regexp"

	"github.com/joypauls/scry/fst"
)

// symbols for file
const dirLabel = "🗂 " // "📁 "
const fileLabel = ""

var arrowLeft = '←'
var arrowRight = '→'
var arrowUp = '▲'
var arrowDown = '▼'

// TODO
// // initialize one time display-related configs at program start
// // this could probably be a configuration struct
// // after everything is declared
// func init() {
// 	if runewidth.EastAsianWidth {
// 		arrowLeft = '<'
// 		arrowRight = '>'
// 	}
// }

// minimum for maxLength is 5 (/... leading, / trailing), enforce?
func formatPath(p *fst.Path, maxLen int) string {
	if len(p.String()) == 1 {
		return p.String()
	} else if len(p.String()) < maxLen {
		return p.String() + string(fp.Separator)
	}
	// iterate until the path is shortened enough
	re := regexp.MustCompile(`^\/[^\/]*`)
	clipped := p.String()
	for len(clipped)+5 > maxLen {
		clipped = re.ReplaceAllString(clipped, "")
		// eventually could terminate to just "/'" if last node is > maxLength?
	}
	return fmt.Sprintf("/...%s%c", clipped, fp.Separator)
}

func formatHeader(p *fst.Path, maxLen int) string {
	useEmoji := true
	maxHeaderLen := (7 * maxLen) / 10 // 70% of width
	header := formatPath(p, maxHeaderLen)
	if useEmoji {
		header = "🔮 " + header
	}
	return header
}

func formatFileName(name string, isDir bool) string {
	if isDir {
		return dirLabel + " " + name
	}
	return name
}

// this method needs some tlc
func formatFile(name string, size string, isDir bool, isSymLink bool, symLinkTarget string, w int) string {
	// shouldn't be hardcoded but don't know a great way yet
	// 1 + nameWidth + statsWidth = w
	statsWidth := 24
	nameWidth := w - statsWidth
	format := fmt.Sprintf("%%-%d.%ds  %%9s", nameWidth, nameWidth)
	if isDir {
		// very hacky way to accomodate double width rune
		format = fmt.Sprintf("%%-%d.%ds  %%9s", nameWidth, nameWidth)
	}
	namePretty := formatFileName(name, isDir)
	// check for symlink
	if isSymLink {
		namePretty = namePretty + fmt.Sprintf(" %c %s", arrowRight, symLinkTarget)
	}

	return fmt.Sprintf(format,
		namePretty,
		size,
	)
}
