// Package species/species.go
package species

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	items := make(menu.List, 0, len(species.GetAll()))

	for i, specInfo := range species.GetAll() {
		items = append(items, menu.Item{
			Header: specInfo.Name,
			Desc:   specInfo.Traits,
			Data:   species.Species(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Species",
		"Select Species",
		items,
		false,
		func(c *character.Character, val species.Species) tea.Cmd {
			c.Species = val
			c.Size = species.Get(val).Size[0]
			if species.Get(c.Species).Size.RequiresChoice() {
				return pager.SwitchTo(pager.SpeciesSizeSelection)
			} else if c.Species == species.Dragonborn {
				return pager.SwitchTo(pager.SpeciesAncestrySelection)
			} else if c.Species == species.Elf {
				return pager.SwitchTo(pager.SpeciesLineageSelection)
			}
			return pager.SwitchTo(pager.Charsheet)
		},
		displaySpeciesInfo,
	)
}

func displaySpeciesInfo(c *character.Character, s species.Species, contentWidth int) string {
	info := species.Get(s)

	// Header
	title := style.Title.Render(info.Name)

	// Stat Grid / Key-Value Blocks
	creatureType := style.Label.Render("Creature Type: ") + info.CreatureType
	size := style.Label.Render("Size: ") + info.Size.GetAll()
	speed := style.Label.Render("Speed: ") + style.Badge.Render(info.Speed.String())

	// Combine stats into a neat horizontal/vertical block
	statsGroup := lipgloss.JoinVertical(
		lipgloss.Left,
		creatureType,
		size,
		speed,
	)

	// Special traits section
	traits := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("248")).
		Render(info.Traits)

	// Main flavor description
	desc := style.Desc.Render(info.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		desc,
		style.Divider(contentWidth-2),
		statsGroup,
		style.Header.Render("Traits"),
		traits,
	)
}
