// Package title
package title

import (
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type titleScreenModel struct {
	termWidth, termHeight int
	menu                  menu.Menu
	title                 string
}

func (m titleScreenModel) Init() tea.Cmd {
	return nil
}

func makeTitleScreenMenuItems() menu.List {
	newCharacter := menu.Item{
		Header: "New Character",
		Data:   pager.Charsheet,
	}

	loadCharacter := menu.Item{
		Header: "Load Character",
		Data:   pager.Charsheet,
	}

	quit := menu.Item{
		Header: "Quit",
		Data:   pager.Quit,
	}

	var items menu.List
	return append(items, newCharacter, loadCharacter, quit)
}

func NewTitleScreenModel() titleScreenModel {
	return titleScreenModel{
		menu: menu.New(
			"D&D Wizard",
			makeTitleScreenMenuItems(),
			18,
			false,
		),
	}
}

func (m titleScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if v, ok := menu.GetItemData[pager.CurrentView](m.menu); ok {
				if v == pager.Quit {
					return m, tea.Quit
				}
				return m, pager.SwitchTo(v)
			}
		}
	}
	m.menu, cmd = m.menu.Update(msg, 15)
	return m, cmd
}

func (m titleScreenModel) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		m.title, m.menu.View(),
	)

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
