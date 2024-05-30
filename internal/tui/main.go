package tui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/lazyhttp/lazyhttp/internal/requests"
	"github.com/rivo/tview"
)

func MainPage(location string, isDirectory bool) {
	app := tview.NewApplication()

	infoBox := tview.NewTextView()
	infoBox.
		SetBorder(true).
		SetTitle("info").
		SetTitleAlign(tview.AlignLeft)

	requestDirectory := tview.NewTextView()
	requestDirectory.
		SetBorder(true).
		SetTitle("requests").
		SetTitleAlign(tview.AlignLeft)

	recents := tview.NewTextView()
	recents.
		SetBorder(true).
		SetTitle("recents").
		SetTitleAlign(tview.AlignLeft)

	methodSelect := tview.NewDropDown().
		SetOptions([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}, nil).
		SetCurrentOption(0)

	methodSelect.
		SetBorder(true).
		SetTitle("Method").
		SetTitleAlign(tview.AlignLeft)

	urlTextArea := tview.NewTextArea().
		SetWrap(true).
		SetText("https://www.google.com", false)

	urlTextArea.
		SetBorder(true).
		SetTitle("1: Input Text Here").
		SetTitleAlign(tview.AlignLeft)

	headersTextArea := tview.NewTextArea().
		SetWrap(true).
		SetText("headers", false)

	headersTextArea.
		SetBorder(true).
		SetTitle("headers").
		SetTitleAlign(tview.AlignLeft)

	bodyTextArea := tview.NewTextArea().
		SetWrap(true).SetText("SomeBody", false)

	responseTextView := tview.NewTextView().
		SetWrap(true)

	responseTextView.
		SetBorder(true).
		SetTitle("2: Mirrored Text Here").
		SetTitleAlign(tview.AlignLeft)

	helpTextView := tview.NewTextView().SetText("Press F1 for help")

	mainView := tview.NewGrid().
		SetRows(2, 0, 0, 1).
		SetColumns(1, 2, 0, 0).
		//left side
		AddItem(infoBox, 0, 0, 1, 1, 0, 0, true).
		AddItem(requestDirectory, 1, 0, 1, 1, 0, 0, false).
		AddItem(recents, 2, 0, 1, 1, 0, 0, false).
		// middle
		AddItem(methodSelect, 0, 1, 1, 1, 0, 0, false).
		AddItem(urlTextArea, 0, 2, 1, 1, 0, 0, false).
		AddItem(headersTextArea, 1, 1, 1, 2, 0, 0, false).
		AddItem(bodyTextArea, 2, 1, 1, 2, 0, 0, false).
		//right
		AddItem(responseTextView, 0, 3, 3, 1, 0, 0, false).
		//help
		AddItem(helpTextView, 3, 1, 1, 3, 0, 0, false)

	pages := tview.NewPages()
	help := makeHelp(pages)

	pages.AddAndSwitchToPage("main", mainView, true).
		AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF1:
			pages.ShowPage("help") //TODO: Check when clicking outside help window with the mouse. Then clicking help again.
			return nil

		case tcell.KeyCR:
		case tcell.KeyF5:
			if event.Key() == tcell.KeyCR && event.Modifiers() != tcell.ModCtrl {
				return event
			}

			_, method := methodSelect.GetCurrentOption()
			respose, requestError := fireRequest(method, urlTextArea.GetText())

			if requestError != nil {
				responseTextView.SetText(requestError.Error())
				return nil
			}
			responseTextView.SetText(respose)

			return nil
		}

		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
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

func makeHelp(pages *tview.Pages) *tview.Frame {
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
	return help
}

var helpText1 = `[green]Navigation

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

var helpText2 = `[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`

var helpText3 = `[green]Selecting Text[white]

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
