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
}

// constructor for Layout
func MakeLayout(w, h int) Layout {
	padding := 2 // min for this is 2
	f := Layout{}
	f.width, f.height = w, h
	f.xStart = 1
	f.xEnd = w - 1
	f.yStart = 0 + padding
	f.yEnd = h - 1 - padding
	f.windowHeight = f.yEnd - f.yStart + 1
	return f
}
