package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/app"
	"github.com/joypauls/scry/fst"
	misc "github.com/joypauls/scry/internal"
)

var testFS = misc.GetTestFS()

func makeFullApp(t *testing.T) *app.App {
	dirRaw, err := testFS.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	c := app.Config{
		ShowHidden: false,
		UseEmoji:   true,
		InitDir:    fst.NewPath("."),
		Home:       fst.NewPath("."),
	}
	app := &app.App{
		Layout:    app.MakeLayout(100, 50),
		Directory: fst.NewDirectoryFromSlice(dirRaw, false),
		Config:    c,
		Path:      fst.NewPath("."),
	}
	return app
}

// right now just testing that nothing catastrophic happens
func TestHandleKeyEvent(t *testing.T) {
	app := makeFullApp(t)
	s := misc.MakeTestScreen(t)

	e := tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone)
	err := handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyLeft, ' ', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyRight, ' ', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyRune, 'h', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err != nil {
		t.Fatalf("Key event handler encountered error: %s", err)
	} else {
	}

	e = tcell.NewEventKey(tcell.KeyEsc, ' ', tcell.ModNone)
	err = handleKeyEvent(e, s, app)
	if err == nil {
		t.Fatalf("App closed - should have returned dummy error.")
	} else {
	}
}

// func TestUsageText(t *testing.T) {
// 	app := makeFullApp(t)
// 	s := tu.MakeTestScreen(t)
// 	e :=

// 	numLines := strings.Count(formatUsageText(), "\n")
// 	if res, exp := numLines, 10; res != exp {
// 		t.Errorf("Result: %d, Expected: %d", res, exp)
// 	}
// }
