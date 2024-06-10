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
	infoHeight int = 2
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
	ProgramInfo textarea.Model
	Url         textarea.Model
	Response    textarea.Model
	Help        help.Model

	Requests textarea.Model
	// History     relativeSizedView
	// HttpMethod textarea.Model
	Headers textarea.Model
	// Body        relativeSizedView
	// Statistics  relativeSizedView
}

func initialModel() *Model {
	model := Model{
		IsDirectory: false,
		Location:    ".",
		Response:    newTextarea(),
		ProgramInfo: newTextarea(),
		Url:         newTextarea(),
		Requests:    newTextarea(),
		Headers:     newTextarea(),
		// HttpMethod:  newTextarea(),
		Help: help.New(),
	}
	model.Url.Focus()
	return &model
}

func newTextarea() textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "Type something"
	t.ShowLineNumbers = true
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return t
}

func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) View() string {
	vertLeft := lipgloss.JoinVertical(lipgloss.Left,
		m.ProgramInfo.View(),
		m.Requests.View(),
	)

	vertMiddle := lipgloss.JoinVertical(lipgloss.Left,
		m.Url.View(),
		m.Headers.View(),
	)

	horizontalViews := lipgloss.JoinHorizontal(
		lipgloss.Top,
		vertLeft,
		vertMiddle,
		m.Response.View(),
	)
	return horizontalViews + "\n\n" + "help ?"
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

		m.ProgramInfo.SetHeight(infoHeight)
		m.Requests.SetHeight(msg.Height - helpHeight - infoHeight - 6)

		m.Url.SetHeight(infoHeight)
		m.Headers.SetHeight(msg.Height - helpHeight - infoHeight - 6)

		m.Response.SetHeight(msg.Height - helpHeight - 4)

		m.ProgramInfo.SetWidth(msg.Width / numberCols)
		m.Requests.SetWidth(msg.Width / numberCols)

		m.Url.SetWidth(msg.Width / numberCols)
		m.Headers.SetWidth(msg.Width / numberCols)

		m.Response.SetWidth(msg.Width / numberCols)

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
