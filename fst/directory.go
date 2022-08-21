// File System Tools
// Dealing with directories (defining as a collection of files/directories here)
package fst

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type SortMethod int

const (
	NameAsc SortMethod = iota
	NameDesc
	DirectoryAsc
	DirectoryDesc
)

func readDirectory(p *Path) ([]os.DirEntry, error) {
	contents, err := os.ReadDir(p.String()) // DirEntry slice
	if err != nil {
		if os.IsPermission(err) {
			return []os.DirEntry{}, errors.New("Permission Denied 🔒")
		}
		// should handle this with more care
		// we don't need to display the path because it is on the top status bar
		// return []os.DirEntry{}, fmt.Errorf("Couldn't read the directory: %s", p.String())
		return []os.DirEntry{}, errors.New("Error reading directory: Unknown ¯\\_(ツ)_/¯")
	}
	return contents, nil
}

func processDirectory(contents []os.DirEntry, p *Path, showHidden bool) []File {
	// is this the right way to build this slice?
	if showHidden {
		files := make([]File, len(contents))
		for i, f := range contents {
			files[i] = MakeFile(f, p)
		}
		return files
	}
	files := make([]File, 0)
	for _, f := range contents {
		match, _ := regexp.MatchString("^\\.", f.Name())
		if !match {
			files = append(files, MakeFile(f, p))
		}
	}
	return files
}

// Directory{} should stay identical to reading an empty directory
type Directory struct {
	files []File
	size  int
	err   error // where to store error if there's a problem processing the directory
}

func (d *Directory) File(i int) File {
	if i < 0 || i >= d.size {
		fmt.Println(i)
		log.Fatal("Requested a file from index that doesn't exist")
	}
	return d.files[i]
}

func (d *Directory) Files() []File {
	return d.files
}

// func (d *Directory) SortNameDesc() {
// 	sort.Slice(d.files, func(i, j int) bool {
// 		// this feels inefficient
// 		return strings.ToLower(d.files[i].Name) < strings.ToLower(d.files[j].Name)
// 	})
// }

// Sorts slice of Files in place.
func (d *Directory) Sort(method SortMethod) {
	switch method {
	// sort by file name
	case NameAsc:
		sort.Slice(d.files, func(i, j int) bool {
			return strings.ToLower(d.files[i].Name()) < strings.ToLower(d.files[j].Name())
		})
	case NameDesc:
		sort.Slice(d.files, func(i, j int) bool {
			return strings.ToLower(d.files[i].Name()) > strings.ToLower(d.files[j].Name())
		})
	// sort by whether file is a directory and secondarily by name
	case DirectoryAsc:
		sort.Slice(d.files, func(i, j int) bool {
			if d.files[i].IsDir() && d.files[j].IsDir() {
				return strings.ToLower(d.files[i].Name()) > strings.ToLower(d.files[j].Name())
			} else if !d.files[i].IsDir() && !d.files[j].IsDir() {
				return strings.ToLower(d.files[i].Name()) > strings.ToLower(d.files[j].Name())
			} else if d.files[i].IsDir() {
				if !d.files[j].IsDir() {
					return false
				}
			} else if d.files[j].IsDir() {
				return true
			}
			return false
		})
	case DirectoryDesc:
		sort.Slice(d.files, func(i, j int) bool {
			if d.files[i].IsDir() && d.files[j].IsDir() {
				return strings.ToLower(d.files[i].Name()) < strings.ToLower(d.files[j].Name())
			} else if !d.files[i].IsDir() && !d.files[j].IsDir() {
				return strings.ToLower(d.files[i].Name()) < strings.ToLower(d.files[j].Name())
			} else if d.files[i].IsDir() {
				if !d.files[j].IsDir() {
					return true
				}
			} else if d.files[j].IsDir() {
				return false
			}
			return true
		})
	}
}

func (d *Directory) Size() int {
	return d.size
}

func (d *Directory) IsEmpty() bool {
	return d.size == 0
}

func (d *Directory) IsProblem() bool {
	return d.err != nil
}

func (d *Directory) Error() string {
	return d.err.Error()
}

func (d *Directory) Read(p *Path, showHidden bool) {
	dir, err := readDirectory(p)
	if err != nil {
		d.files = []File{}
	} else {
		d.files = processDirectory(dir, p, showHidden)
		d.Sort(DirectoryDesc) // just sort for default for now
	}
	d.err = err
	d.size = len(d.files)
}

func NewDirectory(p *Path, showHidden bool) *Directory {
	d := new(Directory)
	d.Read(p, showHidden)
	return d
}

// This is for tests
func NewDirectoryFromSlice(dir []os.DirEntry, p *Path, showHidden bool) *Directory {
	d := new(Directory)
	d.files = processDirectory(dir, p, showHidden)
	d.size = len(d.files)
	d.Sort(NameAsc)
	return d
}
