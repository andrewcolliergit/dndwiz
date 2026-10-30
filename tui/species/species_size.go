package species

import (
	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
)

func NewSizeModel(c *character.Character) tea.Model {
	sizes := character.GetSpecies(c.Species).Size
	items := make(menu.List, 0, len(sizes))

	for _, s := range sizes {
		items = append(items, menu.Item{
			Header: s,
			Desc:   "",
			Data:   s,
		})
	}

	return menu.NewPicker(
		c,
		"Species Size",
		"Select Species Size",
		items,
		false,
		func(c *character.Character, selectedSize string) tea.Cmd {
			// Update the character's primary ability selection
			c.Size = selectedSize
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displaySpeciesSizeInfo,
	)
}

func displaySpeciesSizeInfo(c *character.Character, s string, contentWidth int) string {
	return style.Title.Render(s)
}
