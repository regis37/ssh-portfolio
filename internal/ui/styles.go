package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	// FocusColor is the border color when a panel has focus; adapts per theme.
	FocusColor lipgloss.Color

	NavBorder     lipgloss.Style
	ContentBorder lipgloss.Style
	NavTitle      lipgloss.Style
	ContentTitle  lipgloss.Style
	Selected      lipgloss.Style
	Normal        lipgloss.Style
	// BodyText styles raw section content (fg + bg) so panel bg is consistent.
	BodyText lipgloss.Style
	Footer   lipgloss.Style
	Counter  lipgloss.Style
	HelpKey  lipgloss.Style
	HelpDesc lipgloss.Style
}

func darkTheme() Theme {
	panelBg    := lipgloss.Color("#0D0B1E")
	borderCol  := lipgloss.Color("#3D3570")
	focusCol   := lipgloss.Color("#9B78FF")
	titleCol   := lipgloss.Color("#9B78FF")
	textCol    := lipgloss.Color("#C9C1FF")
	mutedCol   := lipgloss.Color("#5C5190")
	selectedBg := lipgloss.Color("#7D56F4")

	return Theme{
		FocusColor: focusCol,
		NavBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		ContentBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		NavTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		ContentTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		Selected: lipgloss.NewStyle().
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg).
			Padding(0, 1),
		BodyText: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg),
		Footer: lipgloss.NewStyle().
			Foreground(mutedCol),
		Counter: lipgloss.NewStyle().
			Foreground(mutedCol),
		HelpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		HelpDesc: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg),
	}
}

func lightTheme() Theme {
	panelBg    := lipgloss.Color("#EEE8FF")
	borderCol  := lipgloss.Color("#B4A8E0")
	focusCol   := lipgloss.Color("#5A30D4")
	titleCol   := lipgloss.Color("#4520AA")
	textCol    := lipgloss.Color("#1A0F3C")
	mutedCol   := lipgloss.Color("#7060AA")
	selectedBg := lipgloss.Color("#5A30D4")

	return Theme{
		FocusColor: focusCol,
		NavBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		ContentBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		NavTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		ContentTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		Selected: lipgloss.NewStyle().
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg).
			Padding(0, 1),
		BodyText: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg),
		Footer: lipgloss.NewStyle().
			Foreground(mutedCol),
		Counter: lipgloss.NewStyle().
			Foreground(mutedCol),
		HelpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		HelpDesc: lipgloss.NewStyle().
			Foreground(textCol).
			Background(panelBg),
	}
}

// focusedNavBorder returns the nav border style with the theme's focus color.
func focusedNavBorder(t Theme) lipgloss.Style {
	return t.NavBorder.BorderForeground(t.FocusColor)
}

// focusedContentBorder returns the content border style with the theme's focus color.
func focusedContentBorder(t Theme) lipgloss.Style {
	return t.ContentBorder.BorderForeground(t.FocusColor)
}
