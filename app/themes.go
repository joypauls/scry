package app

import "github.com/gdamore/tcell/v2"

type Theme struct {
	Default     tcell.Style
	DefaultEmph tcell.Style
	Selected    tcell.Style
	Highlight   tcell.Style
}

var themes = map[string]Theme{
	"fey": {
		Default:     tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset),
		DefaultEmph: tcell.StyleDefault.Foreground(tcell.ColorLightPink.TrueColor()).Background(tcell.ColorReset),
		Selected:    tcell.StyleDefault.Foreground(tcell.ColorLightPink.TrueColor()).Background(tcell.ColorDarkSlateBlue.TrueColor()),
		Highlight:   tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkSlateBlue.TrueColor()),
	},
	"l33t": {
		Default:     tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset),
		DefaultEmph: tcell.StyleDefault.Foreground(tcell.ColorLawnGreen.TrueColor()).Background(tcell.ColorReset),
		Selected:    tcell.StyleDefault.Foreground(tcell.ColorLawnGreen.TrueColor()).Background(tcell.ColorReset),
		Highlight:   tcell.StyleDefault.Foreground(tcell.ColorLawnGreen.TrueColor()).Background(tcell.ColorReset),
	},
	// "witch": {
	// 	Default:   tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
	// 	Selected:  tcell.StyleDefault.Foreground(tcell.ColorLightPink.TrueColor()).Background(tcell.ColorDarkSlateBlue.TrueColor()),
	// 	Highlight: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorIndigo.TrueColor()),
	// },
	// "witch": {
	// 	Default:   tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
	// 	Selected:  tcell.StyleDefault.Foreground(tcell.ColorLightCoral).Background(tcell.ColorBlack),
	// 	Highlight: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorIndigo),
	// },
	// "witch2": {
	// 	Default:   tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
	// 	Highlight: tcell.StyleDefault.Foreground(tcell.ColorLightCoral).Background(tcell.ColorBlack),
	// 	Selected:  tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorIndigo),
	// },
	// "fey": {
	// 	Default:   tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
	// 	Selected:  tcell.StyleDefault.Foreground(tcell.ColorLightPink.TrueColor()).Background(tcell.ColorDarkSlateBlue.TrueColor()),
	// 	Highlight: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkSlateBlue.TrueColor()),
	// },
	// "wizard": {
	// 	Default:   tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
	// 	Selected:  tcell.StyleDefault.Foreground(tcell.ColorSpringGreen).Background(tcell.ColorBlack),
	// 	Highlight: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorMediumBlue),
	// },
}
