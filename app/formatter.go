/* Dealing with complex formatting for display */
package app

import (
	"fmt"
	"os"
	fp "path/filepath"
	"regexp"

	"github.com/joypauls/scry/fst"
)

// symbols for file
const dirLabel = "📁"
const fileLabel = ""

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

func formatFileName(f fst.File) string {
	if f.IsDir {
		return dirLabel + " " + f.Name
	}
	return f.Name
}

func formatFile(f fst.File, p *fst.Path) string {
	format := "%-30.30s  %9s  %8s"
	if f.IsDir {
		// very hacky way to accomodate double width rune
		format = "%-29.29s  %9s  %8s"
	}
	name := formatFileName(f)
	// check for symlink
	if f.IsSymLink {
		// this should be done when the file is read
		target, err := os.Readlink(fp.Join(p.String(), name))
		if err != nil {
			target = "?"
		}
		name = name + fmt.Sprintf(" %c %s", arrowRight, target)
	}
	// return fmt.Sprintf("%s %-4s  %#-4o  %-9s  %s ",
	// 	label,
	// 	fmt.Sprintf("%02d-%02d-%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100),
	// 	f.Perm,
	// 	f.Size.String(),
	// 	name,
	// )
	return fmt.Sprintf(format,
		name,
		f.Size.String(),
		fmt.Sprintf("%2d/%02d/%d", f.Time.Month(), f.Time.Day(), f.Time.Year()%100),
	)
}
