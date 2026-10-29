package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewAlignmentModel(c *character.Character) tea.Model {
	alignments := character.GetAllAlignments()
	items := make(SidebarItemList, 0, len(alignments))

	for i, alnInfo := range alignments {
		items = append(items, SidebarItem{
			title: alnInfo.Name,
			desc:  alnInfo.Desc,
			data:  character.Alignment(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Alignments",
		"Select Alignment",
		items,
		false,
		func(c *character.Character, val character.Alignment) tea.Cmd {
			c.Alignment = val
			return switchToView(viewCharsheet)
		},
		displayAlignmentInfo,
	)
}

func displayAlignmentInfo(c *character.Character, a character.Alignment, contentWidth int) string {
	aln := character.GetAlignment(a)

	// Header & Accents
	title := titleStyle.Render(aln.Name)

	// Body content
	desc := descStyle.Render(aln.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		// renderDivider(contentWidth-2),
		desc,
	)
}
