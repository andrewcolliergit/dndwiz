// Package alignment
package alignment

import (
	"dndwiz/character"
	"dndwiz/character/alignment"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	alignments := alignment.GetAll()
	items := make(menu.List, 0, len(alignments))

	for i, alnInfo := range alignments {
		items = append(items, menu.Item{
			Header: alnInfo.Name,
			Desc:   alnInfo.Desc,
			Data:   alignment.Alignment(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Alignments",
		"Select Alignment",
		items,
		false,
		func(c *character.Character, val alignment.Alignment) tea.Cmd {
			c.Alignment = val
			return pager.SwitchTo(pager.Charsheet)
		},
		displayAlignmentInfo,
	)
}

func displayAlignmentInfo(c *character.Character, a alignment.Alignment, contentWidth int) string {
	aln := alignment.Get(a)

	// Header & Accents
	title := style.Title.Render(aln.Name)

	// Body content
	desc := style.Desc.Render(aln.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		// renderDivider(contentWidth-2),
		desc,
	)
}
