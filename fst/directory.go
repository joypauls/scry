// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"log"
	"os"
)

func readFiles(p *Path) []*File {
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
	files []*File
}

func (d *Directory) File(i int) *File {
	if i < 0 || i >= len(d.files) {
		log.Fatal("Requested a file from index that doesn't exist")
	}
	return d.files[i]
}

func (d *Directory) Files() []*File {
	return d.files
}

func (d *Directory) Size() int {
	return len(d.files)
}

func NewDirectory(p *Path) *Directory {
	d := new(Directory)
	d.files = readFiles(p)
	return d
}
