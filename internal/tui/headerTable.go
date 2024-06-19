package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyHeader = "he"
	columnKeyValue  = "va"

	metaDataEnabledKey = "enabled"
)

var (
	enabledStyle  = baseStyle
	disabledStyle = baseStyle.Strikethrough(true)
)

type HeaderRows struct {
	Header, Value string
	Enabled       bool
}

func (r HeaderRows) toRow() table.Row {
	return table.NewRow(table.RowData{
		columnKeyHeader:    r.Header,
		columnKeyValue:     r.Value,
		metaDataEnabledKey: r.Enabled,
	})
}

type HeaderTable struct {
	t table.Model

	height, width int
	data          []*HeaderRows

	style lipgloss.Style
}

func NewHeaderTable() HeaderTable {
	rows := []*HeaderRows{
		{"H1", "V1", true},
	}

	r := make([]table.Row, len(rows))
	for i := range rows {
		r[i] = rows[i].toRow()
	}

	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	t := table.New([]table.Column{
		table.NewColumn(columnKeyHeader, "Header", 10),
		table.NewFlexColumn(columnKeyValue, "Value", 1),
	}).
		WithRows(r).
		WithBaseStyle(style).
		BorderRounded().
		WithRowStyleFunc(func(rsfi table.RowStyleFuncInput) lipgloss.Style {
			if rsfi.Row.Data[metaDataEnabledKey] == false {
				return disabledStyle
			}
			return enabledStyle
		})

	m := HeaderTable{t: t, data: rows, style: style}
	return m
}

func (ht *HeaderTable) SetStyle(style lipgloss.Style) {
	ht.style = style
}

func (ht *HeaderTable) SetHeight(height int) {
	ht.height = height
}

func (ht *HeaderTable) SetWidth(width int) {
	ht.width = width
}

func (ht *HeaderTable) AddRows(rows ...*HeaderRows) {
}

func (ht *HeaderTable) Init() tea.Cmd {
	return nil
}

func (ht *HeaderTable) View() string {
	t := ht.t.
		WithMinimumHeight(ht.height).
		WithTargetWidth(ht.width)
	return t.View()
}

func (ht *HeaderTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return ht, nil
}
