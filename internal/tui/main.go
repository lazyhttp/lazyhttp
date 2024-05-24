package tui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Main(location string, isDirectory bool) {
	app := tview.NewApplication()

	locationText := fmt.Sprintf("This is the specified location: %v", location)
	textArea := tview.NewTextArea().
		SetWrap(true).
		SetPlaceholder(locationText)

	textArea.
		SetBorder(true).
		SetTitle("1: Input Text Here").
		SetTitleAlign(tview.AlignLeft)

	mirroredTextView := tview.NewTextView().
		SetWrap(true)

	mirroredTextView.
		SetBorder(true).
		SetTitle("2: Mirrored Text Here").
		SetTitleAlign(tview.AlignLeft)

	helpTextView := tview.NewTextView().SetText("Press F1 for help")

	updateMirror := func() {
		currentText := textArea.GetText()
		mirroredTextView.SetText(currentText)
	}
	textArea.SetChangedFunc(updateMirror)
	updateMirror()

	var directoryText = "This is a file"
	if isDirectory {
		directoryText = "This is a directory"
	}
	directoryTextView := tview.NewTextView().SetText(directoryText)

	pages := tview.NewPages()

	mainView := tview.NewGrid().
		SetRows(0, 1).
		//SetColumns(0, 1).
		AddItem(textArea, 0, 0, 1, 1, 0, 0, true).
		AddItem(mirroredTextView, 0, 1, 1, 1, 0, 0, false).
		AddItem(helpTextView, 1, 0, 1, 1, 0, 0, false).
		AddItem(directoryTextView, 1, 1, 1, 1, 0, 0, false)

	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(helpText1)

	help2 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(helpText2)

	help3 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(helpText3)

	help := tview.NewFrame(help1).
		SetBorders(1, 1, 0, 0, 2, 2)
	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	pages.AddAndSwitchToPage("main", mainView, true).
		AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help") //TODO: Check when clicking outside help window with the mouse. Then clicking help again.
			return nil
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

var helpText1 string = `[green]Navigation

[yellow]Left arrow[white]: Move left.
[yellow]Right arrow[white]: Move right.
[yellow]Down arrow[white]: Move down.
[yellow]Up arrow[white]: Move up.
[yellow]Ctrl-A, Home[white]: Move to the beginning of the current line.
[yellow]Ctrl-E, End[white]: Move to the end of the current line.
[yellow]Ctrl-F, page down[white]: Move down by one page.
[yellow]Ctrl-B, page up[white]: Move up by one page.
[yellow]Alt-Up arrow[white]: Scroll the page up.
[yellow]Alt-Down arrow[white]: Scroll the page down.
[yellow]Alt-Left arrow[white]: Scroll the page to the left.
[yellow]Alt-Right arrow[white]:  Scroll the page to the right.
[yellow]Alt-B, Ctrl-Left arrow[white]: Move back by one word.
[yellow]Alt-F, Ctrl-Right arrow[white]: Move forward by one word.

[blue]Press Enter for more help, press Escape to return.`

var helpText2 string = `[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`

var helpText3 string = `[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.
[yellow]Ctrl-L[white] to select entire text.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.
		
[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`
