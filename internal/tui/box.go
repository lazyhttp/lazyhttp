package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Box struct {
	width, height int
	text          string
	style         lipgloss.Style
}

func NewBox() Box {
	return Box{
		width: 5, height: 3,
		text:  "coming soon",
		style: focusedPlaceholderStyle,
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
	if diff < 0 {
		startEmpty = 2
		s := strings.Split(b.text, " ")
		firstLine := fmt.Sprintf("%[1]*s", -b.width, fmt.Sprintf("%[1]*s", (b.width+len(s[0]))/2, s[0]))
		secondLine := fmt.Sprintf("%[1]*s", -b.width, fmt.Sprintf("%[1]*s", (b.width+len(s[0]))/2, s[0]))
		lines[0] = firstLine
		lines[1] = secondLine
	}

	for i := startEmpty; i < b.height; i++ {
		lines[i] = strings.Repeat(" ", b.width)
	}

	return strings.Join(lines, "\n")
}

func (b *Box) SetWidth(width int) {
	b.width = width
}

func (b *Box) SetHeight(height int) {
	b.height = height
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
