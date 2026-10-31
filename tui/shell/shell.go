// Package shell/shell.go
package shell

import (
	"dndwiz/character"
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
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
		content   viewport.Model
		textInput textinput.Model
		tX, tY    int
	}

	ShellConfig[T any] struct {
		Layout      ShellType
		Title       string
		MenuTitle   string
		MenuWidth   int // menu width (defaults to 32 if 0)
		MenuHeight  int // menu height
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
		ti.Width = cfg.MenuWidth
		if ti.Width <= 0 {
			ti.Width = 30
		}
		s.textInput = ti
	} else {
		items := cfg.GetItems(c)
		s.menu = NewMenu(cfg.MenuTitle, items, cfg.MenuWidth, cfg.MenuHeight, cfg.ShowDesc)
		s.content = viewport.New(10, 10) // minimum, will be resized by tea.Msg in update
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
		m.content.Update(msg)
	}

	return m, cmd
}

func (m Shell[T]) View() string {
	if m.cfg.Layout != LayoutInput {
		m.menu.setItems(m.cfg.GetItems(m.c))
	}

	var view string
	switch m.cfg.Layout {
	case LayoutInput:
		ui := lipgloss.JoinVertical(
			lipgloss.Left,
			m.cfg.Title,
			"",
			m.textInput.View(),
		)
		view = lipgloss.Place(m.tX, m.tY, lipgloss.Center, lipgloss.Center, ui)

	case LayoutSidebar:

		menuWidth := m.menu.GetWidth()
		if menuWidth <= 0 {
			menuWidth = minMenuWidth
		}

		_, contentFrameHeight := style.ContentBorder.GetFrameSize()
		contentWidth := max(m.tX-menuWidth, 10)
		contentHeight := m.tY - contentFrameHeight

		m.content.Height = contentHeight
		m.content.Width = contentWidth

		if val, ok := getItemData[T](m.menu); ok && m.cfg.Content != nil {
			rawContent := m.cfg.Content(m.c, val, contentWidth)
			m.content.SetContent(rawContent)
		}

		contentView := style.ContentBorder.Render(m.content.View())

		view = lipgloss.JoinHorizontal(lipgloss.Top, m.menu.View(), contentView)

	case LayoutCentered:
		view = lipgloss.Place(
			m.tX, m.tY,
			lipgloss.Center, lipgloss.Center,
			m.menu.View(),
		)

	default:
		view = m.menu.View()

	}
	return view
}
