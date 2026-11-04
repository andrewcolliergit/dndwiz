package style

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Hint struct{ Key, Desc string }

var helpStyles = help.New().Styles

func TypesetHelp(hints ...Hint) string {
	sep := helpStyles.ShortSeparator.Inline(true).Render(" • ")

	parts := make([]string, 0, len(hints))
	for _, h := range hints {
		parts = append(parts,
			helpStyles.ShortKey.Inline(true).Render(h.Key)+" "+
				helpStyles.ShortDesc.Inline(true).Render(h.Desc),
		)
	}
	return strings.Join(parts, sep)
}

func TypesetTitle(title string, subtitle string, width int) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		Title.Render(title),
		TypesetDesc(subtitle, width),
	)
}

func TypesetBody(body string, width int) string {
	return Body.Width(width).Render(body)
}

func TypesetDesc(desc string, width int) string {
	return Desc.Width(width).Render(desc)
}

func TypesetBadges(badges []string, width int) string {
	var list []string

	for _, badge := range badges {
		list = append(list, Badge.Render(badge))
	}
	str := strings.Join(list, " ")
	return lipgloss.NewStyle().Width(width).Render(str)
}

func TypesetDivider(width int) string {
	if width <= 0 {
		width = 20
	}
	return divider.Render(strings.Repeat("─", width))
}

func TypesetTextBox(title, box string, hints ...Hint) string {
	const (
		width  = 50
		fieldW = 26 // must be >= the input's rendered width
	)

	center := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)
	field := lipgloss.NewStyle().Width(fieldW).Align(lipgloss.Left).Render(box)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		center.Render(list.DefaultStyles().Title.Render(title)),
		"",
		center.Render(field),
		"",
		center.Render(TypesetHelp(hints...)),
	)

	framed := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(width, 20, lipgloss.Center, lipgloss.Center, framed)
}

func TypesetScores(c *character.Character, selected ability.Ability, focused bool) string {
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

func TypesetRollsRow(scores *[6]int, selected int, focused bool) string {
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
