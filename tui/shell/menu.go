// Package shell/menu.go
package shell

import (
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	minMenuHeight int = 4
	minMenuWidth  int = 32
	maxMenuWidth  int = 64
)

type MenuItem struct {
	Label string
	Desc  string
	Data  any
}

func (i MenuItem) Title() string       { return i.Label }
func (i MenuItem) Description() string { return i.Desc }
func (i MenuItem) FilterValue() string { return i.Label }

type Menu struct {
	List   list.Model
	Width  int
	Height int
}

func newList(title string, items []MenuItem, showDesc bool) list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = showDesc

	l := list.New(toListItems(items), delegate, 0, 0)
	l.Title = title
	l.SetShowHelp(true)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.KeyMap.Quit = key.NewBinding()
	l.KeyMap.ForceQuit = key.NewBinding()
	l.KeyMap.ShowFullHelp = key.NewBinding()
	l.KeyMap.CloseFullHelp = key.NewBinding()
	return l
}

func toListItems(items []MenuItem) []list.Item {
	out := make([]list.Item, len(items))
	for i, it := range items {
		out[i] = it
	}
	return out
}

func selectedData[T any](l list.Model) (T, bool) {
	var zero T
	item, ok := l.SelectedItem().(MenuItem)
	if !ok {
		return zero, false
	}
	val, ok := item.Data.(T)
	return val, ok
}

func (m Menu) Update(msg tea.Msg, parentHeight int) (Menu, tea.Cmd) {
	if m.Height == 0 && parentHeight >= minMenuHeight {
		_, border := style.MenuBorder.GetFrameSize()
		m.List.SetSize(m.Width, parentHeight-border)
	}

	if m.Width == 0 {
		m.List.SetWidth(minMenuWidth)
	}

	if m.List.Width() > maxMenuWidth {
		m.List.SetWidth(maxMenuWidth)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
