// Package pages/primary_ability.go
package pages

import (
	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/class"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var PrimaryAbilityConfig = shell.ShellConfig[ability.Ability]{
	ContentTitle: "Primary Ability",
	MenuTitle:    "Select Primary Ability",
	Layout:       shell.LayoutSidebar,
	MenuWidth:    32,
	ShowDesc:     false,

	GetItems: func(c *character.Character) []shell.MenuItem {
		abilities := class.Get(c.Class).PrimaryAbility.From
		items := make([]shell.MenuItem, 0, len(abilities))

		for _, a := range abilities {
			items = append(items, shell.MenuItem{
				Label: ability.Get(a).Name,
				Desc:  "",
				Data:  a,
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedAbility ability.Ability) tea.Cmd {
		c.PrimaryAbilities = ability.Abilities{selectedAbility}
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, a ability.Ability, contentWidth int) string {
		abi := ability.Get(a)

		title := style.Title.Render(abi.Name)

		desc := lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Render(abi.Desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			style.TypesetDivider(contentWidth-2),
			desc,
		)
	},
}
