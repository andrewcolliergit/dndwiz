// Package equipmentselection
package equipmentselection

import (
	"dndwiz/character"
	"dndwiz/character/class"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewModel(c *character.Character) tea.Model {
	classes := class.GetAll()
	items := make(menu.List, 0, len(classes))

	for i, clsInfo := range classes {
		items = append(items, menu.Item{
			Header: clsInfo.Name,
			Desc:   clsInfo.Desc,
			Data:   class.Class(i + 1),
		})
	}

	return menu.NewPicker(
		c,
		"Classes",
		"Select Class",
		items,
		false,
		func(c *character.Character, val class.Class) tea.Cmd {
			c.Class = val
			cls := class.Get(val)
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

func displayEquipmentInfo(c *character.Character, cls class.Class, contentWidth int) string {
	info := class.Get(cls)

	// Header & Accents
	title := style.Title.Render(info.Name)
	primaryAbilities := style.Label.Render("Primary Ability: ") + style.Badge.Render(info.PrimaryAbility.String())
	hitPointDie := style.Label.Render("Hit Point Die: ") + class.GetHitPointDieString(cls)
	savingThrow := style.Label.Render("Saving Throw Proficiency: ") + info.SavingThrow.String()
	skillProficiencies := style.Label.Render("Skill Proficiences: ") + info.SkillProficiency.String()
	weaponProficiencies := style.Label.Render("Weapon Proficiencies: ") + info.WeaponProficiencies
	armorTraining := style.Label.Render("Armor Training: ") + info.ArmorTraining
	startingEquipment := style.Label.Render("Starting Equipment: ") + info.StartingEquipment.String()

	// Body content
	desc := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Render(info.Desc)

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
