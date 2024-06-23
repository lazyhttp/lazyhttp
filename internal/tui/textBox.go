package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInput struct {
	t       textarea.Model
	err     error
	focused bool
}

func NewTextarea() TextInput {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "Type something"
	t.ShowLineNumbers = false
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return TextInput{t: t, focused: false}
}

func (m TextInput) Init() tea.Cmd {
	return textarea.Blink
}

func (m TextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.t.Focused() {
				m.t.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.focused {
				m.focused = true
				cmd = m.t.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.t, cmd = m.t.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m TextInput) View() string {
	return m.t.View()
}

func (m *TextInput) SetWidth(width int)   { m.t.SetWidth(width) }
func (m *TextInput) SetHeight(height int) { m.t.SetHeight(height) }
func (m *TextInput) Focus() tea.Cmd {
	m.focused = true
	return m.t.Focus()
}

func (m *TextInput) Blur() {
	m.focused = false
	m.t.Blur()
}
