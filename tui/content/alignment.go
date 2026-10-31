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

// AlignmentSelectionConfig is the declarative blueprint for selecting an alignment.
var AlignmentSelectionConfig = shell.ShellConfig[alignment.Alignment]{
	Title:     "Alignments",
	MenuTitle: "Select Alignment",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	// 1. Populate alignment options dynamically
	GetItems: func(c *character.Character) shell.MenuItemList {
		alignments := alignment.GetAll()
		items := make(shell.MenuItemList, 0, len(alignments))

		for i, alnInfo := range alignments {
			items = append(items, shell.MenuItem{
				Header: alnInfo.Name,
				Desc:   alnInfo.Desc,
				Data:   alignment.Alignment(i + 1),
			})
		}
		return items
	},

	// 2. Handle alignment selection and navigation back to character sheet
	OnSelect: func(c *character.Character, val alignment.Alignment) tea.Cmd {
		c.Alignment = val
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	// 3. Render alignment detail pane
	Content: func(c *character.Character, a alignment.Alignment, contentWidth int) string {
		return displayAlignmentInfo(a)
	},
}

// --- Helper Formatting Functions ---

func displayAlignmentInfo(a alignment.Alignment) string {
	aln := alignment.Get(a)

	// Header & Accents
	title := style.Title.Render(aln.Name)

	// Body content
	desc := style.Desc.Render(aln.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		desc,
	)
}
