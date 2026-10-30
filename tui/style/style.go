// Package style
package style

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"

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

// Menu Styles

func GetSidebarMenuStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("62"))
}

func Divider(width int) string {
	if width <= 0 {
		width = 20
	}
	return divider.Render(strings.Repeat("─", width))
}

func RenderAbilityScoresRow(c *character.Character, selected ability.Ability, focused bool) string {
	selectedAbility := int(selected) - 1
	boxStyle := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	var boxes []string

	for i, score := range c.AbilityScore {
		// Determine style based on whether this box is selected
		currentBoxStyle := boxStyle
		if i == selectedAbility {
			currentBoxStyle = boxStyle.Inherit(Label) // Or apply labelStyle directly to the content/box
		}

		scorestr := fmt.Sprintf("%d", score)
		if score == 0 {
			scorestr = "-"
		}

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			ability.Get(ability.Ability(i+1)).Abbr,
			scorestr,
		)

		boxes = append(boxes, currentBoxStyle.Render(content))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)

	outerBorderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder())

	if focused {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("205")) // Pink/active color
	} else {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("240")) // Muted/inactive color
	}

	return outerBorderStyle.Render(row)
}

func RenderTitle(title string, subtitle string) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		Title.Render(title),
		Desc.Render(subtitle),
	)
}

func RenderAbilityRollsRow(scores *[6]int, selected int, focused bool) string {
	boxStyle := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	var boxes []string

	for i, score := range scores {
		// Determine style based on whether this box is selected
		currentBoxStyle := boxStyle
		if i == selected && focused && score != 0 {
			currentBoxStyle = boxStyle.Inherit(Label) // Or apply labelStyle directly to the content/box
		}

		scorestr := fmt.Sprintf("%d", score)

		if score == 0 {
			scorestr = "-"
		}

		num := i + 1
		suffix := "th"

		switch num {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		}

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			fmt.Sprintf("%d%s", num, suffix),
			scorestr,
		)

		boxes = append(boxes, currentBoxStyle.Render(content))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)

	outerBorderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder())

	if focused {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("205")) // Pink/active color
	} else {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("240")) // Muted/inactive color
	}

	return outerBorderStyle.Render(row)
}
