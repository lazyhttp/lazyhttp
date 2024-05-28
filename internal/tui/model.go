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
	ProgramInfo relativeSizedView
	Requests    relativeSizedView
	History     relativeSizedView
	HttpMethod  relativeSizedView
	Url         relativeSizedView
	Headers     relativeSizedView
	Body        relativeSizedView
	Response    relativeSizedView
	Statistics  relativeSizedView
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {

	leftPanel := lipgloss.JoinVertical(lipgloss.Top,
		m.ProgramInfo.View(),
		m.Requests.View(),
		m.History.View())

	urlPanel := lipgloss.JoinHorizontal(lipgloss.Top, m.HttpMethod.View(), m.Url.View())
	requestPanel := lipgloss.JoinVertical(lipgloss.Top,
		urlPanel,
		m.Headers.View(),
		m.Body.View())

	responsePanel := lipgloss.JoinVertical(lipgloss.Top,
		m.Response.View(),
		m.Statistics.View())

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
		m.ProgramInfo.Update(msg)
		m.Requests.Update(msg)
		m.History.Update(msg)
		m.HttpMethod.Update(msg)
		m.Url.Update(msg)
		m.Headers.Update(msg)
		m.Body.Update(msg)
		m.Response.Update(msg)
		m.Statistics.Update(msg)
		return m, nil
	}
	return m, nil
}
