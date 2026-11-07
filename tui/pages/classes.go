// Package pages/classes.go
package pages

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/class"
	"dndwiz/character/score"
	"dndwiz/character/skill"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ClassSelectionConfig = shell.ShellConfig[class.Class]{
	ContentTitle: "Classes",
	MenuTitle:    "Select Class",
	Layout:       shell.LayoutSidebar,
	ShowDesc:     false,
	HelpContent: func(w int) string {
		p := func(s string) string { return style.TypesetBody(s, w) }
		term := func(name, text string) string {
			return lipgloss.JoinVertical(lipgloss.Left,
				style.Label.Render(name),
				p(text),
				"",
			)
		}

		return lipgloss.JoinVertical(lipgloss.Left,
			style.Header.Render("What is a class?"),
			p("Your class is what your character does: fight with a sword, cast spells, sneak through "+
				"shadows, heal allies. It is the biggest choice you make, because it decides most of what "+
				"you can do on your turn in the game. Your species says where you come from, and your "+
				"class says what you do."),
			"",
			p("Each class also gives you new abilities as your character levels up. This app shows what "+
				"a level 1 character starts with."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each class. Nothing is chosen until you press "+
				"Enter or Space, so browse as much as you like. You can come back and change your class "+
				"at any time."),
			"",
			p("Most classes take you straight back to the character sheet. The Fighter asks one "+
				"follow-up question: whether Strength or Dexterity is your primary ability."),
			"",
			p("If the details are longer than the screen, scroll them with Shift+↑/↓."),
			style.TypesetDivider(w),

			style.Header.Render("What the details mean"),
			term("Primary Ability",
				"The ability score that matters most for this class. If you put your best score here, "+
					"the class's main features work better. Some classes list two abilities. Monks, "+
					"Paladins, and Rangers need both, so both are set for you. Fighters need only one of "+
					"their two, so you choose."),
			term("Hit Point Die",
				"Your hit points (HP) are how much damage you can take before you fall unconscious. "+
					"The die size shows how tough the class is: d12 is the toughest, d6 the most fragile. "+
					"At 1st level you simply start with the die's maximum plus your Constitution modifier."),
			term("Saving Throws",
				"When a spell, trap, or monster tries to affect you, you often roll a saving throw to "+
					"resist it. Each class is trained in two abilities and adds a bonus when rolling "+
					"those saves."),
			term("Skill Proficiencies",
				"Skills are things you can be good at, like Stealth, Persuasion, or Perception. The "+
					"number is how many you may choose, and the list shows which ones this class can "+
					"pick from. This app displays the options but does not ask you to choose yet."),
			term("Weapon Proficiencies",
				"Which weapons you are trained with. Simple weapons are everyday gear like clubs and "+
					"daggers. Martial weapons are military arms like longswords and greataxes. Being "+
					"trained adds your proficiency bonus when you attack."),
			term("Armor Training",
				"Which armor and shields you can wear without penalty. Heavier armor protects you more "+
					"but is only available to some classes. \"None\" means you rely on your agility, "+
					"your abilities, or magic instead."),
			term("Starting Equipment",
				"You choose one bundle, written as (A), (B), and so on. One option is usually a set of "+
					"gear and the last is a pile of gold pieces (GP) to spend on your own gear. If you "+
					"are new, pick the gear."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see"),
			term("Spellcasting",
				"Some classes use magic. Spells come from a particular ability: Wizards use "+
					"Intelligence, Clerics and Druids use Wisdom, and Bards, Sorcerers, and Warlocks use "+
					"Charisma. Cantrips are simple spells you can cast as often as you like. Other "+
					"spells use up spell slots, which come back after rest."),
			term("Long Rest",
				"About eight hours of sleep and light activity. Many class abilities recharge after one."),
			term("Modifier",
				"The bonus or penalty an ability score gives you. A score of 10 gives +0, and every "+
					"two points above that adds +1."),
			style.TypesetDivider(w),

			style.Header.Render("The classes"),
			term("Barbarian",
				"A fierce warrior who fights in a Rage, taking less damage and hitting harder. Very "+
					"tough (d12) and simple to play. Good if you want to charge in and not worry about "+
					"spells."),
			term("Bard",
				"A performer whose music and words are magic. You cast spells, inspire allies to do "+
					"better on their rolls, and are skilled at many things. A flexible support class, "+
					"best for players who like talking and creativity."),
			term("Cleric",
				"A servant of a god who channels divine power. You heal allies, protect the group, and "+
					"can wear sturdy armor. A reliable choice if you like being the one the team depends on."),
			term("Druid",
				"A guardian of nature. You cast nature spells and can transform into animals with Wild "+
					"Shape. Rich and flexible, but with more choices to manage."),
			term("Fighter",
				"A master of weapons and armor. Fighters are sturdy, use every weapon, and have few "+
					"rules to track. One of the easiest classes to learn, which makes it a good first "+
					"character."),
			term("Monk",
				"A martial artist who fights with fists and speed, needing no armor. Quick and "+
					"mobile, powered by both Dexterity and Wisdom, so your ability scores matter more here."),
			term("Paladin",
				"A holy knight bound by an oath. You wear heavy armor, hit hard, and can heal with a "+
					"touch. Combines fighting with some magic, powered by Strength and Charisma."),
			term("Ranger",
				"A hunter and scout of the wild. You use bows and blades, track enemies, and cast "+
					"a little nature magic. Good for a character who feels at home outdoors."),
			term("Rogue",
				"A stealthy specialist. You strike from hiding for extra damage (Sneak Attack), are "+
					"skilled at many tasks, and excel at sneaking, locks, and traps. Light armor and "+
					"d8 hit points, so you avoid hits instead of absorbing them."),
			term("Sorcerer",
				"A spellcaster born with magic in their blood. You know fewer spells than a Wizard but "+
					"can reshape them. Fragile (d6) and best kept behind the party."),
			term("Warlock",
				"A spellcaster who made a pact with a powerful being in return for magic. Your spells "+
					"are fewer but strong and return quickly after a short rest. Unusual and "+
					"story-rich, so it is best if you like a dramatic backstory."),
			term("Wizard",
				"A scholar of magic with the widest range of spells, which you record in a spellbook. "+
					"The most powerful spellcaster, but fragile (d6) and complex, so it is better for "+
					"players who enjoy planning."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Think about what you want to do in a fight and out of one. If you want something "+
				"straightforward, Fighter, Barbarian, and Rogue have few rules to remember. If you "+
				"like magic, Cleric is forgiving, while Wizard and Druid reward experience. You can "+
				"always change your class on this screen before you finish."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},

	GetItems: func(c *character.Character) []shell.MenuItem {
		classes := class.GetAll()
		items := make([]shell.MenuItem, 0, len(classes))

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

		if c.AbilityScoreMethod == score.StdArray {
			c.AbilityScore = cls.StdArray // keep the array in step with the class
		}

		if cls.PrimaryAbility.RequiresChoice() {
			c.PrimaryAbilities = ability.Abilities{cls.PrimaryAbility.From[0]} // default; user can change
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
		info := class.Get(selected)

		// Header & Accents
		primaryAbilities := style.Label.Render("Primary Ability: ") + style.TypesetBadges([]string{info.PrimaryAbility.String()}, contentWidth)
		hitPointDie := style.Label.Render("Hit Point Die: ") + class.GetHitPointDieString(selected)
		weaponProficiencies := style.Label.Render("Weapon Proficiencies: ") + info.WeaponProficiencies
		armorTraining := style.Label.Render("Armor Training: ") + info.ArmorTraining
		startingEquipment := style.Label.Render("Starting Equipment: ") + info.StartingEquipment.String()

		return lipgloss.JoinVertical(
			lipgloss.Left,

			style.TypesetTitle(info.Name, info.Desc, contentWidth),
			style.TypesetDivider(contentWidth-2),
			primaryAbilities,
			hitPointDie,
			savingThrowStr(class.Get(selected).SavingThrow),
			skillProficienciesStr(selected),
			weaponProficiencies,
			armorTraining,
			startingEquipment,
		)
	},
}

// --- Helper Methods ---

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
