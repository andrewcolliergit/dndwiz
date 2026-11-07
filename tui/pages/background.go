// Package pages/background.go
package pages

import (
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/background"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BackgroundSelectionConfig = shell.ShellConfig[background.Background]{
	ContentTitle: "Backgrounds",
	MenuTitle:    "Select Background",
	Layout:       shell.LayoutSidebar,
	MenuWidth:    32,
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
			style.Header.Render("What is a background?"),
			p("Your background is what your character did before they became an adventurer. "+
				"Were they raised in a temple, trained as a soldier, or scraping by in the "+
				"underbelly of a city? That history gives you a head start in a few areas."),
			"",
			p("A background is about the past, not the present. It is separate from your class, "+
				"so any background works with any class. Pick the one that tells the story you want."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each background. Nothing is chosen until "+
				"you press Enter or Space. After you choose, you return to the character sheet. You "+
				"can come back and change your background at any time."),
			style.TypesetDivider(w),

			style.Header.Render("What the details mean"),
			term("Ability Scores",
				"Each background lists three abilities. In the rules, you raise one of them by 2 and "+
					"another by 1, or raise all three by 1, and no score can go above 20. This is "+
					"how your history makes you a little stronger, smarter, or tougher."),
			term("Feat",
				"A feat is a special talent. Your background gives you one for free, called an "+
					"Origin feat. The Feat entry in the glossary below explains each one."),
			term("Skills",
				"You become trained in two skills, like Stealth or Insight. Being trained means you "+
					"add a bonus when you roll for that skill."),
			term("Tools",
				"You become trained with one kind of tool, such as Thieves' Tools or Calligrapher's "+
					"Supplies. Tools let you do things like pick locks or write beautiful documents."),
			term("Equipment Options",
				"You choose one bundle. Option (A) is a set of useful gear plus some gold pieces "+
					"(GP). Option (B) is only gold, to spend on whatever you like. If you are new, "+
					"pick (A), because it comes ready to use."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see"),
			term("Proficiency",
				"Being trained in something. When you use a skill, tool, or weapon you are proficient "+
					"in, you add a bonus to the roll."),
			term("Cantrip",
				"A simple spell you can cast as often as you like."),
			term("Initiative",
				"A roll that decides who acts first when a fight begins."),
			style.TypesetDivider(w),

			style.Header.Render("The backgrounds"),
			term("Acolyte",
				"You grew up serving a temple or faith. Your abilities are Intelligence, Wisdom, "+
					"and Charisma. You are trained in Insight (reading people) and Religion, and with "+
					"Calligrapher's Supplies. Your feat is Magic Initiate: you learn two cantrips and "+
					"one first-level spell from the Cleric spell list, and you can cast that spell "+
					"once per long rest. A good fit for priests and healers."),
			term("Criminal",
				"You lived outside the law, whether as a pickpocket, a smuggler, or a burglar. "+
					"Your abilities are Dexterity, Constitution, and Intelligence. You are trained in "+
					"Sleight of Hand and Stealth, and with Thieves' Tools. Your feat is Alert, which "+
					"adds your proficiency bonus to initiative so you tend to act first. A good fit "+
					"for sneaky characters."),
			term("Sage",
				"You spent years studying in libraries and academies. Your abilities are "+
					"Constitution, Intelligence, and Wisdom. You are trained in Arcana (magical "+
					"knowledge) and History, and with Calligrapher's Supplies. Your feat is Magic "+
					"Initiate, this time using the Wizard spell list. A good fit for scholars and "+
					"spellcasters."),
			term("Soldier",
				"You served in an army or a militia. Your abilities are Strength, Dexterity, and "+
					"Constitution. You are trained in Athletics and Intimidation, and with a gaming "+
					"set of your choice. Your feat is Savage Attacker: once per turn, you can roll a "+
					"weapon's damage dice twice and use the higher result. A good fit for fighters."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Choose the story you like best. If you want a bonus that fits your class, look at "+
				"the abilities: a Soldier's Strength and Dexterity help a warrior, and a Sage's "+
				"Intelligence helps a Wizard. Any pairing works, so none of these is a wrong choice."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		backgrounds := background.GetAll()
		items := make([]shell.MenuItem, 0, len(backgrounds))

		for i, bgInfo := range backgrounds {
			items = append(items, shell.MenuItem{
				Label: bgInfo.Name,
				Desc:  "",
				Data:  background.Background(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, val background.Background) tea.Cmd {
		c.Background = val
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, bg background.Background, contentWidth int) string {
		var abilitiesArr []string
		for _, a := range background.Get(bg).Abilities {
			abilityAbbr := ability.Get(a).Abbr
			abilitiesArr = append(abilitiesArr, style.Badge.Render(abilityAbbr))
		}
		abilitiesStr := strings.Join(abilitiesArr, " ")
		bgi := background.Get(bg)

		title := style.Title.Render(bgi.Name)
		desc := style.Desc.Width(contentWidth).Render(bgi.Desc)

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
			style.TypesetDivider(divWidth),
			abilities,
			feat,
			skills,
			tools,
			style.Header.Render("Equipment Options"),
			equip,
		)
	},
}
