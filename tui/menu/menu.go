// Package menu/menu.go
package menu

import (
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Header string
	Desc   string
	Data   any
}

type List []Item

func (i Item) Title() string       { return i.Header }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Header }

type Menu struct {
	List  list.Model
	Width int
}

func (m Menu) SelectedData() any {
	if selected := m.List.SelectedItem(); selected != nil {
		if item, ok := selected.(Item); ok {
			return item.Data
		}
	}
	return nil
}

func GetItemData[T any](m Menu) (T, bool) {
	var zero T
	selected := m.List.SelectedItem()
	if selected == nil {
		return zero, false
	}

	item, ok := selected.(Item)
	if !ok {
		return zero, false
	}

	val, ok := item.Data.(T)
	return val, ok
}

func New(title string, items List, width int, showDesc bool) Menu {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = showDesc
	l := list.New(listItems, delegate, width, 0)
	l.Title = title
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)

	return Menu{
		List:  l,
		Width: width,
	}
}

func (m *Menu) SetItems(items List) {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	m.List.SetItems(listItems)
}

func (m Menu) Update(msg tea.Msg, parentHeight int) (Menu, tea.Cmd) {
	if parentHeight > 4 {
		m.List.SetSize(m.Width-2, parentHeight-4)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Menu) View() string {
	return style.GetSidebarMenuStyle(m.Width).Render(m.List.View())
}
