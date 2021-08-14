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

const coldef = termbox.ColorDefault // termbox.Attribute
var arrowLeft = '←'
var arrowRight = '→'

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

func drawFrame(app *App) {
	// top line
	draw(0, 0, coldef, coldef, app.path.Cur())
	// bottom line
	coordStr := fmt.Sprintf("(%d)", app.index)
	draw(app.layout.xEnd-len(coordStr)+1, app.layout.yEnd, coldef, coldef, coordStr)
	draw(0, app.layout.height-1, coldef, coldef, "[ESC] quit, [h] help")
}

func drawWindow(app *App) {
	// put check for empty here
	limit := minInt(app.layout.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			app.layout.xStart,
			app.layout.yStart+i,
			i == app.index,
			app.dir.File(i+app.head),
		)
	}
}

// Main object managing the app functionality and display.
type App struct {
	path     *fst.Path
	dir      *fst.Directory
	layout   *Layout
	index    int // 0 <= index < maxIndex
	maxIndex int
	head     int // start of window
	// tail     int // end of window
}

func (app *App) Index() int {
	return app.index
}

func (app *App) AddIndex(change int) {
	app.index = minInt(maxInt(app.index+change, 0), app.maxIndex)
}

func (app *App) ResetIndex() {
	app.index = 0
}

// func (app *App) MoveWindow(change int) {
// 	app.offset = minInt(maxInt(app.index+change, 0), app.maxIndex)
// }

// Move to the current parent.
func (app *App) GoToParent() {
	app.path.Set(app.path.Parent())
	app.dir = fst.NewDirectory(app.path) // this shouldn't be a whole new object
	app.ResetIndex()
	app.head = 0
}

// Move to the current selection if it's a directory, otherwise do nothing.
func (app *App) GoToChild() {
	f := app.dir.File(app.index)
	if f.IsDir {
		app.path.Set(fp.Join(app.path.Cur(), f.Name))
		app.dir = fst.NewDirectory(app.path)
		app.ResetIndex()
		app.head = 0
	} // else do nothing
}

func (app *App) Refresh() {
	termbox.Clear(coldef, coldef) // reset

	app.maxIndex = minInt(app.layout.windowHeight-1, app.dir.Size()-1)

	drawFrame(app)
	drawWindow(app) // main content

	termbox.Flush() // clean
}

func NewApp() *App {
	app := new(App)
	app.path = fst.InitPath() // init at cwd
	app.dir = fst.NewDirectory(app.path)
	app.layout = NewLayout()
	app.index = 0
	app.maxIndex = minInt(app.layout.windowHeight-1, app.dir.Size()-1)
	app.head = 0
	// app.tail = app.maxIndex
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
					if app.maxIndex+app.head+1 < app.dir.Size() {
						app.head++
					}
					// keep index the same (at bottom)
				} else {
					app.AddIndex(1)
				}
			case termbox.KeyArrowUp:
				app.AddIndex(-1)
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
