// This package is for handling all the file stuff.
package filetools

import (
	"fmt"
	"log"
	"os"
)

// SI defined base for multiple byte units
const SIBase = 1000
const SIPrefixes = "kMGTPE"

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

//////////////////////
// Directory Reader //
//////////////////////

type FileStats struct {
	name       string
	sizeRaw    int64
	sizePretty string
}

type Directory struct {
	files []FileStats
}

func (fs *FileStats) Populate(d os.DirEntry) {
	fs.name = d.Name()
	fileInfo, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	fs.sizeRaw = fileInfo.Size()
	fs.sizePretty = humanizeBytes(fs.sizeRaw)
}

func parseDir(isDir bool) string {
	directoryLabel := "file"
	if isDir {
		directoryLabel = "dir"
	}
	return directoryLabel
}

///////////////////////
// General Utilities //
///////////////////////

// returning the path of pwd
func GetCurDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
