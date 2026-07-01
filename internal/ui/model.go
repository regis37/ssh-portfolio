package ui

import tea "github.com/charmbracelet/bubbletea"

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
}

func New() Model {
	return Model{
		cursor:   0,
		focus:    focusNav,
		darkMode: true,
		theme:    darkTheme(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) toggleTheme() {
	m.darkMode = !m.darkMode
	if m.darkMode {
		m.theme = darkTheme()
	} else {
		m.theme = lightTheme()
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
