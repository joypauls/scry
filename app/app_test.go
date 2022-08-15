package app

import (
	"os"
	"testing"

	"github.com/joypauls/scry/fst"
	misc "github.com/joypauls/scry/internal"
)

var testFS = misc.GetTestFS()

func makeFullApp(t *testing.T) *App {
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
	app.maxIndex = app.Size() - 1 // slightly different logic but meh
	return app
}

func makeEmptyApp(t *testing.T) *App {
	dirRaw := []os.DirEntry{}
	c := Config{
		ShowHidden: false,
		UseEmoji:   true,
	}
	app := &App{
		Layout:    MakeLayout(100, 50),
		Directory: fst.NewDirectoryFromSlice(dirRaw, false),
		Config:    c,
		Path:      fst.NewPath("."),
		maxIndex:  len(dirRaw), // slightly different logic but meh
	}
	return app
}

func TestAppIndex(t *testing.T) {
	app := makeFullApp(t)

	// starting value
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// incremented value
	app.addIndex(1)
	if res, exp := app.Index(), 1; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// decremented value
	app.addIndex(-1)
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
}

func TestAppResetIndex(t *testing.T) {
	// stuff in the directory
	app := makeFullApp(t)
	app.addIndex(100)
	app.resetIndex()
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// empty directory
	app = makeEmptyApp(t)
	app.addIndex(100)
	app.resetIndex()
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
}

func TestAppShifting(t *testing.T) {
	app := makeFullApp(t)

	// normal scrolling
	app.addIndex(1)
	app.Up()
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// loop around to last element
	app.Up()
	if res, exp := app.Index(), 2; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// move offset when there is one
	app.resetIndex()
	app.maxIndex = 1
	app.offset = 1
	app.Up()
	if res, exp := app.offset, 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	app.resetIndex()

	// normal scrolling
	app.addIndex(1)
	app.Down()
	if res, exp := app.Index(), 2; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// loop around to first element
	app.Down()
	if res, exp := app.Index(), 0; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}

	// shift offset when at bottom and it's not at maximum
	app.resetIndex()
	app.addIndex(1)
	app.maxIndex = 1
	app.Down()
	if res, exp := app.offset, 1; res != exp {
		t.Errorf("Result: %d, Expected: %d", res, exp)
	}
}

func TestAppDraw(t *testing.T) {
	app := makeFullApp(t)
	s := misc.MakeTestScreen(t)

	// not a very useful test
	app.Draw(s)

	// not a very useful test
	app = makeEmptyApp(t)
	app.Draw(s)
}

func TestAppRefresh(t *testing.T) {
	app := makeFullApp(t)
	s := misc.MakeTestScreen(t)

	// not a very useful test
	app.Refresh(s)
}
