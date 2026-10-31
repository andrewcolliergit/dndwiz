package content

import (
	"dndwiz/character"
	draconic "dndwiz/character/species/draconic_ancestry"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var DraconicAncestryConfig = shell.ShellConfig[draconic.Ancestry]{
	Title:     "Draconic Ancestry",
	MenuTitle: "Select Draconic Ancestry",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  true,

	GetItems: func(c *character.Character) shell.MenuItemList {
		ancestries := draconic.GetAll()
		items := make(shell.MenuItemList, 0, len(ancestries))

		for i, a := range ancestries {
			items = append(items, shell.MenuItem{
				Label: a.Name,
				Desc:  a.DamageType,
				Data:  draconic.Ancestry(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedDraconicAncestry draconic.Ancestry) tea.Cmd {
		c.DraconicAncestry = selectedDraconicAncestry
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, d draconic.Ancestry, contentWidth int) string {
		dai := draconic.Get(d)

		title := style.Title.Render(dai.Name)
		dmgType := style.Label.Render("Damage Type: ") + style.Badge.Render(dai.DamageType)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			style.Divider(contentWidth-2),
			dmgType,
		)
	},
}
