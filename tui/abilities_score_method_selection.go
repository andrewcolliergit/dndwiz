package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewAbilityScoresMethodSelectionModel(c *character.Character) tea.Model {
	items := MenuItemList{
		{
			title: "Point Buy",
			desc:  "",
			data:  viewAbilityScoresBuy,
		},
	}

	if c.Class != character.ClassNone {
		stdarray := character.GetClass(c.Class).StdArray

		items = append(items, MenuItem{
			title: "Standard Array",
			desc:  stdarray.String(),
			data:  viewCharsheet,
		})
	}

	if !c.HasRolledAbilityScores {
		items = append(items, MenuItem{
			title: "Roll",
			desc:  "",
			data:  viewAbilityScoresRoll,
		})
	}

	return newSelectionPicker(
		c,
		"Ability Scores Selection Method",
		"Ability Score Allocation Method",
		items,
		false,
		func(c *character.Character, val currentView) tea.Cmd {
			switch val {
			case viewCharsheet:
				cls := character.GetClass(c.Class)
				c.AbilityScore = cls.StdArray
				c.AbilityScoreMethod = character.AbilityScoreMethodStdArray
			case viewAbilityScoresRoll:
				c.HasRolledAbilityScores = true
				c.AbilityScoreMethod = character.AbilityScoreMethodRoll
			case viewAbilityScoresBuy:
				c.AbilityScoreMethod = character.AbilityScoreMethodPointBuy
			}
			return switchToView(val)
		},
		displayAbilityScoreMethodInfo,
	)
}

func displayAbilityScoreMethodInfo(c *character.Character, view currentView, contentWidth int) string {
	var title, desc string

	switch view {
	case viewAbilityScoresRoll:
		title = "Roll"
		desc = "You may opt to roll for your skill points. Using this method, 6 scores will be rolled and you may assign them to the abilities of your choice.\nNote that once you have rolled, you will not have the option to roll again. You may, however, elect to Point Buy or take the standard array after rolling."
	case viewAbilityScoresBuy:
		title = "Point Buy"
		desc = "You may elect to customize your skill points as you see fit using Point Buy. With this method, you will be given a budget of 27 points to spend on skill point levels, where each level purchased becomes increasingly expensive."
	case viewCharsheet: // Standard Array
		tmpCharacter := character.NewCharacter()
		tmpCharacter.AbilityScore = character.GetClass(c.Class).StdArray
		box := renderAbilityScoresRow(&tmpCharacter, -1, false)
		title = "Standard Array"
		desc = "Rather than customizing your ability scores, you may elect to take the standard score array for your chosen class. The standard skill point array for the " + character.GetClass(c.Class).Name + " class is:\n\n" + box
	}

	// Header & Accents
	title = titleStyle.Render(title)

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
