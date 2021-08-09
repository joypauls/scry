// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"log"
	"os"
)

func getFiles(p *Path) []*File {
	rawFiles, err := os.ReadDir(p.Cur())
	if err != nil {
		log.Fatal(err)
	}
	var files []*File
	for _, f := range rawFiles {
		files = append(files, NewFile(f))
	}
	return files
}

type Directory struct {
	Path  *Path
	Files []*File
}

func NewDirectory(p *Path) *Directory {
	d := new(Directory)
	d.Path = p
	d.Files = getFiles(p)
	return d
}
