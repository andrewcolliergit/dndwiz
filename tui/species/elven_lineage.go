package species

import (
	"dndwiz/character"
	elven "dndwiz/character/species/elven_lineage"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewElfModel(c *character.Character) tea.Model {
	lineages := elven.GetAll()
	items := make(menu.List, 0, len(lineages))

	for i, l := range lineages {
		items = append(items, menu.Item{
			Header: l.Name,
			Desc:   "",
			Data:   elven.Lineage(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Elven Lineage",
		"Select Elven Lineage",
		items,
		false,
		func(c *character.Character, selectedElvenLineage elven.Lineage) tea.Cmd {
			// Update the character's primary ability selection
			c.ElvenLineage = selectedElvenLineage
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displayElvenLineageInfo,
	)
}

func displayElvenLineageInfo(c *character.Character, e elven.Lineage, contentWidth int) string {
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
