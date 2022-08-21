// File System Tools
// Dealing with individual files here.
package fst

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	fp "path/filepath"
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
// Should think about the cost/benefit of reimplimenting DirEntry methods vs just composing?
type File struct {
	name          string
	size          BytesSI
	isDir         bool
	isReg         bool
	isSymLink     bool
	symLinkTarget string
	time          time.Time
	perm          fs.FileMode
}

// // Needed to implement os.DirEntry
// func (f File) Name() string {
// 	return f.Name
// }

func (f File) Name() string {
	return f.name
}

func (f File) Size() BytesSI {
	return f.size
}

func (f File) IsDir() bool {
	return f.isDir
}

func (f File) IsSymLink() bool {
	return f.isSymLink
}

func (f File) SymLinkTarget() string {
	return f.symLinkTarget
}

func (f File) Time() time.Time {
	return f.time
}

func (f File) Perm() fs.FileMode {
	return f.perm
}

func MakeFile(d os.DirEntry, p *Path) File {
	var f File
	f.name = d.Name()
	f.isDir = d.IsDir()
	fi, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	// fi.Size() is always an int64 but converting to float for math reasons
	f.size = BytesSI(float64(fi.Size()))
	f.time = fi.ModTime()
	f.perm = fi.Mode().Perm()
	f.isReg = fi.Mode().IsRegular()
	f.isSymLink = fi.Mode()&os.ModeSymlink != 0

	if f.isSymLink {
		target, err := os.Readlink(fp.Join(p.String(), f.name))
		if err != nil {
			target = "?"
		}
		f.symLinkTarget = target
	}

	return f
}
