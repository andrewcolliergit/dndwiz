package style

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/score"

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

func TypesetHelpBar(width int, hints []Hint) string {
	return lipgloss.NewStyle().
		Width(width).
		MaxWidth(width).
		Align(lipgloss.Center).
		Render(TypesetHelp(hints...))
}

func TypesetTitle(title string, subtitle string, width int) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		Title.Render(title),
		TypesetDesc(subtitle, width, lipgloss.Left),
	)
}

func TypesetBody(body string, width int) string {
	return Body.Width(width).Render(body)
}

func TypesetDesc(desc string, width int, pos lipgloss.Position) string {
	return Desc.Width(width).Align(pos).Render(desc)
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

func TypesetTextBox(title, box string) string {
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
	)

	framed := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(width, 20, lipgloss.Center, lipgloss.Center, framed)
}

// TypesetScoresPicked is TypesetScores plus a "picked" box, drawn as a filled
// block so it reads differently from the hover highlight.
func TypesetScoresPicked(c *character.Character, selected, picked ability.Ability, focused bool) string {
	selectedAbility := int(selected) - 1
	pickedAbility := int(picked) - 1
	boxStyle := lipgloss.NewStyle().
		Width(5).
		Align(lipgloss.Center)

	pickedStyle := boxStyle.
		Background(lipgloss.Color("205")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	var boxes []string

	for i, sc := range c.AbilityScore {
		currentBoxStyle := boxStyle
		if i == selectedAbility {
			currentBoxStyle = boxStyle.Inherit(Label)
		}
		if i == pickedAbility {
			currentBoxStyle = pickedStyle
		}

		scorestr := fmt.Sprintf("%d", sc)
		if sc == 0 {
			scorestr = "-"
		}

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			ability.Get(ability.Ability(i+1)).Abbr,
			scorestr,
			score.ModifierString(sc), // "" when unassigned, so box height stays constant
		)

		boxes = append(boxes, currentBoxStyle.Render(content))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)

	outerBorderStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	if focused {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("205"))
	} else {
		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("240"))
	}

	return outerBorderStyle.Render(row)
}

func TypesetScores(c *character.Character, selected ability.Ability, focused bool) string {
	return TypesetScoresPicked(c, selected, ability.None, focused)
}

// func TypesetScores(c *character.Character, selected ability.Ability, focused bool) string {
// 	selectedAbility := int(selected) - 1
// 	boxStyle := lipgloss.NewStyle().
// 		Width(5).
// 		Align(lipgloss.Center)
//
// 	var boxes []string
//
// 	for i, score := range c.AbilityScore {
// 		// Determine style based on whether this box is selected
// 		currentBoxStyle := boxStyle
// 		if i == selectedAbility {
// 			currentBoxStyle = boxStyle.Inherit(Label) // Or apply labelStyle directly to the content/box
// 		}
//
// 		scorestr := fmt.Sprintf("%d", score)
// 		if score == 0 {
// 			scorestr = "-"
// 		}
//
// 		content := lipgloss.JoinVertical(
// 			lipgloss.Left,
// 			ability.Get(ability.Ability(i+1)).Abbr,
// 			scorestr,
// 		)
//
// 		boxes = append(boxes, currentBoxStyle.Render(content))
// 	}
//
// 	row := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)
//
// 	outerBorderStyle := lipgloss.NewStyle().
// 		Border(lipgloss.RoundedBorder())
//
// 	if focused {
// 		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("205")) // Pink/active color
// 	} else {
// 		outerBorderStyle = outerBorderStyle.BorderForeground(lipgloss.Color("240")) // Muted/inactive color
// 	}
//
// 	return outerBorderStyle.Render(row)
// }

// TypesetDice renders the four-dice box. active is the spinning die (-1 for none),
// dropped is the die excluded from the sum (-1 for none).
func TypesetDice(dice [4]int, active, dropped int, focused bool) string {
	boxStyle := lipgloss.NewStyle().Width(5).Align(lipgloss.Center)

	var boxes []string
	for i, v := range dice {
		val := fmt.Sprintf("%d", v)
		if v == 0 {
			val = "-"
		}

		label := fmt.Sprintf("d%d", i+1)
		s := boxStyle
		switch i {
		case dropped:
			val = Desc.Strikethrough(true).Render(val)
		case active:
			s = boxStyle.Inherit(Label)
		}

		boxes = append(boxes, s.Render(lipgloss.JoinVertical(lipgloss.Left, label, val)))
	}

	border := lipgloss.Color("240")
	if focused {
		border = lipgloss.Color("205")
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(border).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, boxes...))
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
		case 3:
			suffix = "rd"
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
