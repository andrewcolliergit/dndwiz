package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SidebarItem struct {
	title string
	desc  string
	data  any
}

type SidebarItemList []SidebarItem

func (i SidebarItem) Title() string       { return i.title }
func (i SidebarItem) Description() string { return i.desc }
func (i SidebarItem) FilterValue() string { return i.title }

type sidebarModel struct {
	list  list.Model
	width int
}

func (m sidebarModel) SelectedData() any {
	if selected := m.list.SelectedItem(); selected != nil {
		if item, ok := selected.(SidebarItem); ok {
			return item.data
		}
	}
	return nil
}

func GetItemData[T any](m sidebarModel) (T, bool) {
	var zero T
	selected := m.list.SelectedItem()
	if selected == nil {
		return zero, false
	}

	item, ok := selected.(SidebarItem)
	if !ok {
		return zero, false
	}

	val, ok := item.data.(T)
	return val, ok
}

func newSidebar(title string, items SidebarItemList, width int, showDesc bool) sidebarModel {
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

	return sidebarModel{
		list:  l,
		width: width,
	}
}

func (m *sidebarModel) SetItems(items SidebarItemList) {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	m.list.SetItems(listItems)
}

func (m sidebarModel) Update(msg tea.Msg, parentHeight int) (sidebarModel, tea.Cmd) {
	if parentHeight > 4 {
		m.list.SetSize(m.width-2, parentHeight-4)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m sidebarModel) View() string {
	sidebarStyle := lipgloss.NewStyle().
		Width(m.width).
		//Border(lipgloss.NormalBorder(), false, true, false, false).
		//BorderForeground(lipgloss.Color("240"))
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("62"))
	return sidebarStyle.Render(m.list.View())
}
