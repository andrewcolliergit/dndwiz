package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewPrimaryAbilityModel(c *character.Character) tea.Model {
	abilities := character.GetClass(c.Class).PrimaryAbility.From
	items := make(SidebarItemList, 0, len(abilities))

	for _, a := range abilities {
		items = append(items, SidebarItem{
			title: character.GetAbility(a).Name,
			desc:  "",
			data:  a,
		})
	}

	return newSelectionPicker(
		c,
		"Primary Ability",
		"Select Primary Ability",
		items,
		false,
		func(c *character.Character, selectedAbility character.Ability) tea.Cmd {
			// Update the character's primary ability selection
			c.PrimaryAbilities = character.Abilities{selectedAbility}
			return switchToView(viewCharsheet)
		},
		displayPrimaryAbilityInfo,
	)
}

func displayPrimaryAbilityInfo(c *character.Character, a character.Ability, contentWidth int) string {
	abi := character.GetAbility(a)

	title := titleStyle.Render(abi.Name)

	desc := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(abi.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		renderDivider(contentWidth-2),
		desc,
	)
}
