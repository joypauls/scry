/*
This package is handling the printing, terminal functionality, and user input.

Partially inspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go
*/
package app

import (
	"fmt"
	"log"
	fp "path/filepath"

	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// this global config sucks let's get rid of it
const coldef = termbox.ColorDefault // termbox.Attribute
var arrowLeft = '←'
var arrowRight = '→'
var arrowUp = '▲'
var arrowDown = '▼'

// initialize one time display-related configs at program start
// this could probably be a configuration struct
func config() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// draw stuff that is not file content
func drawFrame(app *App) {
	// top line
	draw(0, 0, coldef, coldef, app.path.String())
	if app.offset > 0 {
		draw(0, 1, coldef, coldef, fmt.Sprintf("%c", arrowUp))
	}
	if app.maxIndex+app.offset+1 < app.Size() {
		draw(0, app.height-2, coldef, coldef, fmt.Sprintf("%c", arrowDown))
	}
	// bottom line
	coordStr := fmt.Sprintf("(%d)", app.index)
	draw(app.xEnd-len(coordStr)+1, app.height-1, coldef, coldef, coordStr)
	draw(0, app.height-1, coldef, coldef, "[ESC] quit, [h] help")
}

// fraw file content
func drawWindow(app *App) {
	// put check for empty here
	limit := minInt(app.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			app.xStart,
			app.yStart+i,
			i == app.index,
			app.File(i+app.offset),
		)
	}
}

// Main object managing the app functionality and display.
type App struct {
	Layout
	*fst.Directory
	path     *fst.Path
	home     *fst.Path
	index    int // 0 <= index < maxIndex
	maxIndex int
	offset   int // start of window
}

func (app *App) Index() int {
	return app.index
}

func (app *App) AddIndex(change int) {
	// meeting this condition means that there cur dir is empty
	if app.index != 0 || app.maxIndex != 0 {
		app.index = minInt(maxInt(app.index+change, 0), app.maxIndex)
	}
}

func (app *App) ResetIndex() {
	app.index = 0
	if app.IsEmpty() {
		app.maxIndex = 0
	} else {
		app.maxIndex = minInt(app.windowHeight-1, app.Size()-1)
	}
}

// Move to the current parent.
func (app *App) GoToParent() {
	app.path.Set(app.path.Parent())
	app.Read(app.path)
	app.ResetIndex()
	app.offset = 0
}

// Move to the current selection if it's a directory, otherwise do nothing.
func (app *App) GoToChild() {
	if !app.IsEmpty() {
		f := app.File(app.index + app.offset) // pointer to a File
		if f.IsDir {
			app.path.Set(fp.Join(app.path.String(), f.Name))
			app.Read(app.path)
			app.ResetIndex()
			app.offset = 0
		} // else do nothing
	}
}

// this should handle all drawing on the screen
func (app *App) Draw() {
	drawFrame(app)
	if app.IsEmpty() {
		draw(app.xStart, app.yStart, coldef, coldef, "<EMPTY>")
	} else {
		drawWindow(app)
	}
}

func (app *App) Refresh() {
	termbox.Clear(coldef, coldef) // reset
	app.Draw()
	termbox.Flush() // clean
}

func NewApp() *App {
	app := &App{Layout: MakeLayout(), Directory: fst.NewDirectory(fst.NewPath())}
	app.path = fst.NewPath() // init at wd
	app.home = fst.NewPath() // could do a deep copy but it's cheap so meh
	app.index = 0
	app.maxIndex = minInt(app.windowHeight-1, app.Size()-1)
	app.offset = 0
	return app
}

// Main program loop and user interactions
func Run() {
	// setting up
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	config() // make this not use global shit

	app := NewApp() // init
	// draw the ui for the first time
	app.Refresh()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break loop
			case termbox.KeyArrowDown:
				// handle scrolling down
				if app.index == app.maxIndex {
					if app.maxIndex+app.offset+1 < app.Size() {
						// keep index the same! (at bottom)
						app.offset++
					}
				} else {
					app.AddIndex(1)
				}
			case termbox.KeyArrowUp:
				// handle scrolling up
				if app.index == 0 && app.offset > 0 {
					// keep index the same (at top)
					app.offset--
				} else {
					app.AddIndex(-1)
				}
			case termbox.KeyArrowLeft:
				app.GoToParent()
			case termbox.KeyArrowRight:
				app.GoToChild()
			}
		case termbox.EventError:
			log.Fatal(ev.Err) // os.Exit(1) follows
		}

		// draw after (potential) changes
		app.Refresh()
	}
}
