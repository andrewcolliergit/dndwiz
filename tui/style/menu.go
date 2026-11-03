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
			Border(lipgloss.RoundedBorder(), true, true, true, true).
			Padding(0, 2, 0, 2)
)

func ModalStyle(tX int, tY int, content string) string {
	return lipgloss.Place(tX, tY, lipgloss.Center, lipgloss.Center, content)
}
