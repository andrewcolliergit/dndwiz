// Package content/alignment
package content

import (
	"dndwiz/character"
	"dndwiz/character/alignment"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AlignmentSelectionConfig = shell.ShellConfig[alignment.Alignment]{
	Title:     "Alignments",
	MenuTitle: "Select Alignment",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	GetItems: func(c *character.Character) shell.MenuItemList {
		alignments := alignment.GetAll()
		items := make(shell.MenuItemList, 0, len(alignments))

		for i, alnInfo := range alignments {
			items = append(items, shell.MenuItem{
				Label: alnInfo.Name,
				Desc:  alnInfo.Desc,
				Data:  alignment.Alignment(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, val alignment.Alignment) tea.Cmd {
		c.Alignment = val
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, a alignment.Alignment, contentWidth int) string {
		aln := alignment.Get(a)
		title := style.Title.Render(aln.Name)
		desc := style.Desc.Render(aln.Desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
		)
	},
}
