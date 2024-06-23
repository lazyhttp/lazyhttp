package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Log(message string) {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}

		_, erw := f.WriteString(message)

		if erw != nil {
			fmt.Println("fatal:", erw)
			os.Exit(1)
		}
		defer f.Close()
	}
}
