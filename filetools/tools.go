// This package is handling the file stuff.
package filetools

import (
	"log"
	"os"
)

// returning the path of pwd
func ParseCurDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
