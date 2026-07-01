package ui

import (
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

type focus int

const (
	focusNav focus = iota
	focusContent
)

type Model struct {
	cursor    int
	focus     focus
	darkMode  bool
	showHelp  bool
	width     int
	height    int
	theme     Theme
	scrollPos int
	renderer  *lipgloss.Renderer
}

// New creates a model. Pass the session renderer from bm.MakeRenderer(sess)
// so that lipgloss emits ANSI codes for the SSH client's actual color profile.
func New(r *lipgloss.Renderer) Model {
	return Model{
		cursor:   0,
		focus:    focusNav,
		darkMode: true,
		theme:    darkTheme(r),
		renderer: r,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) toggleTheme() {
	m.darkMode = !m.darkMode
	if m.darkMode {
		m.theme = darkTheme(m.renderer)
	} else {
		m.theme = lightTheme(m.renderer)
	}
}

func (m *Model) navWidth() int {
	w := m.width * 28 / 100
	if w < 20 {
		w = 20
	}
	return w
}

func (m *Model) contentWidth() int {
	return m.width - m.navWidth() - 4
}

func (m *Model) innerHeight() int {
	return m.height - 4
}
