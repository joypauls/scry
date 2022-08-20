// HEAVILY based on https://github.com/charmbracelet/bubbles/blob/v0.13.0/list/defaultitem.go
package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

// Need to implement ItemDelegate
// https://pkg.go.dev/github.com/charmbracelet/bubbles@v0.13.0/list#ItemDelegate
//
// --------------------------------------------------------
// type ItemDelegate interface {
// 	// Render renders the item's view.
// 	Render(w io.Writer, m Model, index int, item Item)

// 	// Height is the height of the list item.
// 	Height() int

// 	// Spacing is the size of the horizontal gap between list items in cells.
// 	Spacing() int

// 	// Update is the update loop for items. All messages in the list's update
// 	// loop will pass through here except when the user is setting a filter.
// 	// Use this method to perform item-level updates appropriate to this
// 	// delegate.
// 	Update(msg tea.Msg, m *Model) tea.Cmd
// }
// --------------------------------------------------------
type FileDelegate struct {
	ShowDescription bool
	Styles          FileStyles
	UpdateFunc      func(tea.Msg, *list.Model) tea.Cmd
	ShortHelpFunc   func() []key.Binding
	FullHelpFunc    func() [][]key.Binding
	height          int
	spacing         int
}

// Use some sensible defaults
func NewFileDelegate() FileDelegate {
	fd := FileDelegate{
		ShowDescription: true,
		Styles:          NewFileStyles(),
		height:          1,
		spacing:         1,
	}

	keys := newDelegateKeyMap()

	fd.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return m.NewStatusMessage(statusMessageStyle("You chose " + title))

			case key.Matches(msg, keys.remove):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("Deleted " + title))
			}
		}

		return nil
	}

	// Stuff for keybinding help
	help := []key.Binding{keys.choose, keys.remove}
	fd.ShortHelpFunc = func() []key.Binding {
		return help
	}
	fd.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return fd
}

// SetHeight sets delegate's preferred height.
func (fd *FileDelegate) SetHeight(i int) {
	fd.height = i
}

// Height returns the delegate's preferred height.
// This has effect only if ShowDescription is true,
// otherwise height is always 1.
func (fd FileDelegate) Height() int {
	// if fd.ShowDescription {
	// 	return fd.height
	// }
	return fd.height
}

// SetSpacing set the delegate's spacing.
func (fd *FileDelegate) SetSpacing(i int) {
	fd.spacing = i
}

// Spacing returns the delegate's spacing.
func (fd FileDelegate) Spacing() int {
	return fd.spacing
}

// Update checks whether the delegate's UpdateFunc is set and calls it.
func (fd FileDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	if fd.UpdateFunc == nil {
		return nil
	}
	return fd.UpdateFunc(msg, m)
}

// Render prints an item.
func (d FileDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, desc  string
		matchedRunes []int
		s            = &d.Styles
	)

	if i, ok := item.(list.DefaultItem); ok {
		title = i.Title()
		desc = i.Description()
	} else {
		return
	}

	// short-circuit
	if m.Width() <= 0 {
		return
	}

	// Prevent text from exceeding list width
	textwidth := uint(m.Width() - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight())
	title = truncate.StringWithTail(title, textwidth, ellipsis)
	// if d.ShowDescription {
	// 	var lines []string
	// 	for i, line := range strings.Split(desc, "\n") {
	// 		if i >= d.height-1 {
	// 			break
	// 		}
	// 		lines = append(lines, truncate.StringWithTail(line, textwidth, ellipsis))
	// 	}
	// 	desc = strings.Join(lines, "\n")
	// }

	// Conditions
	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	if isFiltered && index < len(m.VisibleItems()) {
		// Get indices of matched characters
		matchedRunes = m.MatchesForItem(index)
	}

	if emptyFilter {
		// BRANCH -> filtering but none applied yet
		title = s.DimmedTitle.Render(title)
		desc = s.DimmedDesc.Render(desc)
	} else if isSelected && m.FilterState() != list.Filtering {
		// BRANCH -> selected but not actively filtering (can be filtered already!)
		if isFiltered {
			// Highlight matches
			unmatched := s.SelectedTitle.Inline(true)
			matched := unmatched.Copy().Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.SelectedTitle.Render(title)
		desc = s.SelectedDesc.Render(desc)
	} else {
		// BRANCH -> not selected, any other state
		if isFiltered {
			// Highlight matches
			unmatched := s.NormalTitle.Inline(true)
			matched := unmatched.Copy().Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.NormalTitle.Render(title)
		desc = s.NormalDesc.Render(desc)
	}

	fmt.Fprintf(w, "%s  %s", title, desc)
}

// ShortHelp returns the delegate's short help.
func (fd FileDelegate) ShortHelp() []key.Binding {
	if fd.ShortHelpFunc != nil {
		return fd.ShortHelpFunc()
	}
	return nil
}

// FullHelp returns the delegate's full help.
func (fd FileDelegate) FullHelp() [][]key.Binding {
	if fd.FullHelpFunc != nil {
		return fd.FullHelpFunc()
	}
	return nil
}
