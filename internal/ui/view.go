package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.width == 0 {
		return "Loading…"
	}
	if m.showHelp {
		return m.helpView()
	}
	return m.mainView()
}

func (m Model) mainView() string {
	navW := m.navWidth()
	contentW := m.contentWidth()
	innerH := m.innerHeight()

	// ── Nav panel ──────────────────────────────────────────────
	navBorder := m.theme.NavBorder
	if m.focus == focusNav {
		navBorder = focusedNavBorder(m.theme)
	}

	navTitle := m.theme.NavTitle.Render("NAVIGATION")
	navItems := []string{navTitle}
	for i, s := range sections {
		label := s.Title
		if i == m.cursor {
			navItems = append(navItems, m.theme.Selected.Width(navW-6).Render(label))
		} else {
			navItems = append(navItems, m.theme.Normal.Render(label))
		}
	}
	navContent := strings.Join(navItems, "\n")
	navPanel := navBorder.
		Width(navW).
		Height(innerH).
		Render(navContent)

	// ── Content panel ──────────────────────────────────────────
	contentBorder := m.theme.ContentBorder
	if m.focus == focusContent {
		contentBorder = focusedContentBorder(m.theme)
	}

	sec := sections[m.cursor]
	contentTitle := m.theme.ContentTitle.Render("► " + sec.Title)
	body := contentTitle + "\n\n" + m.theme.BodyText.Render(sec.Content)

	// Apply scroll
	lines := strings.Split(body, "\n")
	maxScroll := len(lines) - innerH + 2
	if maxScroll < 0 {
		maxScroll = 0
	}
	if m.scrollPos > maxScroll {
		m.scrollPos = maxScroll
	}
	start := m.scrollPos
	if start > len(lines) {
		start = len(lines)
	}
	visible := lines[start:]
	// Clamp to panel height
	if len(visible) > innerH-2 {
		visible = visible[:innerH-2]
	}
	bodyText := strings.Join(visible, "\n")

	contentPanel := contentBorder.
		Width(contentW).
		Height(innerH).
		Render(bodyText)

	// ── Join panels ────────────────────────────────────────────
	panels := lipgloss.JoinHorizontal(lipgloss.Top, navPanel, contentPanel)

	// ── Footer ─────────────────────────────────────────────────
	footerLeft := m.theme.Footer.Render("guest@portfolio")
	themeLabel := "dark"
	if !m.darkMode {
		themeLabel = "light"
	}
	focusLabel := "nav"
	if m.focus == focusContent {
		focusLabel = "content"
	}
	counter := m.theme.Counter.Render(
		fmt.Sprintf("[%s] [%s] %d/%d  [?] help",
			themeLabel, focusLabel, m.cursor+1, len(sections)),
	)
	footerGap := m.width - lipgloss.Width(footerLeft) - lipgloss.Width(counter)
	if footerGap < 1 {
		footerGap = 1
	}
	footer := footerLeft + strings.Repeat(" ", footerGap) + counter

	return lipgloss.JoinVertical(lipgloss.Left, panels, footer)
}

func (m Model) helpView() string {
	t := m.theme
	type row struct{ key, desc string }
	rows := []row{
		{"j / ↓", "next section / scroll down"},
		{"k / ↑", "prev section / scroll up"},
		{"enter", "focus content panel"},
		{"tab", "toggle panel focus"},
		{"t", "toggle light/dark theme"},
		{"g / G", "scroll to top / bottom"},
		{"?", "close this help"},
		{"q", "quit"},
	}

	var lines []string
	lines = append(lines, t.ContentTitle.Render("Help — keyboard shortcuts"))
	lines = append(lines, "")
	for _, r := range rows {
		key := t.HelpKey.Width(10).Render(r.key)
		desc := t.HelpDesc.Render(r.desc)
		lines = append(lines, "  "+key+"  "+desc)
	}
	lines = append(lines, "")
	lines = append(lines, t.Footer.Render("press any key to close"))

	box := t.ContentBorder.
		Width(m.width - 4).
		Render(strings.Join(lines, "\n"))

	topPad := (m.height - lipgloss.Height(box)) / 2
	if topPad < 0 {
		topPad = 0
	}
	return strings.Repeat("\n", topPad) + box
}
