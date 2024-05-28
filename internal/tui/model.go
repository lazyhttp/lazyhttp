package tui

import (
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

type Model struct {
	//flags
	IsDirectory bool

	//In Args
	Location string

	//Window
	Height, Width int

	// Views
	ProgramInfo string //textarea.Model
	Requests    string //textarea.Model
	History     string //textarea.Model
	HttpMethod  string //textarea.Model
	Url         string //textarea.Model
	Headers     string //textarea.Model
	Body        string //textarea.Model
	Response    string //textarea.Model
	Statistics  string //textarea.Model
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		return tea.WindowSizeMsg{
			Width:  80,
			Height: 24,
		}
	}
}

func (m Model) View() string {

	heightUnit := m.Height / 10
	widthUnit := m.Width / 10
	leftPanel := lipgloss.JoinVertical(lipgloss.Top,
		getSizedStyle(2, 3*widthUnit).Render(m.ProgramInfo),
		getSizedStyle(3*heightUnit, 3*widthUnit).Render(m.Requests),
		getSizedStyle(3*heightUnit, 3*widthUnit).Render(m.History))

	urlPanel := lipgloss.JoinHorizontal(lipgloss.Top, getSizedStyle(2, 1*widthUnit).Render(m.HttpMethod), getSizedStyle(2, 4*widthUnit).Render(m.Url))
	requestPanel := lipgloss.JoinVertical(lipgloss.Top,
		urlPanel,
		getSizedStyle(3*heightUnit, 6*widthUnit).Render(m.Headers),
		getSizedStyle(3*heightUnit, 6*widthUnit).Render(m.Body))

	responsePanel := lipgloss.JoinVertical(lipgloss.Top,
		getSizedStyle(5*heightUnit, 4*heightUnit).Render(m.Response),
		getSizedStyle(5*heightUnit, 4*heightUnit).Render(m.Statistics))

	return lipgloss.JoinHorizontal(lipgloss.Top,
		leftPanel, requestPanel, responsePanel)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Width = msg.(tea.WindowSizeMsg).Width
		m.Height = msg.(tea.WindowSizeMsg).Height
		return m, nil
	}
	return m, nil
}
