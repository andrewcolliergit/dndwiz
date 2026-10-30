package classes

import (
	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewPrimaryAbilityModel(c *character.Character) tea.Model {
	abilities := character.GetClass(c.Class).PrimaryAbility.From
	items := make(menu.List, 0, len(abilities))

	for _, a := range abilities {
		items = append(items, menu.Item{
			Header: character.GetAbility(a).Name,
			Desc:   "",
			Data:   a,
		})
	}

	return menu.NewPicker(
		c,
		"Primary Ability",
		"Select Primary Ability",
		items,
		false,
		func(c *character.Character, selectedAbility character.Ability) tea.Cmd {
			// Update the character's primary ability selection
			c.PrimaryAbilities = character.Abilities{selectedAbility}
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displayPrimaryAbilityInfo,
	)
}

func displayPrimaryAbilityInfo(c *character.Character, a character.Ability, contentWidth int) string {
	abi := character.GetAbility(a)

	title := style.Title.Render(abi.Name)

	desc := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(abi.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		style.Divider(contentWidth-2),
		desc,
	)
}
