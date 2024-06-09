package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HeaderLine struct {
	key, value string
	enabled    bool
	style      string
}

func NewHeaderLine(key string, value string, enabled bool, style string) *HeaderLine {
	return &HeaderLine{
		key: key, style: style, value: value, enabled: enabled,
	}
}

func (h *HeaderLine) GetCells() [3]*tview.TableCell {
	cells := [3]*tview.TableCell{}
	key := tview.NewTableCell(h.key)
	key.SetAlign(tview.AlignLeft)

	cells[0] = key

	value := tview.NewTableCell(h.value)
	value.SetAlign(tview.AlignRight)
	value.SetExpansion(1)
	cells[1] = value

	checkBox := tview.NewTableCell(h.getCheckbox())
	checkBox.SetClickedFunc(func() bool {
		if h.enabled {
			h.enabled = false
		} else {
			h.enabled = true
		}

		checkBox.SetText(h.getCheckbox())
		return false
	})
	cells[2] = checkBox

	return cells
}

func (h HeaderLine) getCheckbox() string {
	checkBox := "\u2610  "
	if h.enabled {
		checkBox = "\u2611  "
	}

	return checkBox
}

type headerTableData []*HeaderLine

type HeaderTable struct {
	tview.Table
	tableData *headerTableData
	hasFocus  bool
}

func NewHeaderTable() *HeaderTable {
	td := make(headerTableData, 0)
	t := tview.NewTable()
	//Input capture for t focus management
	table := HeaderTable{
		Table:     *t,
		tableData: &td,
	}

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp, tcell.KeyDown, tcell.KeyRight, tcell.KeyLeft:
			table.SetSelectable(true, false)
		case tcell.KeyEscape:
			table.SetSelectable(false, false)
		}
		return event
	})

	return &table
}

func (h *HeaderTable) AddRow(line *HeaderLine) {
	*h.tableData = append(*h.tableData, line)
	cells := line.GetCells()
	row := h.GetRowCount()
	for col, cell := range cells {
		h.SetCell(row+1, col, cell)
	}
}

func (ht *HeaderTable) Focus(delegate func(p tview.Primitive)) {
	ht.hasFocus = true
	ht.Table.SetSelectable(true, false)
	ht.Table.Focus(delegate)
}

func (ht *HeaderTable) Blur() {
	ht.hasFocus = false
	ht.Table.SetSelectable(false, false)
	ht.Table.Blur()
}
