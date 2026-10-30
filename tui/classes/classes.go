// Package classes
package classes

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	classes := character.GetAllClasses()
	items := make(menu.List, 0, len(classes))

	for i, clsInfo := range classes {
		items = append(items, menu.Item{
			Header: clsInfo.Name,
			Desc:   clsInfo.Desc,
			Data:   character.Class(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Classes",
		"Select Class",
		items,
		false,
		func(c *character.Character, val character.Class) tea.Cmd {
			c.Class = val
			cls := character.GetClass(val)
			if cls.PrimaryAbility.RequiresChoice() {
				return pager.SwitchTo(pager.ViewPrimaryAbilitiesSelection)
			} else {
				c.PrimaryAbilities = cls.PrimaryAbility.From
			}

			return pager.SwitchTo(pager.ViewCharsheet)
		},
		displayClassInfo,
	)
}

func skillProficienciesStr(class character.Class) string {
	sp := character.GetClass(class).SkillProficiency

	var skillsList []string
	for _, skill := range sp.From {
		skillName := character.GetSkill(skill).Name
		skillsList = append(skillsList, style.Skill.Render(skillName))
	}

	skills := strings.Join(skillsList, " ")

	header := fmt.Sprintf("Skill Proficiencies (%d)", sp.Has)

	return style.Label.Render(header) + " " + skills
}

func savingThrowStr(st character.ClassSavingThrow) string {
	var stList []string
	for _, throw := range st {
		t := character.GetAbility(throw).Abbr
		stList = append(stList, style.Badge.Render(t))
	}
	info := strings.Join(stList, " ")
	return style.Label.Render("Saving Throw") + " " + info
}

func displayClassInfo(c *character.Character, class character.Class, contentWidth int) string {
	cls := character.GetClass(class)

	// Header & Accents
	title := style.Title.Render(cls.Name)
	primaryAbilities := style.Label.Render("Primary Ability: ") + style.Badge.Render(cls.PrimaryAbility.String())
	hitPointDie := style.Label.Render("Hit Point Die: ") + character.GetHitPointDieString(class)
	weaponProficiencies := style.Label.Render("Weapon Proficiencies: ") + cls.WeaponProficiencies
	armorTraining := style.Label.Render("Armor Training: ") + cls.ArmorTraining
	startingEquipment := style.Label.Render("Starting Equipment: ") + cls.StartingEquipment.String()

	// Body content
	desc := style.Desc.Render(cls.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		style.RenderTitle(title, desc),
		style.Divider(contentWidth-2),
		primaryAbilities,
		hitPointDie,
		savingThrowStr(character.GetClass(class).SavingThrow),
		skillProficienciesStr(class),
		weaponProficiencies,
		armorTraining,
		startingEquipment,
	)
}
