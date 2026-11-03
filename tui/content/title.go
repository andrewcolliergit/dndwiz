package content

import (
	"dndwiz/character"
	dialog "dndwiz/tui/dialogue"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"

	tea "github.com/charmbracelet/bubbletea"
)

type TitleOption string

const (
	OptionNewGame  TitleOption = "new_game"
	OptionLoadGame TitleOption = "load_game"
	OptionQuit     TitleOption = "quit_game"
)

var TitleConfig = shell.ShellConfig[TitleOption]{
	ContentTitle: "D&D Chararacter Sheet Wizard",
	MenuTitle:    "Main Menu",
	MenuHeight:   9,
	Layout:       shell.LayoutCentered,

	GetItems: func(c *character.Character) []shell.MenuItem {
		return []shell.MenuItem{
			{Label: "New Character", Data: OptionNewGame},
			{Label: "Load Character", Data: OptionLoadGame},
			{Label: "Quit", Data: OptionQuit},
		}
	},

	OnSelect: func(c *character.Character, selected TitleOption) tea.Cmd {
		switch selected {
		case OptionNewGame:
			return pager.SwitchTo(pager.Charsheet)
		case OptionLoadGame:
			// Not yet implemented
			return nil
		case OptionQuit:
			return dialog.Confirm(dialog.Options{
				Title:     "Quit?",
				Body:      "Are you sure you would like to quit the program?",
				OnConfirm: tea.Quit,
			})
		}
		return nil
	},
}
