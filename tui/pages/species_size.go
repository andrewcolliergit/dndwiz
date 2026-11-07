// Package pages/species_size.go
package pages

import (
	"dndwiz/character"
	"dndwiz/character/species"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var SpeciesSizeConfig = shell.ShellConfig[string]{
	ContentTitle: "Species Size",
	MenuTitle:    "Select Species Size",
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
			style.Header.Render("What is size?"),
			p("Size describes how big your character is. Most species have only one size, so "+
				"the game picks it for you. Humans and Tieflings can be either Medium or Small, "+
				"which is why you are being asked."),
			"",
			p("This choice is mostly about how you imagine your character. It does not change "+
				"your class, your abilities, or your speed."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up or down to look at each size. Nothing is chosen until you "+
				"press Enter or Space. After you choose, you return to the character sheet, where "+
				"your size appears under Species Attributes. You can come back and change it "+
				"at any time."),
			"",
			p("If you press Esc instead, you go back to the species list, and your character "+
				"keeps the first size listed, which is Medium."),
			style.TypesetDivider(w),

			style.Header.Render("The sizes"),
			term("Medium",
				"Roughly human-sized, about 4 to 7 feet tall. This is the most common size, and "+
					"it has no special advantages or drawbacks. Choose it if you want a "+
					"typical adventurer."),
			term("Small",
				"Shorter, about 2 to 4 feet tall. Small characters are better at fitting into "+
					"tight spaces, but they have a drawback: weapons with the Heavy property, such "+
					"as a greatsword, greataxe, or longbow, are awkward for them and are harder "+
					"to hit with. Choose it if you like the look of a smaller hero, or if your "+
					"class uses lighter weapons."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see"),
			term("Heavy",
				"A weapon property. Heavy weapons are big and hard to swing, so Small creatures "+
					"roll with a penalty when attacking with them. The weapon's description says "+
					"whether it is Heavy."),
			term("Speed",
				"How many feet you can move on your turn. Size does not change it."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("If you plan to use a heavy weapon, such as a Fighter with a greatsword, choose "+
				"Medium. Otherwise the difference is small, so pick whichever fits the "+
				"character in your head. Neither is stronger overall."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		sizes := species.Get(c.Species).Size
		items := make([]shell.MenuItem, 0, len(sizes))

		for _, s := range sizes {
			items = append(items, shell.MenuItem{
				Label: s,
				Desc:  "",
				Data:  s,
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selectedSize string) tea.Cmd {
		c.Size = selectedSize
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, selectedSize string, contentWidth int) string {
		return style.Title.Render(selectedSize)
	},
}
