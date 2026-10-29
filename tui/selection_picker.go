package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type selectionPicker[T any] struct {
	c                     *character.Character
	title                 string
	menu                  sidebarModel
	termWidth, termHeight int
	onSelect              func(*character.Character, T) tea.Cmd
	formatContent         func(*character.Character, T, int) string
}

func newSelectionPicker[T any](
	c *character.Character,
	title string,
	sidebarTitle string,
	items SidebarItemList,
	showDesc bool,
	onSelect func(*character.Character, T) tea.Cmd,
	formatContent func(*character.Character, T, int) string,
) selectionPicker[T] {
	sidebar := newSidebar(sidebarTitle, items, 32, showDesc)

	return selectionPicker[T]{
		c:             c,
		title:         title,
		menu:          sidebar,
		onSelect:      onSelect,
		formatContent: formatContent,
	}
}

func (m selectionPicker[T]) Init() tea.Cmd {
	return nil
}

func (m selectionPicker[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if val, ok := GetItemData[T](m.menu); ok {
				cmd = m.onSelect(m.c, val)
				if cmd == nil {
					cmd = switchToView(viewCharsheet)
				}
				return m, cmd
			}
		}
	}
	m.menu, cmd = m.menu.Update(msg, m.termHeight)
	return m, cmd
}

func (m selectionPicker[T]) View() string {
	if m.termWidth == 0 || m.termHeight <= 2 {
		return "Initializing..."
	}

	sidebarWidth := m.menu.width
	contentWidth := m.termWidth - sidebarWidth - 4

	var content string
	if val, ok := GetItemData[T](m.menu); ok {
		content = m.formatContent(m.c, val, contentWidth)
	} else {
		content = "No selection"
	}

	contentStyle := lipgloss.NewStyle().
		Width(contentWidth).
		Height(m.termHeight - 2).
		PaddingLeft(2)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.menu.View(),
		contentStyle.Render(content),
	)
}
