package content

import (
	"dndwiz/character"
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
	Title:     "D&D Chararacter Sheet Wizard",
	MenuTitle: "Main Menu",
	Layout:    shell.LayoutCentered,
	Width:     30,

	GetItems: func(c *character.Character) shell.MenuItemList {
		return shell.MenuItemList{
			shell.MenuItem{Label: "New Character", Data: OptionNewGame},
			shell.MenuItem{Label: "Load Character", Data: OptionLoadGame},
			shell.MenuItem{Label: "Quit", Data: OptionQuit},
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
			return tea.Quit
		}
		return nil
	},
}
