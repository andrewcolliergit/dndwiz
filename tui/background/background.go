// Package background
package background

import (
	"strings"

	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	backgrounds := character.GetAllBackgrounds()
	items := make(menu.List, 0, len(backgrounds))

	for i, bgInfo := range backgrounds {
		items = append(items, menu.Item{
			Header: bgInfo.Name,
			Desc:   "",
			Data:   character.Background(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Backgrounds",
		"Select Background",
		items,
		false,
		func(c *character.Character, val character.Background) tea.Cmd {
			c.Background = val
			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displayBackgroundInfo,
	)
}

func displayBackgroundInfo(c *character.Character, bg character.Background, contentWidth int) string {
	var abilitiesArr []string
	for _, ability := range character.GetBackground(bg).Abilities {
		abilityAbbr := character.GetAbility(ability).Abbr
		abilitiesArr = append(abilitiesArr, style.Badge.Render(abilityAbbr))
	}
	abilitiesStr := strings.Join(abilitiesArr, " ")
	bgi := character.GetBackground(bg)

	title := style.Title.Render(bgi.Name)
	desc := style.Desc.Render(bgi.Desc)

	abilities := style.Label.Render("Ability Scores: ") + abilitiesStr
	feat := style.Label.Render("Feat: ") + style.Badge.Render(bgi.Feat.String())
	skills := style.Label.Render("Skills: ") + character.GetBackgroundSkills(bg)
	tools := style.Label.Render("Tools: ") + bgi.Tool
	equip := character.GetBackgroundEquipmentOptions(bg)

	divWidth := contentWidth - 2

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		desc,
		style.Divider(divWidth),
		abilities,
		feat,
		skills,
		tools,
		style.Header.Render("Equipment Options"),
		equip,
	)
}
