package app

import (
	"testing"

	"github.com/joypauls/scry/fst"
	tu "github.com/joypauls/scry/internal"
)

var testFS = tu.GetTestFS()

func makeNormalApp(t *testing.T) *App {
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

// func TestDraw(t *testing.T) {
// 	s := makeTestScreen(t, charset)
// 	// table of test cases
// 	tables := []struct {
// 		screen tcell.Screen
// 		x      int
// 		y      int
// 		style  tcell.Style
// 		text   string
// 	}{
// 		{s, 0, 0, defStyle, "some test text "},
// 		{s, 10, 20, defStyle, "/test/path/display/"},
// 		{s, 0, 30, defStyle, "📁 file 📁 stuff 📁 emoji 📁"},
// 	}

// 	// iterate over test tables
// 	for _, table := range tables {
// 		// doesn't throw an error so not sure how else to test this
// 		draw(table.screen, table.x, table.y, table.style, table.text)
// 		s.Show()
// 	}
// }

// func TestDrawFile(t *testing.T) {
// 	s := makeTestScreen(t, charset)
// 	dirRaw, err := testFs.ReadDir(".")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	f := fst.MakeFile(dirRaw[0])
// 	p := fst.NewPath(".")

// 	// table of test cases
// 	tables := []struct {
// 		screen   tcell.Screen
// 		x        int
// 		y        int
// 		selected bool
// 		file     fst.File
// 		path     *fst.Path
// 	}{
// 		{s, 0, 0, false, f, p},
// 		{s, 10, 20, true, f, p},
// 	}

// 	// iterate over test tables
// 	for _, table := range tables {
// 		// doesn't throw an error so not sure how else to test this
// 		drawFile(table.screen, table.x, table.y, table.selected, table.file, table.path)
// 		s.Show()
// 	}
// }

// func TestDrawFrame(t *testing.T) {
// 	s := makeTestScreen(t, charset)
// 	app := makeTestApp(t)

// 	// doesn't throw an error so not sure how else to test this
// 	drawFrame(s, app)
// 	s.Show()
// }

// func TestDrawWindow(t *testing.T) {
// 	s := makeTestScreen(t, charset)
// 	app := makeTestApp(t)

// 	// doesn't throw an error so not sure how else to test this
// 	drawWindow(s, app)
// 	s.Show()
// }
