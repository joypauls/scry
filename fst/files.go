// File System Tools
// Dealing with individual files here.
package fst

import (
	"log"
	"os"
	"time"
)

type File struct {
	Name  string
	Size  int64
	IsDir bool
	Time  time.Time
}

func NewFile(d os.DirEntry) *File {
	f := new(File) // new pointer to a File
	f.Name = d.Name()
	f.IsDir = d.IsDir()
	fileInfo, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	f.Size = fileInfo.Size()
	f.Time = fileInfo.ModTime()
	return f
}
