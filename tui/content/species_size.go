package content

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
)

var SpeciesSizeConfig = shell.ShellConfig[string]{
	Title:     "Species Size",
	MenuTitle: "Select Species Size",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	GetItems: func(c *character.Character) shell.MenuItemList {
		sizes := species.Get(c.Species).Size
		items := make(shell.MenuItemList, 0, len(sizes))

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
