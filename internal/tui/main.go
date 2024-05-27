package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lazyhttp/lazyhttp/internal/requests"
	"os"
)

func MainPage(location string, isDirectory bool) {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	//flags
	IsDirectory bool

	//In Args
	Location string

	// Views
	ProgramInfo textarea.Model
	Collection  textarea.Model
	Recent      textarea.Model
	HttpMethod  textarea.Model
	Url         textarea.Model
	Headers     textarea.Model
	Body        textarea.Model
	Response    textarea.Model
	Statistics  textarea.Model
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m Model) Init() tea.Cmd { return nil }

func (m Model) View() string {
	return baseStyle.Render(m.Body.View()) + "\n"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func initialModel() Model {
	model := Model{
		IsDirectory: false,
		Location:    ".",
		ProgramInfo: makeTextArea("info"),
		Collection:  makeTextArea("collection"),
		Recent:      makeTextArea("recent"),
		HttpMethod:  makeTextArea("method"),
		Url:         makeTextArea("url"),
		Headers:     makeTextArea("header"),
		Body:        makeTextArea("body"),
		Response:    makeTextArea("response"),
		Statistics:  makeTextArea("stats")}
	return model
}

func fireRequest(method, url string) (response string, err error) {

	requestType := method
	requestUrl := url
	if requestType == "GET" {
		response, err := requests.Get(requestUrl)
		if err != nil {
			return "", fmt.Errorf("failed GET request to %v; %w", requestUrl, err)
		}
		return response, nil
	}
	return "NOT SUPPORTED", nil
}

func makeTextArea(text string) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = "Flag description"
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))
	ta.SetValue(text)
	ta.ShowLineNumbers = false
	ta.Blur()
	return ta
}
