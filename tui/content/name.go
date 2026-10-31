// Package content/name
package content

import (
	"dndwiz/character"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"

	tea "github.com/charmbracelet/bubbletea"
)

var NameConfig = shell.ShellConfig[string]{
	Title:       "Character Name",
	Layout:      shell.LayoutInput,
	Placeholder: "Character name...",
	Width:       30,
	OnSubmit: func(c *character.Character, name string) tea.Cmd {
		if name != "" {
			c.Name = name
		}
		// Return to the character sheet after saving
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},
}
