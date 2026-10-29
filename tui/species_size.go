package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
)

func NewSpeciesSizeModel(c *character.Character) tea.Model {
	sizes := character.GetSpecies(c.Species).Size
	items := make(MenuItemList, 0, len(sizes))

	for _, s := range sizes {
		items = append(items, MenuItem{
			title: s,
			desc:  "",
			data:  s,
		})
	}

	return newSelectionPicker(
		c,
		"Species Size",
		"Select Species Size",
		items,
		false,
		func(c *character.Character, selectedSize string) tea.Cmd {
			// Update the character's primary ability selection
			c.Size = selectedSize
			return switchToView(viewCharsheet)
		},
		displaySpeciesSizeInfo,
	)
}

func displaySpeciesSizeInfo(c *character.Character, s string, contentWidth int) string {
	return titleStyle.Render(s)
}
