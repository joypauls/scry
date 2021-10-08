/* Dealing with complex formatting for display */
package app

import (
	fp "path/filepath"
	"regexp"

	"github.com/joypauls/scry/fst"
)

// minimum for maxLength is 5 (/... leading, / trailing), enforce?
func formatPath(p *fst.Path, maxLen int) string {
	if len(p.String())+1 < maxLen {
		return p.String() + string(fp.Separator)
	}
	// iterate until the path is shortened enough
	re := regexp.MustCompile(`^\/[^\/]*`)
	formatted := p.String()
	for len(formatted)+5 > maxLen {
		formatted = re.ReplaceAllString(formatted, "")
		// eventually could terminate to just "/'" if last node is > maxLength?
	}
	return "/..." + formatted + string(fp.Separator)
}
