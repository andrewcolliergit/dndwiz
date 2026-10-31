package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Menu Styles

const (
	Pad int = 4
)

var (
	MenuBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("62"))

	ContentBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true, true, true, true)
)
