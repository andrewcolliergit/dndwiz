package content

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/class"
	"dndwiz/character/skill"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ClassSelectionConfig = shell.ShellConfig[class.Class]{
	Title:     "Classes",
	MenuTitle: "Select Class",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  false,

	GetItems: func(c *character.Character) shell.MenuItemList {
		classes := class.GetAll()
		items := make(shell.MenuItemList, 0, len(classes))

		for i, clsInfo := range classes {
			items = append(items, shell.MenuItem{
				Label: clsInfo.Name,
				Desc:  clsInfo.Desc,
				Data:  class.Class(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selected class.Class) tea.Cmd {
		c.Class = selected
		cls := class.Get(selected)

		if cls.PrimaryAbility.RequiresChoice() {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.ClassAbilitiesSelection)
			}
		}

		c.PrimaryAbilities = cls.PrimaryAbility.From
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, selected class.Class, contentWidth int) string {
		return displayClassInfo(c, selected, contentWidth)
	},
}

func skillProficienciesStr(c class.Class) string {
	sp := class.Get(c).SkillProficiency

	var skillsList []string
	for _, s := range sp.From {
		skillName := skill.Get(s).Name
		skillsList = append(skillsList, style.Skill.Render(skillName))
	}

	skills := strings.Join(skillsList, " ")
	header := fmt.Sprintf("Skill Proficiencies (%d)", sp.Has)

	return style.Label.Render(header) + " " + skills
}

func savingThrowStr(st class.SavingThrow) string {
	var stList []string
	for _, throw := range st {
		t := ability.Get(throw).Abbr
		stList = append(stList, style.Badge.Render(t))
	}
	info := strings.Join(stList, " ")
	return style.Label.Render("Saving Throw") + " " + info
}

func displayClassInfo(c *character.Character, cls class.Class, contentWidth int) string {
	info := class.Get(cls)

	// Header & Accents
	title := style.Title.Render(info.Name)
	primaryAbilities := style.Label.Render("Primary Ability: ") + style.Badge.Render(info.PrimaryAbility.String())
	hitPointDie := style.Label.Render("Hit Point Die: ") + class.GetHitPointDieString(cls)
	weaponProficiencies := style.Label.Render("Weapon Proficiencies: ") + info.WeaponProficiencies
	armorTraining := style.Label.Render("Armor Training: ") + info.ArmorTraining
	startingEquipment := style.Label.Render("Starting Equipment: ") + info.StartingEquipment.String()

	// Body content
	desc := style.Desc.Render(info.Desc)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		style.RenderTitle(title, desc),
		style.Divider(contentWidth-2),
		primaryAbilities,
		hitPointDie,
		savingThrowStr(class.Get(cls).SavingThrow),
		skillProficienciesStr(cls),
		weaponProficiencies,
		armorTraining,
		startingEquipment,
	)
}
