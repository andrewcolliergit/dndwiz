// Package pages/draconic_ancestry.go
package pages

import (
	"dndwiz/character"
	draconic "dndwiz/character/species/draconic_ancestry"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var DraconicAncestryConfig = shell.ShellConfig[draconic.Ancestry]{
	ContentTitle: "Draconic Ancestry",
	MenuTitle:    "Select Draconic Ancestry",
	Layout:       shell.LayoutSidebar,
	MenuWidth:    32,
	ShowDesc:     true,
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
			style.Header.Render("What is a draconic ancestry?"),
			p("Dragonborn are descended from dragons. Your ancestry is the kind of dragon your "+
				"character descends from, such as red, silver, or bronze. It decides two things: the "+
				"color of your scales, and the type of damage your breath and your natural "+
				"protection are tied to."),
			"",
			p("It does not affect your class, your personality, or anything else about your "+
				"character, so choose the dragon whose story or color you like best."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each ancestry. The list on the left shows "+
				"each dragon's damage type. Nothing is chosen until you press Enter or Space. After "+
				"you choose, you return to the character sheet, where your ancestry appears in the "+
				"tagline (for example, \"Gold Dragonborn\"). You can come back and change it at any time."),
			style.TypesetDivider(w),

			style.Header.Render("What your ancestry does"),
			term("Breath Weapon",
				"You can exhale destructive energy in a 15-foot cone. Creatures in the cone make a "+
					"Dexterity saving throw and take damage of your ancestry's type, or half as much "+
					"if they succeed. The damage starts small and grows as you level up. You can use "+
					"it a limited number of times, and the uses come back after a long rest."),
			term("Damage Resistance",
				"You take half damage from the same damage type as your breath. A gold dragonborn, "+
					"for example, breathes fire and resists fire."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see"),
			term("Damage Type",
				"Every attack or spell deals a kind of damage, such as fire or cold. Some monsters "+
					"are immune or resistant to certain types, so it matters which one you deal."),
			term("Saving Throw",
				"A roll made to resist an effect. Your breath forces enemies to make one."),
			term("Long Rest",
				"About eight hours of sleep and light activity. Many abilities recharge after one."),
			style.TypesetDivider(w),

			style.Header.Render("Chromatic and metallic dragons"),
			p("Dragons come in two families. Chromatic dragons (black, blue, green, red, white) "+
				"are traditionally feared and destructive. Metallic dragons (brass, bronze, copper, "+
				"gold, silver) are traditionally noble and protective. Your ancestry's family is "+
				"only flavor. A descendant of a chromatic dragon can be a hero, and the reverse."),
			style.TypesetDivider(w),

			style.Header.Render("The ancestries by damage type"),
			term("Acid", "Black and Copper."),
			term("Cold", "Silver and White."),
			term("Fire", "Brass, Gold, and Red."),
			term("Lightning", "Blue and Bronze."),
			term("Poison", "Green."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Pick the damage type you would most enjoy using, or the color you like most. Fire "+
				"and cold are the most common, and poison is the one many monsters resist, so it is "+
				"the least flexible choice. For a first character, the difference is small."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		ancestries := draconic.GetAll()
		items := make([]shell.MenuItem, 0, len(ancestries))

		for i, a := range ancestries {
			items = append(items, shell.MenuItem{
				Label: a.Name,
				Desc:  a.DamageType,
				Data:  draconic.Ancestry(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedDraconicAncestry draconic.Ancestry) tea.Cmd {
		c.DraconicAncestry = selectedDraconicAncestry
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, d draconic.Ancestry, contentWidth int) string {
		dai := draconic.Get(d)

		title := style.Title.Render(dai.Name)
		dmgType := style.Label.Render("Damage Type: ") + style.Badge.Render(dai.DamageType)
		desc := style.Desc.Width(contentWidth).Render(dai.Desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
			style.TypesetDivider(contentWidth-2),
			dmgType,
		)
	},
}
