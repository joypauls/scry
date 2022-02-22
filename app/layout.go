package app

// Managing the UI layout
type Layout struct {
	width        int
	height       int
	xStart       int
	xEnd         int
	yStart       int
	yEnd         int
	windowHeight int // equals the # of available lines for dir contents
	middle       int
	innerPadding int
}

// constructor for Layout
func MakeLayout(w, h int) Layout {
	outerPadding := 2 // min for this is 2
	l := Layout{}
	l.width, l.height = w, h
	l.xStart = 0
	l.xEnd = w - 1
	l.yStart = 0 + outerPadding
	l.yEnd = h - 1 - outerPadding
	l.windowHeight = l.yEnd - l.yStart + 1
	l.middle = w / 2
	l.innerPadding = 3 // min for this is 2
	return l
}
