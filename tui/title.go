package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type titleScreenModel struct {
	termWidth, termHeight int
	menu                  Menu
	title                 string
}

func (m titleScreenModel) Init() tea.Cmd {
	return nil
}

func makeTitleScreenMenuItems() MenuItemList {
	newCharacter := MenuItem{
		title: "New Character",
		data:  viewCharsheet,
	}

	loadCharacter := MenuItem{
		title: "Load Character",
		data:  viewCharsheet,
	}

	quit := MenuItem{
		title: "Quit",
		data:  viewQuit,
	}

	var items MenuItemList
	return append(items, newCharacter, loadCharacter, quit)
}

func NewTitleScreenModel() titleScreenModel {
	return titleScreenModel{
		menu: newMenu(
			menuStyleSimple,
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
			if v, ok := GetItemData[currentView](m.menu); ok {
				if v == viewQuit {
					return m, tea.Quit
				}
				return m, switchToView(v)
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
