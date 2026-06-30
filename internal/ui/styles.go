package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Base        lipgloss.Style
	NavBorder   lipgloss.Style
	ContentBorder lipgloss.Style
	NavTitle    lipgloss.Style
	ContentTitle lipgloss.Style
	Selected    lipgloss.Style
	Normal      lipgloss.Style
	Footer      lipgloss.Style
	Counter     lipgloss.Style
	HelpKey     lipgloss.Style
	HelpDesc    lipgloss.Style
	Accent      lipgloss.Style
}

func darkTheme() Theme {
	selectedBg := lipgloss.Color("#7D56F4")
	accentColor := lipgloss.Color("#7D56F4")
	borderColor := lipgloss.Color("#444444")
	mutedColor := lipgloss.Color("#666666")
	textColor := lipgloss.Color("#DDDDDD")
	bgColor := lipgloss.Color("#1a1a2e")

	_ = bgColor

	return Theme{
		Base: lipgloss.NewStyle().
			Foreground(textColor),
		NavBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(0, 1),
		ContentBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(0, 1),
		NavTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor).
			MarginBottom(1),
		ContentTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor).
			MarginBottom(1),
		Selected: lipgloss.NewStyle().
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 1),
		Footer: lipgloss.NewStyle().
			Foreground(mutedColor),
		Counter: lipgloss.NewStyle().
			Foreground(mutedColor),
		HelpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor),
		HelpDesc: lipgloss.NewStyle().
			Foreground(textColor),
		Accent: lipgloss.NewStyle().
			Foreground(accentColor),
	}
}

func lightTheme() Theme {
	selectedBg := lipgloss.Color("#5A3FD4")
	accentColor := lipgloss.Color("#5A3FD4")
	borderColor := lipgloss.Color("#AAAAAA")
	mutedColor := lipgloss.Color("#888888")
	textColor := lipgloss.Color("#222222")

	return Theme{
		Base: lipgloss.NewStyle().
			Foreground(textColor),
		NavBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(0, 1),
		ContentBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(0, 1),
		NavTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor).
			MarginBottom(1),
		ContentTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor).
			MarginBottom(1),
		Selected: lipgloss.NewStyle().
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 1),
		Footer: lipgloss.NewStyle().
			Foreground(mutedColor),
		Counter: lipgloss.NewStyle().
			Foreground(mutedColor),
		HelpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor),
		HelpDesc: lipgloss.NewStyle().
			Foreground(textColor),
		Accent: lipgloss.NewStyle().
			Foreground(accentColor),
	}
}

func focusedNavBorder(t Theme) lipgloss.Style {
	return t.NavBorder.BorderForeground(lipgloss.Color("#7D56F4"))
}

func focusedContentBorder(t Theme) lipgloss.Style {
	return t.ContentBorder.BorderForeground(lipgloss.Color("#7D56F4"))
}
