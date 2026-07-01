package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	FocusColor    lipgloss.Color
	NavBorder     lipgloss.Style
	ContentBorder lipgloss.Style
	NavTitle      lipgloss.Style
	ContentTitle  lipgloss.Style
	Selected      lipgloss.Style
	Normal        lipgloss.Style
	BodyText      lipgloss.Style
	Footer        lipgloss.Style
	Counter       lipgloss.Style
	HelpKey       lipgloss.Style
	HelpDesc      lipgloss.Style
}

// ns returns a new Style tied to the given renderer.
// If r is nil it falls back to the lipgloss default renderer.
func ns(r *lipgloss.Renderer) lipgloss.Style {
	if r != nil {
		return r.NewStyle()
	}
	return lipgloss.NewStyle()
}

func darkTheme(r *lipgloss.Renderer) Theme {
	panelBg    := lipgloss.Color("#0D0B1E")
	borderCol  := lipgloss.Color("#3D3570")
	focusCol   := lipgloss.Color("#9B78FF")
	titleCol   := lipgloss.Color("#9B78FF")
	textCol    := lipgloss.Color("#C9C1FF")
	mutedCol   := lipgloss.Color("#5C5190")
	selectedBg := lipgloss.Color("#7D56F4")

	return Theme{
		FocusColor: focusCol,
		NavBorder: ns(r).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		ContentBorder: ns(r).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		NavTitle: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		ContentTitle: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		Selected: ns(r).
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: ns(r).
			Foreground(textCol).
			Background(panelBg).
			Padding(0, 1),
		BodyText: ns(r).
			Foreground(textCol).
			Background(panelBg),
		Footer: ns(r).
			Foreground(mutedCol),
		Counter: ns(r).
			Foreground(mutedCol),
		HelpKey: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		HelpDesc: ns(r).
			Foreground(textCol).
			Background(panelBg),
	}
}

func lightTheme(r *lipgloss.Renderer) Theme {
	panelBg    := lipgloss.Color("#EEE8FF")
	borderCol  := lipgloss.Color("#B4A8E0")
	focusCol   := lipgloss.Color("#5A30D4")
	titleCol   := lipgloss.Color("#4520AA")
	textCol    := lipgloss.Color("#1A0F3C")
	mutedCol   := lipgloss.Color("#7060AA")
	selectedBg := lipgloss.Color("#5A30D4")

	return Theme{
		FocusColor: focusCol,
		NavBorder: ns(r).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		ContentBorder: ns(r).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderCol).
			Background(panelBg).
			Padding(0, 1),
		NavTitle: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		ContentTitle: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		Selected: ns(r).
			Background(selectedBg).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),
		Normal: ns(r).
			Foreground(textCol).
			Background(panelBg).
			Padding(0, 1),
		BodyText: ns(r).
			Foreground(textCol).
			Background(panelBg),
		Footer: ns(r).
			Foreground(mutedCol),
		Counter: ns(r).
			Foreground(mutedCol),
		HelpKey: ns(r).
			Bold(true).
			Foreground(titleCol).
			Background(panelBg),
		HelpDesc: ns(r).
			Foreground(textCol).
			Background(panelBg),
	}
}

func focusedNavBorder(t Theme) lipgloss.Style {
	return t.NavBorder.BorderForeground(t.FocusColor)
}

func focusedContentBorder(t Theme) lipgloss.Style {
	return t.ContentBorder.BorderForeground(t.FocusColor)
}
