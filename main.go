package main

import (
    "fmt"
    "os"
    "log"
)

func parseCurDir() string {
	path, err := os.Getwd()
	if err != nil {
    	log.Println(err)
	}
	return path
}

func parseDir(isDir bool) string {
	directoryLabel := "file"
	if isDir {
		directoryLabel = "dir"
	}
	return directoryLabel
}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// print path to the current directory
	fmt.Println(parseCurDir())
	// temporary seperator
	fmt.Println("-------")

	var fileName string
	var isDir bool
	for _, file := range files {
		fileName = file.Name()
		isDir = file.IsDir()

		fmt.Println(fileName + ", " + parseDir(isDir))
	}
}

