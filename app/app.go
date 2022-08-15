/*
The app package handles the main application logic.

Partially inspired by:
	https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go
	https://github.com/gdamore/tcell/blob/master/TUTORIAL.md
*/
package app

import (
	fp "path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
)

// Set default text style

// this global config sucks let's get rid of it please
// should just handle all this in config setup?
var theme = themes["fey"]

// var hlStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlueViolet)

var arrowLeft = '←'
var arrowRight = '→'
var arrowUp = '▲'
var arrowDown = '▼'

// initialize one time display-related configs at program start
// this could probably be a configuration struct
// after everything is declared
func init() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

// Main object managing the app functionality and display.
type App struct {
	*fst.Directory
	Layout
	Config
	Path *fst.Path
	// these below are critical to protect
	index    int // 0 <= index < maxIndex
	offset   int // start of window
	maxIndex int
}

func (app *App) addIndex(change int) {
	// meeting this condition means that there cur dir is empty
	// if app.index != 0 || app.maxIndex != 0 {
	if !app.IsEmpty() {
		app.index = minInt(maxInt(app.index+change, 0), app.maxIndex)
	}
}

func (app *App) resetIndex() {
	app.index = 0
	if app.IsEmpty() {
		app.maxIndex = 0
	} else {
		app.maxIndex = minInt(app.windowHeight-1, app.Size()-1)
	}
}

func (app *App) Index() int {
	return app.index
}

func (app *App) Top() {
	app.index = 0
	app.offset = 0
}

func (app *App) Bottom() {
	app.index = app.maxIndex
	app.offset = (app.Size() - 1) - app.maxIndex
}

func (app *App) Down() {
	if app.index == app.maxIndex {
		if app.maxIndex+app.offset == app.Size()-1 {
			app.Top()
		} else if app.maxIndex+app.offset < app.Size()-1 {
			// keep index the same! (at bottom)
			app.offset++
		}
	} else {
		app.addIndex(1)
	}
}

func (app *App) Up() {
	if app.index == 0 {
		if app.offset == 0 {
			app.Bottom()
		} else if app.offset > 0 {
			// keep index the same (at top)
			app.offset--
		}
	} else {
		app.addIndex(-1)
	}
}

func (app *App) Walk(p *fst.Path) {
	app.Path.Set(p.String())
	app.Read(app.Path, app.ShowHidden)
	app.resetIndex()
	app.offset = 0
}

// Move to the current selection if it's a directory, otherwise do nothing.
func (app *App) WalkToChild() {
	if !app.IsEmpty() {
		f := app.File(app.index + app.offset) // pointer to a File
		if f.IsDir {
			// this is kinda hacky
			app.Path.Set(fp.Join(app.Path.String(), f.Name))
			app.Walk(app.Path)
		}
		// else do nothing
	}
}

// this should handle all drawing on the screen
func (app *App) Draw(s tcell.Screen) {
	drawFrame(s, app)
	if app.IsProblem() {
		draw(s, app.xStart, app.yStart, theme.Default, app.Error())
	} else if app.IsEmpty() {
		draw(s, app.xStart, app.yStart, theme.Default, "<EMPTY>")
	} else {
		drawWindow(s, app)
	}
}

func (app *App) Refresh(s tcell.Screen) {
	s.Clear()
	app.Draw(s)
}

func NewApp(s tcell.Screen, c Config) *App {
	w, h := s.Size()
	app := &App{
		Layout:    MakeLayout(w, h),
		Directory: fst.NewDirectory(c.InitDir, c.ShowHidden),
		Config:    c,
	}
	app.Path = c.InitDir.Copy()
	app.maxIndex = minInt(app.windowHeight-1, app.Size()-1)
	return app
}
