package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"math"
)

type relativeSizedView struct {
	heightPerc, widthPerc float32
	height, width         int
	text                  string
}

// Init does initial setup for the column.
func (v relativeSizedView) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O for columns.
func (v relativeSizedView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v.setSize(msg.Width, msg.Height)
	}
	return v, cmd
}

func (v relativeSizedView) View() string {
	return baseStyle.
		Height(v.height).
		Width(v.width).
		Render(v.text)
}

func (v relativeSizedView) setSize(width int, height int) {
	v.width = int(math.Floor(float64(v.widthPerc) * float64(width)))
	v.height = int(math.Floor(float64(height) * float64(v.heightPerc)))
}
