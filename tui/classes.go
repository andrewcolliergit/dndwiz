package tui

import (
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

func displayClassInfo(c *character.Character, class character.Class) string {
	cls := character.GetClass(class)

	// Header & Accents
	title := titleStyle.Render(cls.Name)
	primaryAbilities := labelStyle.Render("Primary Ability: ") + badgeStyle.Render(cls.PrimaryAbility.String())
	hitPointDie := labelStyle.Render("Hit Point Die: ") + character.GetHitPointDieString(class)
	savingThrow := labelStyle.Render("Saving Throw Proficiency: ") + cls.SavingThrow.String()
	skillProficiencies := labelStyle.Render("Skill Proficiences: ") + cls.SkillProficiency.String()
	weaponProficiencies := labelStyle.Render("Weapon Proficiencies: ") + cls.WeaponProficiencies
	armorTraining := labelStyle.Render("Armor Training: ") + cls.ArmorTraining
	startingEquipment := labelStyle.Render("Starting Equipment: ") + cls.StartingEquipment.String()

	// Body content
	desc := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(cls.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		primaryAbilities,
		hitPointDie,
		savingThrow,
		skillProficiencies,
		weaponProficiencies,
		armorTraining,
		startingEquipment,
		renderDivider(40),
		desc,
	)
}
