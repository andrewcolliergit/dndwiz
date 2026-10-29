package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type menuStyle int

const (
	menuStyleNone menuStyle = iota
	menuStyleSimple
	menuStyleSidebar
)

type MenuItem struct {
	title string
	desc  string
	data  any
}

type MenuItemList []MenuItem

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.desc }
func (i MenuItem) FilterValue() string { return i.title }

type Menu struct {
	list  list.Model
	width int
	style menuStyle
}

func (m Menu) SelectedData() any {
	if selected := m.list.SelectedItem(); selected != nil {
		if item, ok := selected.(MenuItem); ok {
			return item.data
		}
	}
	return nil
}

func GetItemData[T any](m Menu) (T, bool) {
	var zero T
	selected := m.list.SelectedItem()
	if selected == nil {
		return zero, false
	}

	item, ok := selected.(MenuItem)
	if !ok {
		return zero, false
	}

	val, ok := item.data.(T)
	return val, ok
}

func newMenu(style menuStyle, title string, items MenuItemList, width int, showDesc bool) Menu {
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
		list:  l,
		width: width,
		style: style,
	}
}

func (m *Menu) SetItems(items MenuItemList) {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	m.list.SetItems(listItems)
}

func (m Menu) Update(msg tea.Msg, parentHeight int) (Menu, tea.Cmd) {
	if parentHeight > 4 {
		m.list.SetSize(m.width-2, parentHeight-4)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Menu) View() string {
	return getSidebarMenuStyle(m.width).Render(m.list.View())
}
