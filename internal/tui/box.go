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
	text                string
	style               lipgloss.Style
}

func NewBox(placeholder string) Box {
	return Box{
		maxWidth:  -1,
		maxHeight: -1,
		width:     2,
		height:    2,
		text:      placeholder,
		style:     blurredBorderStyle,
	}
}

func (b *Box) visableLines() string {
	diff := b.width - len(b.text)

	lines := make([]string, b.height)
	startEmpty := 0
	if diff > 2 {
		startEmpty = 1
		firstLine := fmt.Sprintf("%[1]*s", -b.width, fmt.Sprintf("%[1]*s", (b.width+len(b.text))/2, b.text))
		lines[0] = firstLine

	}
	if diff < 0 && strings.Contains(b.text, " ") {
		startEmpty = 2
		s := strings.Split(b.text, " ")
		firstLine := fmt.Sprintf("%[1]*s", -b.width, fmt.Sprintf("%[1]*s", (b.width+len(s[0]))/2, s[0]))
		secondLine := fmt.Sprintf("%[1]*s", -b.width, fmt.Sprintf("%[1]*s", (b.width+len(s[1]))/2, s[1]))
		lines[0] = firstLine
		lines[1] = secondLine
	}

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
		Render(b.visableLines())

	return b.style.
		UnsetWidth().UnsetHeight(). // Style size already applied in contents.
		Render(contents)
}

func (b Box) Update(msg tea.Msg) (Box, tea.Cmd) {
	return b, nil
}
