package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewBackgroundModel(c *character.Character) tea.Model {
	backgrounds := character.GetAllBackgrounds()
	items := make(SidebarItemList, 0, len(backgrounds))

	for i, bgInfo := range backgrounds {
		items = append(items, SidebarItem{
			title: bgInfo.Name,
			desc:  "",
			data:  character.Background(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Backgrounds",
		"Select Background",
		items,
		false,
		func(c *character.Character, val character.Background) tea.Cmd {
			c.Background = val
			return switchToView(viewCharsheet)
		},
		displayBackgroundInfo,
	)
}

func displayBackgroundInfo(c *character.Character, bg character.Background) string {
	bgi := character.GetBackground(bg)

	title := titleStyle.Render(bgi.Name)
	desc := lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("246")).Render(bgi.Desc)

	abilities := labelStyle.Render("Ability Scores: ") + character.GetBackgroundAbilities(bg)
	feat := labelStyle.Render("Feat: ") + badgeStyle.Render(bgi.Feat.String())
	skills := labelStyle.Render("Skills: ") + character.GetBackgroundSkills(bg)
	tools := labelStyle.Render("Tools: ") + bgi.Tool
	equip := labelStyle.Render("Equipment Options:\n") + character.GetBackgroundEquipmentOptions(bg)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		desc,
		renderDivider(40),
		abilities,
		feat,
		skills,
		tools,
		renderDivider(40),
		equip,
	)
}
