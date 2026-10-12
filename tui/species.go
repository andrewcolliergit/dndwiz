package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewSpeciesModel(c *character.Character) tea.Model {
	species := character.GetAllSpecies()
	items := make(SidebarItemList, 0, len(species))

	for i, specInfo := range species {
		items = append(items, SidebarItem{
			title: specInfo.Name,
			desc:  specInfo.Traits,
			data:  character.Species(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Species",
		"Select Species",
		items,
		false,
		func(c *character.Character, val character.Species) tea.Cmd {
			c.Species = val
			if character.GetSpecies(c.Species).Size.RequiresChoice() {
				return switchToView(viewSpeciesSizeSelection)
			} else if c.Species == character.SpeciesDragonborn {
				return switchToView(viewDraconicAncestrySelection)
			} else if c.Species == character.SpeciesElf {
				return switchToView(viewElvenLineageSelection)
			}
			return switchToView(viewCharsheet)
		},
		displaySpeciesInfo,
	)
}

func displaySpeciesInfo(c *character.Character, species character.Species) string {
	spec := character.GetSpecies(species)

	// Header
	title := titleStyle.Render(spec.Name)

	// Stat Grid / Key-Value Blocks
	creatureType := labelStyle.Render("Creature Type: ") + spec.CreatureType
	size := labelStyle.Render("Size: ") + spec.Size.GetAll()
	speed := labelStyle.Render("Speed: ") + badgeStyle.Render(spec.Speed.String())

	// Combine stats into a neat horizontal/vertical block
	statsGroup := lipgloss.JoinVertical(
		lipgloss.Left,
		creatureType,
		size,
		speed,
	)

	// Special traits section
	traitsLabel := labelStyle.Render("Traits:")
	traits := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("248")).
		Render(spec.Traits)

	// Main flavor description
	desc := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(spec.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		statsGroup,
		renderDivider(40),
		traitsLabel,
		traits,
		renderDivider(40),
		desc,
	)
}
