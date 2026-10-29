package tui

import (
	"fmt"
	"strings"

	"dndwiz/character"

	"github.com/charmbracelet/lipgloss"
)

const (
	colorSkill  = "6"
	colorTitle  = "86"
	colorLabel  = "205"
	colorDesc   = "246"
	colorHeader = "202"
)

// Text formatting styles
var (
	// Panel Header
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorTitle)).
			MarginBottom(1)

	// Subheaders
	subheaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorHeader)).
			Underline(true).
			MarginTop(1)

	// Section Labels (e.g., "Primary Abilities:")
	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorLabel))

	// Descriptions
	descStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color(colorDesc))

	subtitleStyle = descStyle.Inherit(descStyle).MarginBottom(1)

	// Sub-details / Badges
	badgeStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("238")).
			Foreground(lipgloss.Color("255")).
			Padding(0, 1)

	// Skills
	skillStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("238")).
			Foreground(lipgloss.Color(colorSkill)).
			Italic(true).
			Padding(0, 1)
	// Divider Line
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Margin(1, 0)

	// Tagline
	taglineStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("246"))

	// Tooltip with key binds
	tooltipStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Italic(true).
			Align(lipgloss.Center)

	// Popup
	popupStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(1, 2).
			Width(36).
			Align(lipgloss.Left)
)

// Menu Styles

func getSidebarMenuStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("62"))
}

func renderDivider(width int) string {
	if width <= 0 {
		width = 20
	}
	return dividerStyle.Render(strings.Repeat("─", width))
}

func renderAbilityScoresRow(c *character.Character, selected character.Ability, focused bool) string {
	selectedAbility := int(selected) - 1
	boxStyle := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	var boxes []string

	for i, score := range c.AbilityScore {
		// Determine style based on whether this box is selected
		currentBoxStyle := boxStyle
		if i == selectedAbility {
			currentBoxStyle = boxStyle.Inherit(labelStyle) // Or apply labelStyle directly to the content/box
		}

		scorestr := fmt.Sprintf("%d", score)
		if score == 0 {
			scorestr = "-"
		}

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			character.GetAbility(character.Ability(i+1)).Abbr,
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

func renderTitle(title string, subtitle string) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(title),
		descStyle.Render(subtitle),
	)
}

func renderAbilityRollsRow(scores *[6]int, selected int, focused bool) string {
	boxStyle := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	var boxes []string

	for i, score := range scores {
		// Determine style based on whether this box is selected
		currentBoxStyle := boxStyle
		if i == selected && focused && score != 0 {
			currentBoxStyle = boxStyle.Inherit(labelStyle) // Or apply labelStyle directly to the content/box
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
