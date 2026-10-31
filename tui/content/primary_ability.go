package content

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

// PrimaryAbilityConfig is the declarative blueprint for selecting a class's primary ability.
var PrimaryAbilityConfig = shell.ShellConfig[ability.Ability]{
	Title:     "Primary Ability",
	MenuTitle: "Select Primary Ability",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	// 1. Populate valid primary abilities dynamically based on the character's class choice
	GetItems: func(c *character.Character) shell.MenuItemList {
		abilities := class.Get(c.Class).PrimaryAbility.From
		items := make(shell.MenuItemList, 0, len(abilities))

		for _, a := range abilities {
			items = append(items, shell.MenuItem{
				Header: ability.Get(a).Name,
				Desc:   "",
				Data:   a,
			})
		}
		return items
	},

	// 2. Handle selection and return to the character sheet
	OnSelect: func(c *character.Character, selectedAbility ability.Ability) tea.Cmd {
		c.PrimaryAbilities = ability.Abilities{selectedAbility}
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	// 3. Render the detail info for the highlighted ability
	Content: func(c *character.Character, a ability.Ability, contentWidth int) string {
		return displayPrimaryAbilityInfo(a, contentWidth)
	},
}

// --- Helper Formatting Functions ---

func displayPrimaryAbilityInfo(a ability.Ability, contentWidth int) string {
	abi := ability.Get(a)

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
