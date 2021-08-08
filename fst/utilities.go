package fst

import (
	"log"
	"os"
)

// returning the path of pwd
func GetCurDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
