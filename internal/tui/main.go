package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lazyhttp/lazyhttp/internal/requests"
	"os"
)

func MainPage(location string, isDirectory bool) {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() Model {
	model := Model{
		IsDirectory: false,
		Location:    ".",
		ProgramInfo: relativeSizedView{1.0 / 6.0, 1.0 / 4.0, 1, 1, "info"},
		Requests:    relativeSizedView{3.0 / 6.0, 1.0 / 4, 1, 1, "collection"},
		History:     relativeSizedView{3.0 / 6.0, 1.0 / 4, 1, 1, "recent"},
		HttpMethod:  relativeSizedView{1.0 / 6, 1.0 / 4, 1, 1, "method"},
		Url:         relativeSizedView{1.0 / 6, 1.0 / 4, 1, 1, "url"},
		Headers:     relativeSizedView{3.0 / 6, 2.0 / 4, 1, 1, "header"},
		Body:        relativeSizedView{3.0 / 6, 2.0 / 4, 1, 1, "body"},
		Response:    relativeSizedView{4.0 / 6, 1.0 / 4, 1, 1, "response"},
		Statistics:  relativeSizedView{3.0 / 6, 1.0 / 4, 1, 1, "stats"}}
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
