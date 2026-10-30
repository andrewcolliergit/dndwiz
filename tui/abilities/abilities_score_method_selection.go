// Package abilities
package abilities

import (
	"dndwiz/character"
	"dndwiz/character/class"
	"dndwiz/character/score"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewOptionModel(c *character.Character) tea.Model {
	items := menu.List{
		{
			Header: "Point Buy",
			Desc:   "",
			Data:   pager.AbilityScoresBuy,
		},
	}

	if c.Class != class.None {
		stdarray := class.Get(c.Class).StdArray

		items = append(items, menu.Item{
			Header: "Standard Array",
			Desc:   stdarray.String(),
			Data:   pager.Charsheet,
		})
	}

	if !c.HasRolledAbilityScores {
		items = append(items, menu.Item{
			Header: "Roll",
			Desc:   "",
			Data:   pager.AbilityScoresRoll,
		})
	}

	return menu.NewPicker(
		c,
		"Ability Scores Selection Method",
		"Ability Score Allocation Method",
		items,
		false,
		func(c *character.Character, val pager.CurrentView) tea.Cmd {
			switch val {
			case pager.Charsheet:
				cls := class.Get(c.Class)
				c.AbilityScore = cls.StdArray
				c.AbilityScoreMethod = score.StdArray
			case pager.AbilityScoresRoll:
				c.HasRolledAbilityScores = true
				c.AbilityScoreMethod = score.Roll
			case pager.AbilityScoresBuy:
				c.AbilityScoreMethod = score.PointBuy
			}
			return pager.SwitchTo(val)
		},
		displayAbilityScoreMethodInfo,
	)
}

func displayAbilityScoreMethodInfo(c *character.Character, view pager.CurrentView, contentWidth int) string {
	var title, desc string

	switch view {
	case pager.AbilityScoresRoll:
		title = "Roll"
		desc = "You may opt to roll for your skill points. Using this method, 6 scores will be rolled and you may assign them to the abilities of your choice.\nNote that once you have rolled, you will not have the option to roll again. You may, however, elect to Point Buy or take the standard array after rolling."
	case pager.AbilityScoresBuy:
		title = "Point Buy"
		desc = "You may elect to customize your skill points as you see fit using Point Buy. With this method, you will be given a budget of 27 points to spend on skill point levels, where each level purchased becomes increasingly expensive."
	case pager.Charsheet: // Standard Array
		tmpCharacter := character.NewCharacter()
		tmpCharacter.AbilityScore = class.Get(c.Class).StdArray
		box := style.RenderAbilityScoresRow(&tmpCharacter, -1, false)
		title = "Standard Array"
		desc = "Rather than customizing your ability scores, you may elect to take the standard score array for your chosen class. The standard skill point array for the " + class.Get(c.Class).Name + " class is:\n\n" + box
	}

	// Header & Accents
	title = style.Title.Render(title)

	// Body content
	desc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		desc,
	)
}
