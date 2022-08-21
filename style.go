package main

import "github.com/charmbracelet/lipgloss"

const (
	bullet   = "•"
	ellipsis = "…"
)

var docStyle = lipgloss.NewStyle().Margin(0, 1)
var titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#4621AD")).Padding(0, 1)
var statusMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).Render

// DefaultItemStyles defines styling for a default list item.
// See DefaultItemView for when these come into play.
type FileStyles struct {
	// The Normal state.
	NormalName lipgloss.Style
	NormalSize lipgloss.Style

	// The selected item state.
	SelectedName lipgloss.Style
	SelectedSize lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	DimmedName lipgloss.Style
	DimmedSize lipgloss.Style

	// Charcters matching the current filter, if any.
	FilterMatch lipgloss.Style
}

// NewDefaultItemStyles returns style definitions for a default item. See
// DefaultItemView for when these come into play.
func NewFileStyles() (s FileStyles) {
	s.NormalName = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2)

	s.NormalSize = s.NormalName.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
		Align(lipgloss.Right)

	s.SelectedName = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		Padding(0, 0, 0, 1)

	s.SelectedSize = lipgloss.NewStyle().
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		Padding(0, 0, 0, 2).
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	s.DimmedName = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
		Padding(0, 0, 0, 1)

	s.DimmedSize = s.DimmedName.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
