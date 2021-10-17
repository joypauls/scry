package app

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

// only supporting this for now
var charset = "UTF-8"

// var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

// Lifted from https://github.com/gdamore/tcell/blob/master/sim_test.go
func makeTestScreen(t *testing.T, charset string) tcell.SimulationScreen {
	s := tcell.NewSimulationScreen(charset)
	if s == nil {
		t.Fatalf("Failed to get simulation screen")
	}
	if e := s.Init(); e != nil {
		t.Fatalf("Failed to initialize screen: %v", e)
	}
	return s
}

func TestDraw(t *testing.T) {
	s := makeTestScreen(t, charset)
	// table of test cases
	tables := []struct {
		screen tcell.Screen
		x      int
		y      int
		style  tcell.Style
		text   string
	}{
		{s, 0, 0, defStyle, "some test text "},
		{s, 10, 20, defStyle, "/test/path/display/"},
		{s, 0, 30, defStyle, "📁 file 📁 stuff 📁 emoji 📁"},
	}

	// iterate over test tables
	for _, table := range tables {
		// doesn't throw an error so not sure how else to test this
		draw(table.screen, table.x, table.y, table.style, table.text)
	}
}
