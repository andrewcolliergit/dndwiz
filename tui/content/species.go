package content

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var SpeciesSelectionConfig = shell.ShellConfig[species.Species]{
	Title:     "Species",
	MenuTitle: "Select Species",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	GetItems: func(c *character.Character) shell.MenuItemList {
		allSpecies := species.GetAll()
		items := make(shell.MenuItemList, 0, len(allSpecies))

		for i, specInfo := range allSpecies {
			items = append(items, shell.MenuItem{
				Label: specInfo.Name,
				Desc:  specInfo.Traits,
				Data:  species.Species(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selected species.Species) tea.Cmd {
		c.Species = selected
		spec := species.Get(selected)
		c.Size = spec.Size[0]

		if spec.Size.RequiresChoice() {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesSizeSelection)
			}
		} else if c.Species == species.Dragonborn {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesAncestrySelection)
			}
		} else if c.Species == species.Elf {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesLineageSelection)
			}
		}

		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, selected species.Species, contentWidth int) string {
		info := species.Get(selected)

		// Header
		title := style.Title.Render(info.Name)

		// Stat Grid / Key-Value Blocks
		creatureType := style.Label.Render("Creature Type: ") + info.CreatureType
		size := style.Label.Render("Size: ") + info.Size.GetAll()
		speed := style.Label.Render("Speed: ") + style.Badge.Render(info.Speed.String())

		// Combine stats into a neat vertical block
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
	},
}
