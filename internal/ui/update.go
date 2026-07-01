package ui

import (
	"github.com/charmbracelet/ssh"

	tea "github.com/charmbracelet/bubbletea"
	bm "github.com/charmbracelet/wish/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.scrollPos = 0
		return m, nil

	case tea.KeyMsg:
		if m.showHelp {
			m.showHelp = false
			return m, nil
		}
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "?":
			m.showHelp = true
		case "t":
			m.toggleTheme()
		case "tab":
			if m.focus == focusNav {
				m.focus = focusContent
			} else {
				m.focus = focusNav
			}
			m.scrollPos = 0
		case "j", "down":
			if m.focus == focusNav {
				if m.cursor < len(sections)-1 {
					m.cursor++
					m.scrollPos = 0
				}
			} else {
				m.scrollPos++
			}
		case "k", "up":
			if m.focus == focusNav {
				if m.cursor > 0 {
					m.cursor--
					m.scrollPos = 0
				}
			} else {
				if m.scrollPos > 0 {
					m.scrollPos--
				}
			}
		case "enter":
			m.focus = focusContent
			m.scrollPos = 0
		case "g":
			m.scrollPos = 0
		case "G":
			m.scrollPos = 9999
		}
	}
	return m, nil
}

// TeaHandler is the wish bubbletea middleware handler.
// It uses bm.MakeRenderer(sess) to obtain a renderer tied to the SSH client's
// terminal color profile, so lipgloss emits the correct ANSI codes.
func TeaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	renderer := bm.MakeRenderer(sess)
	m := New(renderer)
	pty, _, _ := sess.Pty()
	m.width = pty.Window.Width
	m.height = pty.Window.Height
	opts := bm.MakeOptions(sess)
	opts = append(opts, tea.WithAltScreen())
	return m, opts
}
