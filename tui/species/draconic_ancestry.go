package species

import (
	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewDragonModel(c *character.Character) tea.Model {
	ancestries := character.GetAllDraconicAncestries()
	items := make(menu.List, 0, len(ancestries))

	for i, a := range ancestries {
		items = append(items, menu.Item{
			Header: a.Name,
			Desc:   a.DamageType,
			Data:   character.DraconicAncestry(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Draconic Ancestry",
		"Select Draconic Ancestry",
		items,
		true,
		func(c *character.Character, selectedDraconicAncestry character.DraconicAncestry) tea.Cmd {
			// Update the character's primary ability selection
			c.DraconicAncestry = selectedDraconicAncestry
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displayDraconicAncestryInfo,
	)
}

func displayDraconicAncestryInfo(c *character.Character, d character.DraconicAncestry, contentWidth int) string {
	dai := character.GetDraconicAncestry(d)

	title := style.Title.Render(dai.Name)
	dmgType := style.Label.Render("Damage Type: ") + style.Badge.Render(dai.DamageType)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		style.Divider(contentWidth-2),
		dmgType,
	)
}
