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
		ProgramInfo: "info",
		Requests:    "collection",
		History:     "recent",
		HttpMethod:  "method",
		Url:         "url",
		Headers:     "header",
		Body:        "body",
		Response:    "response",
		Statistics:  "stats"}
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
