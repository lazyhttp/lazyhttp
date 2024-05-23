package tui

import (
	"github.com/rivo/tview"
)

func Main(location string, isDirectory bool) {
	app := tview.NewApplication()
	textArea := tview.NewTextArea().SetWrap(false).SetPlaceholder("placeholder")
	textArea.SetTitle("Hello, World!").SetBorder(true)
	locationTextArea := tview.NewTextView().SetText(location)
	var directoryText string = "This is a file"
	if isDirectory {
		directoryText = "This is a directory"
	}
	directoryTextArea := tview.NewTextView().SetText(directoryText)

	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(locationTextArea, 1, 0, 1, 1, 0, 0, false).
		AddItem(directoryTextArea, 1, 1, 1, 1, 0, 0, false)

	if err := app.SetRoot(mainView, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}

}
