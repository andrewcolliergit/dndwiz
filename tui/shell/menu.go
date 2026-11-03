// Package shell/menu.go
package shell

import (
	"dndwiz/tui/style"

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

type MenuItemList []MenuItem

func (i MenuItem) Title() string       { return i.Label }
func (i MenuItem) Description() string { return i.Desc }
func (i MenuItem) FilterValue() string { return i.Label }

type Menu struct {
	List   list.Model
	Width  int
	Height int
}

func NewMenu(title string, items MenuItemList, width int, height int, showDesc bool) Menu {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = showDesc
	l := list.New(listItems, delegate, width, height)
	l.Title = title
	l.KeyMap.Quit.SetEnabled(false)
	l.KeyMap.ForceQuit.SetEnabled(false)
	l.SetShowHelp(true)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return Menu{
		List:   l,
		Width:  width,
		Height: height,
	}
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

func (m Menu) View() string {
	return style.MenuBorder.Render(m.List.View())
}

func (m Menu) GetWidth() int {
	w := m.List.Width()
	if w <= 0 {
		w = m.Width
	}
	if w <= 0 {
		w = minMenuWidth
	}

	frameWidth, _ := style.MenuBorder.GetFrameSize()
	return frameWidth + w
}

func (m Menu) GetHeight() int {
	h := m.List.Height()
	_, frameHeight := style.MenuBorder.GetFrameSize()
	return frameHeight + h
}

// --- Helper Methods ---

func (m *Menu) setItems(items MenuItemList) {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	m.List.SetItems(listItems)
}

func getItemData[T any](m Menu) (T, bool) {
	var zero T
	selected := m.List.SelectedItem()
	if selected == nil {
		return zero, false
	}

	item, ok := selected.(MenuItem)
	if !ok {
		return zero, false
	}

	val, ok := item.Data.(T)
	return val, ok
}
