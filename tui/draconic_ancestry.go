package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewDraconicAncestryModel(c *character.Character) tea.Model {
	ancestries := character.GetAllDraconicAncestries()
	items := make(MenuItemList, 0, len(ancestries))

	for i, a := range ancestries {
		items = append(items, MenuItem{
			title: a.Name,
			desc:  a.DamageType,
			data:  character.DraconicAncestry(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Draconic Ancestry",
		"Select Draconic Ancestry",
		items,
		true,
		func(c *character.Character, selectedDraconicAncestry character.DraconicAncestry) tea.Cmd {
			// Update the character's primary ability selection
			c.DraconicAncestry = selectedDraconicAncestry
			return switchToView(viewCharsheet)
		},
		displayDraconicAncestryInfo,
	)
}

func displayDraconicAncestryInfo(c *character.Character, d character.DraconicAncestry, contentWidth int) string {
	dai := character.GetDraconicAncestry(d)

	title := titleStyle.Render(dai.Name)
	dmgType := labelStyle.Render("Damage Type: ") + badgeStyle.Render(dai.DamageType)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		renderDivider(contentWidth-2),
		dmgType,
	)
}
