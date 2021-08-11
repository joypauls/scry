// This package is handling the printing, terminal functionality, and user input.
package app

// partially nspired by https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go

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
	draw(0, 0, coldef, coldef, app.dir.Path.Cur())
	// bottom line
	coordStr := fmt.Sprintf("(%d)", app.index)
	draw(app.layout.xEnd-len(coordStr)+1, app.layout.yEnd, coldef, coldef, coordStr)
	draw(0, app.layout.yEnd, coldef, coldef, "[ESC] quit, [h] help")
}

func drawWindow(app *App) {
	for i, f := range app.dir.Files {
		drawFile(0, 0+app.layout.topPad+i, i == app.index, f)
	}
}

// main object managing the app functionality and display
type App struct {
	path     *fst.Path
	dir      *fst.Directory
	layout   *Layout
	index    int
	maxIndex int
}

func InitApp() *App {
	app := new(App)
	app.path = fst.InitPath()
	app.dir = fst.NewDirectory(app.path)
	app.layout = NewLayout()
	app.index = 0
	app.maxIndex = len(app.dir.Files)
	return app
}

func (app *App) IncrementIndex(change int) {
	app.index = minInt(maxInt(app.index+change, 0), app.maxIndex)
}

func (app *App) Refresh() {
	termbox.Clear(coldef, coldef) // reset

	app.maxIndex = len(app.dir.Files) - 1 // update num files

	drawFrame(app)
	drawWindow(app) // main content

	termbox.Flush() // clean
}

// Main program loop and user interactions
func Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	config()

	app := InitApp()

	// draw the UI for the first time
	app.Refresh()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break loop
			case termbox.KeyArrowDown:
				app.IncrementIndex(1)
			case termbox.KeyArrowUp:
				app.IncrementIndex(-1)
			case termbox.KeyArrowLeft:
				app.path.Set(app.path.Parent())
				app.dir = fst.NewDirectory(app.path) // this shouldn't be a whole new object
				app.index = 0
			case termbox.KeyArrowRight:
				sel := app.dir.Files[app.index]
				if sel.IsDir {
					app.path.Set(fp.Join(app.path.Cur(), sel.Name))
					app.dir = fst.NewDirectory(app.path)
					app.index = 0
				}
			}
		case termbox.EventError:
			log.Fatal(ev.Err) // os.Exit(1) follows
		}

		app.Refresh()
	}
}
