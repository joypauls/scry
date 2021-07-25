package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

// SI defined base for multiple byte units
const SIUnit = 1000
const SIPrefixes = "kMGTPE"

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

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func humanizeBytes(bytes int64) string {
	if bytes < SIUnit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(SIUnit), 0
	for n := bytes / SIUnit; n >= SIUnit; n /= SIUnit {
		div *= SIUnit
		exp++
	}
	// how does this impact precision?
	// want always max of 3 left of decimal, max of 1 after
	return fmt.Sprintf("%.1f %cB",
		float64(bytes)/float64(div), SIPrefixes[exp])
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

	for _, file := range files {
		fileName := file.Name()      // string
		isDir := file.IsDir()        // bool
		fileInfo, err := file.Info() // FileInfo
		if err != nil {
			log.Fatal(err)
		}
		fileSize := fileInfo.Size() // int

		// fmt.Println(fileName + ", " + parseDir(isDir))
		fmt.Printf("|%-5s|%-9s|%s\n",
			parseDir(isDir),
			humanizeBytes(fileSize),
			fileName,
		)
	}

	fmt.Printf("%.1f\n", math.Mod(1433.023, 1000))
	fmt.Println(1433 / 1000) // integer division truncates
	fmt.Println(humanizeBytes(1433))
}
