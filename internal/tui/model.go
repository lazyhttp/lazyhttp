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
|Info	|Method		|Url			|Response			|
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

type Model struct {
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

	HttpMethod Box
	Url        textarea.Model
	Headers    HeaderTable
	Body       textarea.Model

	Response Box

	Help help.Model
}

func initialModel() *Model {
	model := Model{
		IsDirectory: false,
		Location:    ".",
		Response:    NewBox("coming soon"),
		ProgramInfo: NewBox("coming soon"),
		Url:         NewTextarea(),
		Requests:    NewBox("coming soon"),
		Headers:     NewHeaderTable(),
		HttpMethod:  NewBox("DELETE"),
		Body:        NewTextarea(),

		History: NewBox("coming soon"),
		Help:    help.New(),
	}

	model.ProgramInfo.SetMaxWidth(25)
	model.Requests.SetMaxWidth(25)
	model.History.SetMaxWidth(25)

	model.Url.Focus()
	return &model
}

func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) View() string {
	vertLeft := lipgloss.JoinVertical(lipgloss.Left,
		m.ProgramInfo.View(),
		m.Requests.View(),
		m.History.View(),
	)

	requestMiddle := lipgloss.JoinHorizontal(lipgloss.Top,
		m.HttpMethod.View(),
		m.Url.View())

	vertMiddle := lipgloss.JoinVertical(lipgloss.Left,
		requestMiddle,
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

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

		m.ProgramInfo.SetHeight(infoHeight)
		leftHeight := msg.Height - helpHeight - infoHeight
		m.Requests.SetHeight(leftHeight/2 - 4)
		m.History.SetHeight(leftHeight/2 - 4)

		m.HttpMethod.SetHeight(infoHeight)
		m.Url.SetHeight(infoHeight)

		m.Headers.SetHeight(leftHeight/2 - 2)
		m.Body.SetHeight(leftHeight/2 - 4)

		m.Response.SetHeight(msg.Height - helpHeight - 4)

		m.ProgramInfo.SetWidth(msg.Width/2 - 2)
		m.Requests.SetWidth(msg.Width/2 - 2)
		m.History.SetWidth(msg.Width/2 - 2)
		leftSideWidth := m.ProgramInfo.GetWidth()

		methodWidth := 10
		m.HttpMethod.SetWidth(methodWidth)
		m.Url.SetWidth((msg.Width-leftSideWidth)/2 - methodWidth - 5)

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
	newPI, cmdi := m.ProgramInfo.Update(msg)
	newUrl, cmdu := m.Url.Update(msg)
	newRes, cmdr := m.Response.Update(msg)

	m.ProgramInfo = newPI
	m.Url = newUrl
	m.Response = newRes
	cmds = append(cmds, cmdi, cmdu, cmdr)

	return m, tea.Batch(cmds...)
}
