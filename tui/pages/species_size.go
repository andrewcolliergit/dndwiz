// Package species_size.go
package pages

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
)

var SpeciesSizeConfig = shell.ShellConfig[string]{
	ContentTitle: "Species Size",
	MenuTitle:    "Select Species Size",
	Layout:       shell.LayoutSidebar,
	MenuWidth:    32,
	ShowDesc:     false,

	GetItems: func(c *character.Character) []shell.MenuItem {
		sizes := species.Get(c.Species).Size
		items := make([]shell.MenuItem, 0, len(sizes))

		for _, s := range sizes {
			items = append(items, shell.MenuItem{
				Label: s,
				Desc:  "",
				Data:  s,
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedSize string) tea.Cmd {
		c.Size = selectedSize
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, selectedSize string, contentWidth int) string {
		return style.Title.Render(selectedSize)
	},
}
