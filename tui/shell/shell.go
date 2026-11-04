// Package shell/shell.go
package shell

import (
	"dndwiz/character"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ShellType int

const (
	LayoutCentered ShellType = iota
	LayoutSidebar
)

type (
	Shell[T any] struct {
		c       *character.Character
		cfg     ShellConfig[T]
		menu    list.Model
		content viewport.Model
		tX, tY  int
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
		GetItems        func(*character.Character) []MenuItem
		OnSelect        func(*character.Character, T) tea.Cmd
		OnSubmit        func(*character.Character, string) tea.Cmd
		Content         func(*character.Character, T, int) string // Optional for centerd layout
	}
)

func NewShell[T any](c *character.Character, cfg ShellConfig[T]) Shell[T] {
	return Shell[T]{
		c:   c,
		cfg: cfg,
		menu: newList(
			cfg.MenuTitle,
			cfg.GetItems(c),
			cfg.ShowDesc,
		),
		content: viewport.New(10, 10), // minimum, will be resized by tea.Msg in update
	}
}

func (m Shell[T]) Init() tea.Cmd {
	return nil
}

func (m *Shell[T]) resizeList() {
	_, frameH := style.MenuBorder.GetFrameSize()

	w := m.cfg.MenuWidth
	if w == 0 {
		w = minMenuWidth
	}
	w = min(w, maxMenuWidth)

	h := m.cfg.MenuHeight
	if h == 0 {
		h = m.tY - frameH
	}
	m.menu.SetSize(w, max(h, minMenuHeight))
}

func (m *Shell[T]) syncContent() {
	if m.cfg.Layout != LayoutSidebar {
		return
	}

	menuWidth := lipgloss.Width(style.MenuBorder.Render(m.menu.View()))
	frameW, frameH := style.ContentBorder.GetFrameSize()
	contentWidth := max(m.tX-menuWidth-frameW, 10)

	m.content.Width = contentWidth
	m.content.Height = max(m.tY-frameH, 1)

	if val, ok := selectedData[T](m.menu); ok && m.cfg.Content != nil {
		raw := m.cfg.Content(m.c, val, contentWidth)
		m.content.SetContent(lipgloss.NewStyle().Width(contentWidth).Render(raw))
	}
}

func (m Shell[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case pager.RefreshMsg:
		m.refresh()
		return m, nil
	case tea.WindowSizeMsg:
		m.tX, m.tY = msg.Width, msg.Height
		m.resizeList()

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if val, ok := selectedData[T](m.menu); ok && m.cfg.OnSelect != nil {
				return m, m.cfg.OnSelect(m.c, val)
			}
			return m, nil
		case "shift+up", "K":
			m.content.ScrollUp(1)
			return m, nil
		case "shift+down", "J":
			m.content.ScrollDown(1)
			return m, nil
		}
	}

	prev := m.menu.Index()
	m.menu, cmd = m.menu.Update(msg)
	if m.menu.Index() != prev {
		m.content.GotoTop()
	}
	m.refresh()

	return m, cmd
}

func (m Shell[T]) View() string {
	switch m.cfg.Layout {
	case LayoutSidebar:
		return lipgloss.JoinHorizontal(lipgloss.Top,
			style.MenuBorder.Render(m.menu.View()),
			style.ContentBorder.Render(m.content.View()),
		)
	case LayoutCentered:
		return lipgloss.Place(m.tX, m.tY, lipgloss.Center, lipgloss.Center,
			style.MenuBorder.Render(m.menu.View()))
	}
	return ""
}

func (m *Shell[T]) refresh() {
	if m.cfg.Layout != LayoutSidebar && m.cfg.Layout != LayoutCentered {
		return
	}
	m.menu.SetItems(toListItems(m.cfg.GetItems(m.c)))
	m.syncContent()
}
