// File System Tools
// Dealing with individual files here.
package fst

import (
	"fmt"
	"log"
	"os"
	"time"
)

// SI defined base for multiple byte units
const SIBase = 1000

var SIPrefixes = [6]rune{'k', 'M', 'G', 'T', 'P', 'E'}

// symbols
const dirLabel = "📁"
const fileLabel = "  "

// Converts an integer number of bytes to SI units.
func humanizeBytes(bytes int64) string {
	if bytes < SIBase {
		return fmt.Sprintf("%d B", bytes) // < 1kB
	}
	magnitude := int64(SIBase)
	maxExp := 0
	for i := bytes / SIBase; i >= SIBase; i /= SIBase {
		magnitude *= SIBase
		maxExp++
	}
	return fmt.Sprintf(
		"%.1f %cB",
		float64(bytes)/float64(magnitude), // want quotient to be float
		SIPrefixes[maxExp],
	)
}

/////////////////
// File Reader //
/////////////////

type File struct {
	Name       string
	SizeRaw    int64
	SizePretty string
	Label      string
	Time       time.Time
}

func NewFile(d os.DirEntry) *File {
	f := new(File) // new pointer to a File
	f.Name = d.Name()
	f.Label = fileLabel
	if d.IsDir() {
		f.Label = dirLabel
	}
	fileInfo, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	f.SizeRaw = fileInfo.Size()
	f.SizePretty = humanizeBytes(f.SizeRaw)
	f.Time = fileInfo.ModTime()
	return f
}
