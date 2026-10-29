package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewElvenLineageModel(c *character.Character) tea.Model {
	lineages := character.GetAllElvenLineages()
	items := make(SidebarItemList, 0, len(lineages))

	for i, l := range lineages {
		items = append(items, SidebarItem{
			title: l.Name,
			desc:  "",
			data:  character.ElvenLineage(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Elven Lineage",
		"Select Elven Lineage",
		items,
		false,
		func(c *character.Character, selectedElvenLineage character.ElvenLineage) tea.Cmd {
			// Update the character's primary ability selection
			c.ElvenLineage = selectedElvenLineage
			return switchToView(viewCharsheet)
		},
		displayElvenLineageInfo,
	)
}

func displayElvenLineageInfo(c *character.Character, e character.ElvenLineage, contentWidth int) string {
	el := character.GetElvenLineage(e)

	title := titleStyle.Render(el.Name)

	lvl1 := labelStyle.Render("Level 1: ") + el.Lvl1
	lvl3 := labelStyle.Render("Level 3: ") + el.Lvl3
	lvl5 := labelStyle.Render("Level 5: ") + el.Lvl5

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		renderDivider(contentWidth-2),
		lvl1,
		lvl3,
		lvl5,
	)
}
