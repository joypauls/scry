/* Dealing with complex formatting for display */
package app

import (
	"fmt"
	fp "path/filepath"
	"regexp"

	"github.com/joypauls/scry/fst"
)

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
