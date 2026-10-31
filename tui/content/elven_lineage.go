package content

import (
	"dndwiz/character"
	elven "dndwiz/character/species/elven_lineage"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ElvenLineageConfig is the declarative blueprint for selecting an elven lineage.
var ElvenLineageConfig = shell.ShellConfig[elven.Lineage]{
	Title:     "Elven Lineage",
	MenuTitle: "Select Elven Lineage",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	// 1. Populate elven lineages dynamically
	GetItems: func(c *character.Character) shell.MenuItemList {
		lineages := elven.GetAll()
		items := make(shell.MenuItemList, 0, len(lineages))

		for i, l := range lineages {
			items = append(items, shell.MenuItem{
				Header: l.Name,
				Desc:   "",
				Data:   elven.Lineage(i + 1),
			})
		}
		return items
	},

	// 2. Handle selection and navigation back to character sheet
	OnSelect: func(c *character.Character, selectedElvenLineage elven.Lineage) tea.Cmd {
		c.ElvenLineage = selectedElvenLineage
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	// 3. Render lineage detail pane
	Content: func(c *character.Character, e elven.Lineage, contentWidth int) string {
		return displayElvenLineageInfo(e, contentWidth)
	},
}

// --- Helper Formatting Functions ---

func displayElvenLineageInfo(e elven.Lineage, contentWidth int) string {
	el := elven.Get(e)

	title := style.Title.Render(el.Name)

	lvl1 := style.Label.Render("Level 1: ") + el.Lvl1
	lvl3 := style.Label.Render("Level 3: ") + el.Lvl3
	lvl5 := style.Label.Render("Level 5: ") + el.Lvl5

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		style.Divider(contentWidth-2),
		lvl1,
		lvl3,
		lvl5,
	)
}
