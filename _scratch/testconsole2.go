package main

import (
	"fmt"

	"github.com/containerd/console"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	current := console.Current()
	defer current.Reset()

	if err := current.SetRaw(); err != nil {
		panic(err)
	}

	term := terminal.NewTerminal(current, "")
	term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		// fmt.Println("callback:", line, pos, key)

		return "", 0, false
	}

	line, err := term.ReadLine()
	fmt.Println("result:", line, err)
}
