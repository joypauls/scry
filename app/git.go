/* Parsing git stuff */
package app

import (
	"os/exec"
	fp "path/filepath"

	fst "github.com/joypauls/scry/fst"
)

// func getGitPath(p *fst.Path) (*fst.Path, error) {
func getGitPath(p *fst.Path) (string, error) {
	// either returns path, ".git" if current dir, else an error
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = p.String()
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	sOut := string(out)
	if sOut == ".git\n" {
		return fst.NewPath(fp.Join(p.String(), sOut)).String(), nil
	}
	return sOut, nil
	// return fst.NewPath(sOut), nil
}
