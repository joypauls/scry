package app

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
)

var sysValue int

// only supporting this for now
var charset = "UTF-8"

// var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

// Emulating https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/io/fs/readdir_test.go
var testFs = fstest.MapFS{
	"hello.txt": {
		Data:    []byte("hello, world"),
		Mode:    0456,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
	// MapFS implicitly adds the directory file "sub" for us
	"sub/goodbye.txt": {
		Data:    []byte("goodbye, world"),
		Mode:    0456,
		ModTime: time.Now(),
		Sys:     &sysValue,
	},
}

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

func makeTestApp(t *testing.T) *App {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	c := Config{
		ShowHidden: false,
		UseEmoji:   true,
	}
	app := &App{
		Layout:    MakeLayout(100, 50),
		Directory: fst.NewDirectoryFromSlice(dirRaw, false),
		Config:    c,
		path:      fst.NewPath("."),
	}
	return app
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
		s.Show()
	}
}

func TestDrawFile(t *testing.T) {
	s := makeTestScreen(t, charset)
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	f := fst.MakeFile(dirRaw[0])
	p := fst.NewPath(".")

	// table of test cases
	tables := []struct {
		screen   tcell.Screen
		x        int
		y        int
		selected bool
		file     fst.File
		path     *fst.Path
	}{
		{s, 0, 0, false, f, p},
		{s, 10, 20, true, f, p},
	}

	// iterate over test tables
	for _, table := range tables {
		// doesn't throw an error so not sure how else to test this
		drawFile(table.screen, table.x, table.y, table.selected, table.file, table.path)
		s.Show()
	}
}

func TestDrawFrame(t *testing.T) {
	s := makeTestScreen(t, charset)
	app := makeTestApp(t)

	// doesn't throw an error so not sure how else to test this
	drawFrame(s, app)
	s.Show()
}

func TestDrawWindow(t *testing.T) {
	s := makeTestScreen(t, charset)
	app := makeTestApp(t)

	// doesn't throw an error so not sure how else to test this
	drawWindow(s, app)
	s.Show()
}
