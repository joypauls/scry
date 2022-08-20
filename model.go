package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joypauls/scry/fst"
)

// Model: state of the program
type model struct {
	list      list.Model
	directory *fst.Directory
}

// need to define this
func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update: modifying model based on input
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Switch for key press events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	// Updating loop
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View: logic to display model
func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func newModel() model {

	// Read current directory as a test
	path := fst.NewPath("")
	d := fst.NewDirectory(path, false)

	var numItems = d.Size()
	// elements need to implement Item interface here
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = MakeItem(d.File(i))
	}

	// Create and style list
	delegate := NewFileDelegate()
	// delegate := list.NewDefaultDelegate()
	fileList := list.New(items, delegate, 0, 0)
	fileList.Title = "Files"
	fileList.Styles.Title = titleStyle

	return model{
		list:      fileList,
		directory: d,
	}
}
