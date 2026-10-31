package content

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
)

// SpeciesSizeConfig is the declarative blueprint for selecting a species size.
var SpeciesSizeConfig = shell.ShellConfig[string]{
	Title:     "Species Size",
	MenuTitle: "Select Species Size",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	// 1. Populate available sizes based on the character's species
	GetItems: func(c *character.Character) shell.MenuItemList {
		sizes := species.Get(c.Species).Size
		items := make(shell.MenuItemList, 0, len(sizes))

		for _, s := range sizes {
			items = append(items, shell.MenuItem{
				Header: s,
				Desc:   "",
				Data:   s,
			})
		}
		return items
	},

	// 2. Handle selection and navigation back to character sheet
	OnSelect: func(c *character.Character, selectedSize string) tea.Cmd {
		c.Size = selectedSize
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	// 3. Render species size detail pane
	Content: func(c *character.Character, selectedSize string, contentWidth int) string {
		return displaySpeciesSizeInfo(selectedSize)
	},
}

// --- Helper Formatting Functions ---

func displaySpeciesSizeInfo(s string) string {
	return style.Title.Render(s)
}
