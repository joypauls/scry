/*
The app package handles the main application logic.

Partially inspired by:
	https://github.com/nsf/termbox-go/blob/master/_demos/editbox.go
	https://github.com/gdamore/tcell/blob/master/TUTORIAL.md
*/
package app

import (
	"fmt"
	"log"
	fp "path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/fst"
	"github.com/mattn/go-runewidth"
)

// Set default text style

// this global config sucks let's get rid of it
var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
var selStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlueViolet)
var arrowLeft = '←'
var arrowRight = '→'
var arrowUp = '▲'
var arrowDown = '▼'

// initialize one time display-related configs at program start
// this could probably be a configuration struct
func initVars() {
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

// draw stuff that is not directory contents
func drawFrame(s tcell.Screen, app *App) {

	// top bar content
	maxHeaderLen := (7 * app.width) / 10 // 70% of width
	header := formatPath(app.path, maxHeaderLen)
	if app.UseEmoji {
		header = "🔮 " + header
	}
	draw(s, 0, 0, defStyle, header)

	if app.offset > 0 {
		draw(s, 0, 1, defStyle, fmt.Sprintf("%c", arrowUp))
	}
	if app.maxIndex+app.offset+1 < app.Size() {
		draw(s, 0, app.height-2, defStyle, fmt.Sprintf("%c", arrowDown))
	}
	// bottom line
	coordStr := fmt.Sprintf("(%d)", app.index)
	draw(s, app.xEnd-len(coordStr)+1, app.height-1, defStyle, coordStr)
	draw(s, 0, app.height-1, defStyle, "[esc] quit [h] home [b] initial")
}

// Actual file contents
func drawWindow(s tcell.Screen, app *App) {
	limit := minInt(app.windowHeight, app.maxIndex)
	for i := 0; i <= limit; i++ {
		drawFile(
			s,
			app.xStart,
			app.yStart+i,
			i == app.index,
			*app.File(i + app.offset),
			*app.path,
		)
	}
}

// Main object managing the app functionality and display.
type App struct {
	Layout
	*fst.Directory
	Config
	path     *fst.Path
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
	app.Read(app.path, app.ShowHidden)
	app.ResetIndex()
	app.offset = 0
}

// Move to the current selection if it's a directory, otherwise do nothing.
func (app *App) GoToChild() {
	if !app.IsEmpty() {
		f := app.File(app.index + app.offset) // pointer to a File
		if f.IsDir {
			app.path.Set(fp.Join(app.path.String(), f.Name))
			app.Read(app.path, app.ShowHidden)
			app.ResetIndex()
			app.offset = 0
		} // else do nothing
	}
}

// this should handle all drawing on the screen
func (app *App) Draw(s tcell.Screen) {
	drawFrame(s, app)
	if app.IsEmpty() {
		draw(s, app.xStart, app.yStart, defStyle, "<EMPTY>")
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
	app.path = c.InitDir.Copy()
	app.index = 0
	app.maxIndex = minInt(app.windowHeight-1, app.Size()-1)
	app.offset = 0
	return app
}

// Main program loop and user interactions
func Run(c Config) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	initVars() // make this not use global shit

	app := NewApp(s, c) // init
	// draw the ui for the first time
	app.Refresh(s)

	quit := func() {
		s.Fini()
	}

renderloop:
	for {
		s.Show()            // Update screen
		ev := s.PollEvent() // Poll event

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
				break renderloop
			} else if ev.Key() == tcell.KeyDown {
				// handle scrolling down
				if app.index == app.maxIndex {
					if app.maxIndex+app.offset == app.Size()-1 {
						app.index = 0
						app.offset = 0
					} else if app.maxIndex+app.offset < app.Size()-1 {
						// keep index the same! (at bottom)
						app.offset++
					}
				} else {
					app.AddIndex(1)
				}
			} else if ev.Key() == tcell.KeyUp {
				// handle scrolling up
				if app.index == 0 {
					if app.offset == 0 {
						app.index = app.maxIndex
						app.offset = (app.Size() - 1) - app.maxIndex
					} else if app.offset > 0 {
						// keep index the same (at top)
						app.offset--
					}
				} else {
					app.AddIndex(-1)
				}
			} else if ev.Key() == tcell.KeyLeft {
				app.GoToParent()
			} else if ev.Key() == tcell.KeyRight {
				app.GoToChild()
			} else if ev.Key() == tcell.KeyRune {
				if ev.Rune() == 'h' || ev.Rune() == 'H' {
					// go to user home directory (if it was found)
					app.path.Set(app.Home.String())
					app.Read(app.path, app.ShowHidden)
					app.ResetIndex()
					app.offset = 0
				} else if ev.Rune() == 'b' || ev.Rune() == 'B' {
					// go to initial directory
					app.path.Set(app.InitDir.String())
					app.Read(app.path, app.ShowHidden)
					app.ResetIndex()
					app.offset = 0
				}
			}
		case *tcell.EventError:
			// can we access the actual tcell error? idk
			log.Fatal("Panic! At the Unknown Input") // os.Exit(1) follows log.Fatal()
		}
		// draw after (potential) changes
		app.Refresh(s)
	}
}
