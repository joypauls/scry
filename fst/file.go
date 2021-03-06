// File System Tools
// Dealing with individual files here.
package fst

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

// this type implementations based on https://golang.org/doc/effective_go
type BytesSI float64

const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

func (b BytesSI) String() string {
	// can't even represent numbers in the etabyte range with float64 so don't support now
	switch {
	// case b >= YB:
	// 	return fmt.Sprintf("%.1f YB", b/YB)
	// case b >= ZB:
	// 	return fmt.Sprintf("%.1f ZB", b/ZB)
	// case b >= EB:
	// 	return fmt.Sprintf("%.1f EB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.1f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.1f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.1f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.1f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.1f KB", b/KB)
	}
	return fmt.Sprintf("%.0f B", b) // num of bytes is an exact quantity
}

// Should use custom enum to restrict supported file types
// Implements os.DirEntry
// SHould think about the cost/benefit of reimplimenting DirEntry methods vs just composing?
type File struct {
	Name      string
	Size      BytesSI
	IsDir     bool
	IsReg     bool
	IsSymLink bool
	Time      time.Time
	Perm      fs.FileMode
}

// // Needed to implement os.DirEntry
// func (f File) Name() string {
// 	return f.Name
// }

func MakeFile(d os.DirEntry) File {
	var f File
	f.Name = d.Name()
	f.IsDir = d.IsDir()
	fi, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	// fi.Size() is always an int64 but converting to float for math reasons
	f.Size = BytesSI(float64(fi.Size()))
	f.Time = fi.ModTime()
	f.Perm = fi.Mode().Perm()
	f.IsReg = fi.Mode().IsRegular()
	f.IsSymLink = fi.Mode()&os.ModeSymlink != 0
	return f
}
