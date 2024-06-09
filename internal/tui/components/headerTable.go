package components

import (
	"slices"

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

type headerTableData []*HeaderLine

func (h *headerTableData) GetCell(row, column int) *tview.TableCell {
	line := (*h)[row]
	switch column {
	case 0:
		return tview.NewTableCell(line.key)
	case 1:
		return tview.NewTableCell(line.value)
	case 2:
		enabledString := "\u2610"
		if line.enabled {
			enabledString = "\u2611"
		}
		return tview.NewTableCell(enabledString)
	}
	return tview.NewTableCell("")
}

func (h *headerTableData) GetRowCount() int {
	return len(*h)
}

func (h *headerTableData) GetColumnCount() int {
	return 3
}

func (h *headerTableData) SetCell(row, column int, cell *tview.TableCell) {
}

func (h *headerTableData) RemoveRow(row int) {
	*h = slices.Delete(*h, row, row)
}

func (h *headerTableData) RemoveColumn(column int) {
}

func (h *headerTableData) InsertRow(row int) {
	*h = append(*h, &HeaderLine{})
}

func (h *headerTableData) InsertColumn(column int) {
}

func (h *headerTableData) Clear() {
	*h = make(headerTableData, 0)
}

type HeaderTable struct {
	tview.Table
	tableData *headerTableData
}

func NewHeaderTable() *HeaderTable {
	td := make(headerTableData, 0)
	t := tview.NewTable()
	table := HeaderTable{
		Table:     *t,
		tableData: &td,
	}

	return &table
}

func (h *HeaderTable) AddRow(line *HeaderLine) {
	*h.tableData = append(*h.tableData, line)
	(*h).SetContent((*h).tableData)
}
