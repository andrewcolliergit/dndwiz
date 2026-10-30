// Package equipmentselection
package equipmentselection

import (
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
		displayEquipmentInfo,
	)
}

func displayEquipmentInfo(c *character.Character, class character.Class, contentWidth int) string {
	cls := character.GetClass(class)

	// Header & Accents
	title := style.Title.Render(cls.Name)
	primaryAbilities := style.Label.Render("Primary Ability: ") + style.Badge.Render(cls.PrimaryAbility.String())
	hitPointDie := style.Label.Render("Hit Point Die: ") + character.GetHitPointDieString(class)
	savingThrow := style.Label.Render("Saving Throw Proficiency: ") + cls.SavingThrow.String()
	skillProficiencies := style.Label.Render("Skill Proficiences: ") + cls.SkillProficiency.String()
	weaponProficiencies := style.Label.Render("Weapon Proficiencies: ") + cls.WeaponProficiencies
	armorTraining := style.Label.Render("Armor Training: ") + cls.ArmorTraining
	startingEquipment := style.Label.Render("Starting Equipment: ") + cls.StartingEquipment.String()

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
		style.Divider(contentWidth-2),
		desc,
	)
}
