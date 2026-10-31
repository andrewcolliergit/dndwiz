package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Menu Styles

func SidebarMenu(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("62"))
}
