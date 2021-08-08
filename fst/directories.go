// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"log"
	"os"
)

func getFiles(path string) []*File {
	rawFiles, err := os.ReadDir(path)
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
	Path  string
	Files []*File
}

func NewDirectory(path string) *Directory {
	d := new(Directory)
	d.Path = path
	d.Files = getFiles(path)
	return d
}
