package app

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	misc "github.com/joypauls/scry/internal"
)

func makeBasicApp(t *testing.T) *App {
	testFS := misc.GetTestFS()
	dirRaw, err := testFS.ReadDir(".")
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
		Path:      fst.NewPath("."),
	}
	return app
}

func TestDraw(t *testing.T) {
	s := misc.MakeTestScreen(t)
	// table of test cases
	tables := []struct {
		screen tcell.Screen
		x      int
		y      int
		style  tcell.Style
		text   string
	}{
		{s, 0, 0, theme.Default, "some test text "},
		{s, 10, 20, theme.Default, "/test/path/display/"},
		{s, 0, 30, theme.Default, "📁 file 📁 stuff 📁 emoji 📁"},
	}

	// iterate over test tables
	for _, table := range tables {
		// doesn't throw an error so not sure how else to test this
		draw(table.screen, table.x, table.y, table.style, table.text)
		s.Show()
	}
}

func TestDrawFile(t *testing.T) {
	s := misc.MakeTestScreen(t)
	dirRaw, err := testFS.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	f := fst.MakeFile(dirRaw[0])
	p := fst.NewPath(".")

	// table of test cases
	tables := []struct {
		screen   tcell.Screen
		x1       int
		x2       int
		y        int
		selected bool
		file     fst.File
		path     *fst.Path
	}{
		{s, 0, 30, 0, false, f, p},
		{s, 10, 30, 20, true, f, p},
	}

	// iterate over test tables
	for _, table := range tables {
		// doesn't throw an error so not sure how else to test this
		drawFile(table.screen, table.x1, table.x2, table.y, table.selected, table.file, table.path)
		s.Show()
	}
}

func TestDrawFrame(t *testing.T) {
	s := misc.MakeTestScreen(t)
	app := makeBasicApp(t)

	// doesn't throw an error so not sure how else to test this
	drawFrame(s, app)
	s.Show()
}

func TestDrawWindow(t *testing.T) {
	s := misc.MakeTestScreen(t)
	app := makeBasicApp(t)

	// doesn't throw an error so not sure how else to test this
	drawWindow(s, app)
	s.Show()
}
