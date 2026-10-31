// Package style
package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Text formatting styles
var (
	// Panel Header
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		MarginBottom(1)

	// Subheaders
	Header = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("202")).
		Underline(true).
		MarginTop(1)

	// Section Labels (e.g., "Primary Abilities:")
	Label = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	// Descriptions
	Desc = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("246"))

	Subtitle = Desc.Inherit(Desc).MarginBottom(1)

	// Sub-details / Badges
	Badge = lipgloss.NewStyle().
		Background(lipgloss.Color("238")).
		Foreground(lipgloss.Color("255")).
		Padding(0, 1)

	// Skills
	Skill = lipgloss.NewStyle().
		Background(lipgloss.Color("238")).
		Foreground(lipgloss.Color("6")).
		Italic(true).
		Padding(0, 1)
	// divider Line
	divider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Margin(1, 0)

	// Tagline
	Tagline = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("246"))

	// Tooltip with key binds
	Tooltip = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		Align(lipgloss.Center)

	// Popup
	Popup = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Width(36).
		Align(lipgloss.Left)
)
