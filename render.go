// Main program loop and user interactions.
// Primary entrypoint for the fun stuff.
package main

import (
	"errors"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/joypauls/scry/app"
)

var defaultStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

func handleKeyEvent(e tcell.Event, s tcell.Screen, app *app.App) error {
	quit := func() { s.Fini() }
	// abstract out to test
	switch e := e.(type) {
	case *tcell.EventResize:
		s.Sync()
	case *tcell.EventKey:
		if e.Key() == tcell.KeyEscape || e.Key() == tcell.KeyCtrlC {
			quit()
			return errors.New("User closed the app. Nothing to do.")
		} else if e.Key() == tcell.KeyDown {
			app.Down()
		} else if e.Key() == tcell.KeyUp {
			app.Up()
		} else if e.Key() == tcell.KeyLeft {
			// we change the path but that's all so we have to walk to it still
			app.Path.ToParent()
			app.Walk(app.Path)
		} else if e.Key() == tcell.KeyRight {
			app.WalkToChild()
		} else if e.Key() == tcell.KeyRune {
			if e.Rune() == 'h' || e.Rune() == 'H' {
				// go to user home directory (if it was found)
				app.Walk(app.Home)
			} else if e.Rune() == 'b' || e.Rune() == 'B' {
				// go to initial directory
				app.Walk(app.InitDir)
			} else if e.Rune() == 'w' {
				app.Top()
			} else if e.Rune() == 's' {
				app.Bottom()
			}
		}
	case *tcell.EventError:
		// can we access the actual tcell error? idk
		log.Fatal("Panic! At The Unknown Input") // os.Exit(1) follows log.Fatal()
	}
	return nil
}

func render(c app.Config) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	// s.SetStyle(defaultStyle) // what was this for
	s.Clear()

	// initialize app object
	app := app.NewApp(s, c)
	// draw the ui for the first time
	app.Refresh(s)

renderloop:
	for {
		s.Show()                             // Update screen
		event := s.PollEvent()               // Poll event
		err := handleKeyEvent(event, s, app) // Handle event
		if err != nil {                      // Gracefully quit
			break renderloop
		}
		// draw after (potential) changes
		app.Refresh(s)
	}
}
