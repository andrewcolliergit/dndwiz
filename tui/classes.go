package tui

import (
	"fmt"
	"strings"

	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewClassesModel(c *character.Character) tea.Model {
	classes := character.GetAllClasses()
	items := make(SidebarItemList, 0, len(classes))

	for i, clsInfo := range classes {
		items = append(items, SidebarItem{
			title: clsInfo.Name,
			desc:  clsInfo.Desc,
			data:  character.Class(i + 1),
		})
	}

	return newSelectionPicker(
		c,
		"Classes",
		"Select Class",
		items,
		false,
		func(c *character.Character, val character.Class) tea.Cmd {
			c.Class = val
			cls := character.GetClass(val)
			if cls.PrimaryAbility.RequiresChoice() {
				return switchToView(viewPrimaryAbilitiesSelection)
			} else {
				c.PrimaryAbilities = cls.PrimaryAbility.From
			}

			return switchToView(viewCharsheet)
		},
		displayClassInfo,
	)
}

func skillProficienciesStr(class character.Class) string {
	sp := character.GetClass(class).SkillProficiency

	var skillsList []string
	for _, skill := range sp.From {
		skillName := character.GetSkill(skill).Name
		skillsList = append(skillsList, skillStyle.Render(skillName))
	}

	skills := strings.Join(skillsList, " ")

	header := fmt.Sprintf("Skill Proficiencies (%d)", sp.Has)

	return labelStyle.Render(header) + " " + skills
}

func savingThrowStr(st character.ClassSavingThrow) string {
	var stList []string
	for _, throw := range st {
		t := character.GetAbility(throw).Abbr
		stList = append(stList, badgeStyle.Render(t))
	}
	info := strings.Join(stList, " ")
	return labelStyle.Render("Saving Throw") + " " + info
}

func displayClassInfo(c *character.Character, class character.Class, contentWidth int) string {
	cls := character.GetClass(class)

	// Header & Accents
	title := titleStyle.Render(cls.Name)
	primaryAbilities := labelStyle.Render("Primary Ability: ") + badgeStyle.Render(cls.PrimaryAbility.String())
	hitPointDie := labelStyle.Render("Hit Point Die: ") + character.GetHitPointDieString(class)
	weaponProficiencies := labelStyle.Render("Weapon Proficiencies: ") + cls.WeaponProficiencies
	armorTraining := labelStyle.Render("Armor Training: ") + cls.ArmorTraining
	startingEquipment := labelStyle.Render("Starting Equipment: ") + cls.StartingEquipment.String()

	// Body content
	desc := descStyle.Render(cls.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		renderTitle(title, desc),
		renderDivider(contentWidth-2),
		primaryAbilities,
		hitPointDie,
		savingThrowStr(character.GetClass(class).SavingThrow),
		skillProficienciesStr(class),
		weaponProficiencies,
		armorTraining,
		startingEquipment,
	)
}
