// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func readFiles(p *Path, showHidden bool) []File {
	rawFiles, err := os.ReadDir(p.String()) // DirEntry slice
	if err != nil {
		log.Fatal(err)
	}
	// is this the right way to build this slice?
	if showHidden {
		files := make([]File, len(rawFiles))
		for i, f := range rawFiles {
			files[i] = MakeFile(f)
		}
		return files
	}
	files := make([]File, 0)
	for _, f := range rawFiles {
		match, _ := regexp.MatchString("^\\.", f.Name())
		if !match {
			files = append(files, MakeFile(f))
		}
	}
	return files
}

// Directory{} should stay identical to reading an empty directory
type Directory struct {
	files []File
	size  int
}

func (d *Directory) File(i int) *File {
	if i < 0 || i >= d.size {
		fmt.Println(i)
		log.Fatal("Requested a file from index that doesn't exist")
	}
	return &d.files[i]
}

func (d *Directory) Files() []File {
	return d.files
}

func (d *Directory) Size() int {
	return d.size
}

func (d *Directory) IsEmpty() bool {
	return d.size == 0
}

func (d *Directory) Read(p *Path, showHidden bool) {
	d.files = readFiles(p, showHidden)
	d.size = len(d.files)
}

func NewDirectory(p *Path, showHidden bool) *Directory {
	d := new(Directory)
	d.Read(p, showHidden)
	return d
}
