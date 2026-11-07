// Package pages/elven_lineage.go
package pages

import (
	"dndwiz/character"
	elven "dndwiz/character/species/elven_lineage"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ElvenLineageConfig = shell.ShellConfig[elven.Lineage]{
	ContentTitle: "Elven Lineage",
	MenuTitle:    "Select Elven Lineage",
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
			style.Header.Render("What is an elven lineage?"),
			p("Elves come from many different homelands, and over the centuries those homes shaped "+
				"them in different ways. Your lineage is the branch of elvenkind your character "+
				"belongs to. It gives you a bonus that is on top of everything every elf already has, "+
				"so it is the part that makes your elf feel distinct from other elves."),
			"",
			p("A lineage is about heritage, not profession. Any lineage works with any class, so "+
				"pick the one whose story or perks you like best."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each lineage. Nothing is chosen until you "+
				"press Enter or Space. After you choose, you return to the character sheet, where your "+
				"lineage appears in the tagline. You can come back and change it at any time."),
			style.TypesetDivider(w),

			style.Header.Render("What Level 1, 3, and 5 mean"),
			p("Characters grow by gaining levels as they adventure, and every new character starts "+
				"at level 1. Your lineage gives you something now and more later:"),
			"",
			term("Level 1",
				"A benefit you have from the very start of your character's life, such as a cantrip "+
					"or extra speed."),
			term("Level 3",
				"A spell you gain once your character reaches level 3. It is always prepared, so it "+
					"does not count against the spells you choose, and you can cast it once per long "+
					"rest without using a spell slot."),
			term("Level 5",
				"A second spell, gained at level 5, which works the same way as the first."),
			p("The level 3 and level 5 entries on the detail screen show only a spell's name. The "+
				"glossary below says what each one does."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see"),
			term("Darkvision",
				"You can see in dim light and darkness for a limited distance, though only in shades "+
					"of gray. All elves have it, and the Drow have a longer range."),
			term("Cantrip",
				"A simple spell you can cast as often as you like, with no limit and no cost."),
			term("Speed",
				"How many feet you can move on your turn. Thirty feet is typical."),
			term("Long Rest",
				"About eight hours of sleep and light activity. Many abilities recharge after one."),
			term("Spell Slot",
				"A resource that powers most spells. Your lineage spells let you skip spending one "+
					"once per long rest."),
			style.TypesetDivider(w),

			style.Header.Render("The lineages"),
			term("Drow",
				"Descended from elves of the Underdark, the lightless world beneath the surface. "+
					"Your Darkvision reaches farther (120 feet), and you know the Dancing Lights "+
					"cantrip, which creates floating lights. At level 3 you gain Faerie Fire, which "+
					"outlines creatures in light so they are easier to hit, and at level 5 you gain "+
					"Darkness, which fills an area with magical shadow. A good fit for a stealthy or "+
					"mysterious character."),
			term("High Elf",
				"A scholar of the arcane, trained in the refined magic of elven courts. You know "+
					"the Prestidigitation cantrip, which performs small magical tricks, and after each "+
					"long rest you can swap it for any cantrip from the Wizard spell list. At level 3 "+
					"you gain Detect Magic, which senses nearby magic, and at level 5 you gain Misty "+
					"Step, which teleports you a short distance. The most flexible choice, and a "+
					"natural match for spellcasters."),
			term("Wood Elf",
				"A child of ancient forests, swift and at home in the wild. Your speed rises to 35 "+
					"feet, and you know the Druidcraft cantrip, which does small nature effects. At "+
					"level 3 you gain Longstrider, which boosts your speed even more, and at level 5 "+
					"you gain Pass without Trace, which makes you and your allies very hard to notice. "+
					"A good fit for scouts and anyone who wants to move fast."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Choose the story you like best. If you want an easy bonus that works without any "+
				"effort, Wood Elf's extra speed is always on. If you want magic and flexibility, pick "+
				"High Elf. If you like stealth and mystery, pick Drow."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		lineages := elven.GetAll()
		items := make([]shell.MenuItem, 0, len(lineages))

		for i, l := range lineages {
			items = append(items, shell.MenuItem{
				Label: l.Name,
				Desc:  "",
				Data:  elven.Lineage(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedElvenLineage elven.Lineage) tea.Cmd {
		c.ElvenLineage = selectedElvenLineage
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, e elven.Lineage, contentWidth int) string {
		el := elven.Get(e)

		title := style.Title.Render(el.Name)
		desc := style.Desc.Width(contentWidth).Render(el.Desc)

		lvl1 := style.Label.Render("Level 1: ") + el.Lvl1
		lvl3 := style.Label.Render("Level 3: ") + el.Lvl3
		lvl5 := style.Label.Render("Level 5: ") + el.Lvl5

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
			style.TypesetDivider(contentWidth-2),
			lvl1,
			lvl3,
			lvl5,
		)
	},
}
