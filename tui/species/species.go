// Package species/species.go
package species

import (
	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	species := character.GetAllSpecies()
	items := make(menu.List, 0, len(species))

	for i, specInfo := range species {
		items = append(items, menu.Item{
			Header: specInfo.Name,
			Desc:   specInfo.Traits,
			Data:   character.Species(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Species",
		"Select Species",
		items,
		false,
		func(c *character.Character, val character.Species) tea.Cmd {
			c.Species = val
			c.Size = character.GetSpecies(val).Size[0]
			if character.GetSpecies(c.Species).Size.RequiresChoice() {
				return pager.SwitchTo(pager.ViewSpeciesSizeSelection)
			} else if c.Species == character.SpeciesDragonborn {
				return pager.SwitchTo(pager.ViewDraconicAncestrySelection)
			} else if c.Species == character.SpeciesElf {
				return pager.SwitchTo(pager.ViewElvenLineageSelection)
			}
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displaySpeciesInfo,
	)
}

func displaySpeciesInfo(c *character.Character, species character.Species, contentWidth int) string {
	spec := character.GetSpecies(species)

	// Header
	title := style.Title.Render(spec.Name)

	// Stat Grid / Key-Value Blocks
	creatureType := style.Label.Render("Creature Type: ") + spec.CreatureType
	size := style.Label.Render("Size: ") + spec.Size.GetAll()
	speed := style.Label.Render("Speed: ") + style.Badge.Render(spec.Speed.String())

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
		Render(spec.Traits)

	// Main flavor description
	desc := style.Desc.Render(spec.Desc)

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
