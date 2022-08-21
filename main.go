//Based on https://github.com/charmbracelet/bubbletea/tree/master/examples/list-default and other examples
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Starting up main program loop
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	m := newModel()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
