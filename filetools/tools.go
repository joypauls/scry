// This package is for handling all the file stuff.
// - could store # of children for dir
package filetools

import (
	"fmt"
	"log"
	"os"
	"time"
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
	Name       string
	SizeRaw    int64
	SizePretty string
	Label      string
	Time       time.Time
}

func (fs *FileStats) Populate(d os.DirEntry) {
	fs.Name = d.Name()
	fs.Label = "f"
	if d.IsDir() {
		fs.Label = "d"
	}
	fileInfo, err := d.Info() // FileInfo
	if err != nil {
		log.Fatal(err)
	}
	fs.SizeRaw = fileInfo.Size()
	fs.SizePretty = humanizeBytes(fs.SizeRaw)
	fs.Time = fileInfo.ModTime()
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
