package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewAbilityScoresMethodSelectionModel(c *character.Character) tea.Model {
	items := SidebarItemList{
		{
			title: "Point Buy",
			desc:  "",
			data:  viewAbilityScoresBuy,
		},
	}

	if c.Class != character.ClassNone {
		stdarray := character.GetClass(c.Class).StdArray

		items = append(items, SidebarItem{
			title: "Standard Array",
			desc:  stdarray.String(),
			data:  viewCharsheet,
		})
	}

	if !c.HasRolledAbilityScores {
		items = append(items, SidebarItem{
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

func displayAbilityScoreMethodInfo(c *character.Character, view currentView) string {
	var title, desc string

	switch view {
	case viewAbilityScoresRoll:
		title = "Roll"
		desc = "Roll description"
	case viewAbilityScoresBuy:
		title = "Point Buy"
		desc = "Point buy description"
	case viewCharsheet: // Standard Array
		tmpCharacter := character.NewCharacter()
		tmpCharacter.AbilityScore = character.GetClass(c.Class).StdArray
		box := renderAbilityScoresRow(&tmpCharacter, -1, false)
		title = "Standard Array"
		desc = box
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
