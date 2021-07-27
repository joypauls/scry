package main

import (
	"fmt"
	"log"
	"os"

	ft "github.com/joypauls/file-scry/filetools"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	// print path to the current directory
	fmt.Println(ft.GetCurDir())
	for _, file := range files {
		var fs ft.FileStats
		fs.Populate(file)
		fmt.Printf("|%-5s|%-5d|%-9s|%s\n",
			fs.Label,
			fs.Time.Year(),
			fs.SizePretty,
			fs.Name,
		)
	}

	// fmt.Printf("%.1f\n", math.Mod(1433.023, 1000))
	// fmt.Println(1433 / 1000) // integer division truncates
	// formatter := fmt.Sprintf("%%-%ds", 3)
	// fmt.Println(formatter)
	// fmt.Println(fmt.Sprintf(formatter, "O"))
}
