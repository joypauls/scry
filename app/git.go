/* Parsing git stuff */
package app

import (
	"os/exec"
	fp "path/filepath"
	"strings"

	fst "github.com/joypauls/scry/fst"
)

// Tries to get the path to the nearest .git directory ancestor for the given Path.
func getGitPath(p *fst.Path) (*fst.Path, error) {
	// either returns path, ".git" if current dir, else an error
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = p.String()
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	sOut := strings.TrimSuffix(string(out), "\n")
	if sOut == ".git" {
		return fst.NewPath(fp.Join(p.String(), sOut)), nil
	}
	return fst.NewPath(sOut), nil
}

func getGitBranchName(p *fst.Path) (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = p.String()
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	sOut := strings.TrimSuffix(string(out), "\n")
	return sOut, err
}
