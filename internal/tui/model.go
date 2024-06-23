package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
---------------------------------------------------------
|Info	|Url            			|Response			|
|-----------------------------------|					|
|Reqs	|Headers					|					|
|		|							|					|
|		|							|					|
|-------|---------------------------|-------------------|
|History|Body						|Stats				|
|		|							|					|
|		|							|					|
|		|							|					|
|-------------------------------------------------------|
*/

var (
	infoHeight int = 1
	helpHeight int = 1
	numberCols int = 3
	margin     int = 3
)

type Tui struct {
	// flags
	IsDirectory bool

	// In Args
	Location string

	// Window
	Height, Width int

	// Views
	ProgramInfo Box
	Requests    Box
	History     Box

	Url     TextInput
	Headers HeaderTable
	Body    TextInput

	Response Box

	Help help.Model
}

func initialModel() *Tui {
	model := Tui{
		IsDirectory: false,
		Location:    ".",
		Response:    NewBox("help"),
		ProgramInfo: NewBox("help"),
		Url:         NewTextarea(),
		Requests:    NewBox("help"),
		Headers:     NewHeaderTable(),
		Body:        NewTextarea(),

		History: NewBox("help"),
		Help:    help.New(),
	}

	model.ProgramInfo.SetMaxWidth(25)
	model.Requests.SetMaxWidth(25)
	model.History.SetMaxWidth(25)

	model.Body.Focus()
	return &model
}

func (m *Tui) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Tui) View() string {
	vertLeft := lipgloss.JoinVertical(lipgloss.Left,
		m.ProgramInfo.View(),
		m.Requests.View(),
		m.History.View(),
	)

	vertMiddle := lipgloss.JoinVertical(lipgloss.Left,
		m.Url.View(),
		m.Headers.View(),
		m.Body.View(),
	)

	horizontalViews := lipgloss.JoinHorizontal(
		lipgloss.Top,
		vertLeft,
		vertMiddle,
		m.Response.View(),
	)
	return horizontalViews + "\n" + blurredBorderStyle.Render("help ?")
}

func (m *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

		m.ProgramInfo.SetHeight(infoHeight)
		leftHeight := msg.Height - helpHeight - infoHeight
		m.Requests.SetHeight(leftHeight/2 - 4)
		m.History.SetHeight(leftHeight/2 - 3)

		m.Url.SetHeight(infoHeight)

		m.Headers.SetHeight(leftHeight/2 - 2)
		m.Body.SetHeight(leftHeight/2 - 3)

		m.Response.SetHeight(msg.Height - helpHeight - 4)

		m.ProgramInfo.SetWidth(msg.Width/2 - 2)
		m.Requests.SetWidth(msg.Width/2 - 2)
		m.History.SetWidth(msg.Width/2 - 2)
		leftSideWidth := m.ProgramInfo.GetWidth()

		m.Url.SetWidth((msg.Width-leftSideWidth)/2 - 3)

		m.Headers.SetWidth((msg.Width-leftSideWidth)/2 - 1)
		m.Body.SetWidth((msg.Width-leftSideWidth)/2 - 3)

		m.Response.SetWidth((msg.Width-leftSideWidth)/2 - 3)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	// newPI, cmdi := m.ProgramInfo.Update(msg)
	// newUrl, cmdu := m.Url.Update(msg)
	// newRes, cmdr := m.Response.Update(msg)
	//
	// m.ProgramInfo = newPI
	// m.Url = newUrl
	// m.Response = newRes
	// cmds = append(cmds, cmdi, cmdu, cmdr)

	return m, tea.Batch(cmds...)
}
