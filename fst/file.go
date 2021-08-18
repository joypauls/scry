// File System Tools
// Dealing with individual files here.
package fst

import (
	"io/fs"
	"log"
	"os"
	"time"
)

// Should use custom enum to restrict supported file types
type File struct {
	Name  string
	Size  int64
	IsDir bool
	IsReg bool
	Mode  os.FileMode
	Time  time.Time
	Perm  fs.FileMode
}

func MakeFile(d os.DirEntry) File {
	var f File
	f.Name = d.Name()
	f.IsDir = d.IsDir()
	fi, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	f.Size = fi.Size()
	f.Time = fi.ModTime()
	f.Perm = fi.Mode().Perm()
	f.IsReg = fi.Mode().IsRegular()
	f.Mode = fi.Mode()
	return f
}

func (f File) IsSymLink() bool {
	return f.Mode&os.ModeSymlink != 0
}
