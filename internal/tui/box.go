package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Box struct {
	width, height       int
	maxWidth, maxHeight int
	title               string
	style               lipgloss.Style
}

func NewBox(placeholder string) Box {
	return Box{
		maxWidth:  -1,
		maxHeight: -1,
		width:     2,
		height:    2,
		title:     placeholder,
		style:     blurredBorderStyle,
	}
}

func (b *Box) getEmptyLines() string {
	lines := make([]string, b.height)
	startEmpty := 0
	for i := startEmpty; i < b.height; i++ {
		lines[i] = strings.Repeat(" ", b.width)
	}

	return strings.Join(lines, "\n")
}

func (b *Box) SetMaxWidth(width int) {
	b.maxWidth = width
}

func (b *Box) SetWidth(width int) {
	if b.maxWidth <= 0 {
		b.width = width
		return
	}
	if width < b.maxWidth {
	} else {
		b.width = b.maxWidth
	}
}

func (b Box) GetWidth() int {
	return b.width
}

func (b *Box) SetMaxHeight(height int) {
	b.maxHeight = height
}

func (b *Box) SetHeight(height int) {
	if b.maxHeight <= 0 {
		b.height = height
		return
	}
	if height < b.maxHeight {
	} else {
		b.height = b.maxHeight
	}
}

func (b *Box) SetStyle(style lipgloss.Style) {
	b.style = style
}

func (b Box) Init() tea.Cmd {
	return textarea.Blink
}

func (b Box) View() string {
	w := min(b.width, b.style.GetWidth())
	h := min(b.height, b.style.GetHeight())

	w -= b.style.GetBorderStyle().GetLeftSize()
	w -= b.style.GetBorderStyle().GetRightSize()

	contents := lipgloss.NewStyle().
		Width(w).     // pad to width.
		Height(h).    // pad to height.
		MaxHeight(h). // truncate height if taller.
		MaxWidth(w).  // truncate width if wider.
		Render(b.getEmptyLines())

	style := b.style
	// topLeft := style.GetBorderStyle().TopLeft
	// topRight := style.GetBorderStyle().TopRight
	// top := style.GetBorderStyle().Top
	style = style.UnsetBorderTop()

	title := b.title
	if b.width <= 2 {
		title = "h"
	} else if len(title) >= w-2 {
		title = title[0 : w-2]
	}

	repeatedMiddleChar := w - 2 - len(title)
	if repeatedMiddleChar < 0 {
		repeatedMiddleChar = 0
	}

	topLine := strings.Builder{}
	topLine.WriteString(" ")
	topLine.WriteString(fmt.Sprintf(" %v", title))
	topLine.WriteString(strings.Repeat(" ", repeatedMiddleChar))
	topLine.WriteString(" ")

	return style.
		UnsetWidth().UnsetHeight(). // Style size already applied in contents.
		Render(fmt.Sprintf("%v\n%v", topLine.String(), contents))
}

func (b Box) Update(msg tea.Msg) (Box, tea.Cmd) {
	return b, nil
}
