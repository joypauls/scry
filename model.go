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
	path      *fst.Path
}

// need to define this for interface
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

	// Update component(s)
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
	p := fst.NewPath("")
	d := fst.NewDirectory(p, false)

	// Elements need to implement Item interface here
	listItems := make([]list.Item, d.Size())
	for i := 0; i < d.Size(); i++ {
		listItems[i] = MakeFileItem(d.File(i))
	}

	// Create and style list
	fd := NewFileDelegate()
	files := list.New(listItems, fd, 0, 0)

	files.Title = formatHeader(p, 50)
	// files.Title = "Files"
	files.Styles.Title = titleStyle
	files.DisableQuitKeybindings() // handle quit in Update()
	files.SetWidth(50)

	return model{
		list:      files,
		directory: d,
		path:      p,
	}
}
