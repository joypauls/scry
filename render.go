// Main program loop and user interactions.
// Primary entrypoint for the fun stuff.
package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/app"
)

var defaultStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

// func handleKeyboardEvents(app *app.App) tcell.Event {

// }

func Render(c app.Config) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defaultStyle)
	s.Clear()

	// initialize app object
	app := app.NewApp(s, c)
	// draw the ui for the first time
	app.Refresh(s)
	quit := func() { s.Fini() }

renderloop:
	for {
		s.Show()            // Update screen
		ev := s.PollEvent() // Poll event

		// abstract out to test
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
				break renderloop
			} else if ev.Key() == tcell.KeyDown {
				app.ShiftDown()
			} else if ev.Key() == tcell.KeyUp {
				app.ShiftUp()
			} else if ev.Key() == tcell.KeyLeft {
				// we change the path but that's all so we have to walk to it still
				app.Path.ToParent()
				app.Walk(app.Path)
			} else if ev.Key() == tcell.KeyRight {
				app.WalkToChild()
			} else if ev.Key() == tcell.KeyRune {
				if ev.Rune() == 'h' || ev.Rune() == 'H' {
					// go to user home directory (if it was found)
					app.Walk(app.Home)
				} else if ev.Rune() == 'b' || ev.Rune() == 'B' {
					// go to initial directory
					app.Walk(app.InitDir)
				}
			}
		case *tcell.EventError:
			// can we access the actual tcell error? idk
			log.Fatal("Panic! At The Unknown Input") // os.Exit(1) follows log.Fatal()
		}
		// draw after (potential) changes
		app.Refresh(s)
	}
}
