// Package shell/shell.go
package shell

import (
	"dndwiz/character"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ShellType int

const (
	LayoutCentered ShellType = iota
	LayoutSidebar
	LayoutInput
)

type (
	Shell[T any] struct {
		c         *character.Character
		cfg       ShellConfig[T]
		menu      Menu
		textInput textinput.Model
		tX, tY    int
	}

	ShellConfig[T any] struct {
		Title       string
		MenuTitle   string
		Layout      ShellType
		Width       int // menu width (defaults to 32 if 0)
		ShowDesc    bool
		Placeholder string
		GetItems    func(*character.Character) MenuItemList
		OnSelect    func(*character.Character, T) tea.Cmd
		OnSubmit    func(*character.Character, string) tea.Cmd
		Content     func(*character.Character, T, int) string // Optional for centerd layout
	}
)

func NewShell[T any](c *character.Character, cfg ShellConfig[T]) Shell[T] {
	s := Shell[T]{
		c:   c,
		cfg: cfg,
	}

	if cfg.Layout == LayoutInput {
		ti := textinput.New()
		ti.Placeholder = cfg.Placeholder
		ti.Focus()
		ti.Width = cfg.Width
		if ti.Width <= 0 {
			ti.Width = 30
		}
		s.textInput = ti
	} else {
		items := cfg.GetItems(c)
		menuWidth := cfg.Width
		if menuWidth <= 0 {
			menuWidth = 32
		}
		s.menu = NewMenu(cfg.MenuTitle, items, menuWidth, cfg.ShowDesc)
	}

	return s
}

func (m Shell[T]) Init() tea.Cmd {
	return nil
}

func (m Shell[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.tX = msg.Width
		m.tY = msg.Height

	case tea.KeyMsg:
		if m.cfg.Layout == LayoutInput {
			switch msg.String() {
			case "enter":
				if m.cfg.OnSubmit != nil {
					return m, m.cfg.OnSubmit(m.c, m.textInput.Value())
				}
				return m, nil
			}
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case "enter", " ":
			if val, ok := getItemData[T](m.menu); ok && m.cfg.OnSelect != nil {
				cmd := m.cfg.OnSelect(m.c, val)
				return m, cmd
			}
			return m, nil
		}
	}

	// Pass Key strokes to underlying menu model
	if m.cfg.Layout != LayoutInput {
		m.menu, cmd = m.menu.Update(msg, m.tY)
	}

	return m, cmd
}

func (m Shell[T]) View() string {
	if m.cfg.Layout != LayoutInput {
		m.menu.setItems(m.cfg.GetItems(m.c))
	}

	switch m.cfg.Layout {
	case LayoutInput:
		ui := lipgloss.JoinVertical(
			lipgloss.Left,
			m.cfg.Title,
			"",
			m.textInput.View(),
		)
		return lipgloss.Place(m.tX, m.tY, lipgloss.Center, lipgloss.Center, ui)

	case LayoutSidebar:
		menuView := m.menu.View()

		menuWidth := m.cfg.Width
		if menuWidth <= 0 {
			menuWidth = 32
		}

		contentWidth := max(m.tX-menuWidth-4, 10)

		var content string
		if val, ok := getItemData[T](m.menu); ok && m.cfg.Content != nil {
			content = m.cfg.Content(m.c, val, contentWidth)
		}

		return lipgloss.JoinHorizontal(lipgloss.Top, menuView, content)

	case LayoutCentered:
		return lipgloss.Place(
			m.tX, m.tY,
			lipgloss.Center, lipgloss.Center,
			m.menu.View(),
		)
	}

	return m.menu.View()
}
