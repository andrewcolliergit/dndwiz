// Package background
package background

import (
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/background"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	backgrounds := background.GetAll()
	items := make(menu.List, 0, len(backgrounds))

	for i, bgInfo := range backgrounds {
		items = append(items, menu.Item{
			Header: bgInfo.Name,
			Desc:   "",
			Data:   background.Background(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Backgrounds",
		"Select Background",
		items,
		false,
		func(c *character.Character, val background.Background) tea.Cmd {
			c.Background = val
			return pager.SwitchTo(pager.Charsheet)
		},
		displayBackgroundInfo,
	)
}

func displayBackgroundInfo(c *character.Character, bg background.Background, contentWidth int) string {
	var abilitiesArr []string
	for _, a := range background.Get(bg).Abilities {
		abilityAbbr := ability.Get(a).Abbr
		abilitiesArr = append(abilitiesArr, style.Badge.Render(abilityAbbr))
	}
	abilitiesStr := strings.Join(abilitiesArr, " ")
	bgi := background.Get(bg)

	title := style.Title.Render(bgi.Name)
	desc := style.Desc.Render(bgi.Desc)

	abilities := style.Label.Render("Ability Scores: ") + abilitiesStr
	feat := style.Label.Render("Feat: ") + style.Badge.Render(bgi.Feat.String())
	skills := style.Label.Render("Skills: ") + background.GetSkills(bg)
	tools := style.Label.Render("Tools: ") + bgi.Tool
	equip := background.EquipmentOptions(bg)

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
