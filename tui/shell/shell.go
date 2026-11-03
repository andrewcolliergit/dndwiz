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
		Layout          ShellType
		ContentTitle    string
		ContentSubtitle string
		MenuTitle       string
		MenuWidth       int // menu width (defaults to 32 if 0)
		MenuHeight      int // menu height
		ShowDesc        bool
		Placeholder     string
		GetItems        func(*character.Character) MenuItemList
		OnSelect        func(*character.Character, T) tea.Cmd
		OnSubmit        func(*character.Character, string) tea.Cmd
		Content         func(*character.Character, T, int) string // Optional for centerd layout
	}
)

func NewShell[T any](c *character.Character, cfg ShellConfig[T]) Shell[T] {
	s := Shell[T]{
		c:   c,
		cfg: cfg,
	}

	switch cfg.Layout {
	case LayoutInput:
		ti := textinput.New()
		ti.Placeholder = cfg.Placeholder
		ti.Focus()
		ti.Width = cfg.MenuWidth
		if ti.Width <= 0 {
			ti.Width = 30
		}
		s.textInput = ti

	default:
		s.menu = NewMenu(
			cfg.MenuTitle,
			cfg.GetItems(c),
			cfg.MenuWidth,
			cfg.MenuHeight,
			cfg.ShowDesc,
		)
		s.content = viewport.New(10, 10) // minimum, will be resized by tea.Msg in update

	}

	return s
}

func (m Shell[T]) Init() tea.Cmd {
	return nil
}

func (m *Shell[T]) syncContent() {
	if m.cfg.Layout != LayoutSidebar {
		return
	}

	menuWidth := lipgloss.Width(m.menu.View())
	frameW, frameH := style.ContentBorder.GetFrameSize()
	contentWidth := max(m.tX-menuWidth-frameW, 10)

	m.content.Width = contentWidth
	m.content.Height = max(m.tY-frameH, 1)

	if val, ok := getItemData[T](m.menu); ok && m.cfg.Content != nil {
		raw := m.cfg.Content(m.c, val, contentWidth)
		m.content.SetContent(lipgloss.NewStyle().Width(contentWidth).Render(raw))
	}
}

func (m Shell[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.tX = msg.Width
		m.tY = msg.Height

	case tea.KeyMsg:
		// Text input layout: handle and return early
		if m.cfg.Layout == LayoutInput {
			if msg.String() == "enter" {
				if m.cfg.OnSubmit != nil {
					return m, m.cfg.OnSubmit(m.c, m.textInput.Value())
				}
				return m, nil
			}
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		// Menu layouts
		switch msg.String() {
		case "enter", " ":
			if val, ok := getItemData[T](m.menu); ok && m.cfg.OnSelect != nil {
				return m, m.cfg.OnSelect(m.c, val)
			}
			return m, nil
		case "shift+up", "K":
			m.syncContent()
			m.content.ScrollUp(1)
			return m, nil
		case "shift+down", "J":
			m.syncContent()
			m.content.ScrollDown(1)
			return m, nil
		}
	}

	// Pass keystrokes to the underlying menu model
	if m.cfg.Layout != LayoutInput {
		prev := m.menu.List.Index()
		m.menu, cmd = m.menu.Update(msg, m.tY)
		if m.menu.List.Index() != prev {
			m.content.GotoTop()
		}
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
			m.cfg.ContentTitle,
			"",
			m.textInput.View(),
		)
		view = lipgloss.Place(m.tX, m.tY, lipgloss.Center, lipgloss.Center, ui)

	case LayoutSidebar:
		m.syncContent()
		menuView := m.menu.View()
		contentView := style.ContentBorder.Render(m.content.View())
		view = lipgloss.JoinHorizontal(lipgloss.Top, menuView, contentView)

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
