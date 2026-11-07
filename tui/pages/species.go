// Package pages/species.go
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

var SpeciesSelectionConfig = shell.ShellConfig[species.Species]{
	ContentTitle: "Species",
	MenuTitle:    "Select Species",
	Layout:       shell.LayoutSidebar,
	MenuWidth:    32,
	ShowDesc:     false,
	HelpContent: func(w int) string {
		p := func(s string) string { return style.TypesetBody(s, w) }
		sp := func(name, traits string) string {
			return lipgloss.JoinVertical(lipgloss.Left,
				style.Label.Render(name),
				p(traits),
				"",
			)
		}

		return lipgloss.JoinVertical(lipgloss.Left,
			style.Header.Render("What is a species?"),
			p("Your species is what kind of person your character is: human, elf, dwarf, and so on. "+
				"In older editions of the game this was called a race. It sets a few basic facts about your "+
				"body (how big you are and how fast you move) and gives you special traits, which are small "+
				"abilities that every member of that species has."),
			"",
			p("Species is about where your character comes from, not what they do for a living. "+
				"That comes from your class. Any species can be any class, so pick the one you find "+
				"most interesting, not the one you think is strongest."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each species. Nothing is chosen until you press "+
				"Enter or Space, so browse as much as you like. You can come back and change your "+
				"species at any time."),
			"",
			p("Some species ask a follow-up question after you select them. This app currently asks about "+
				"size (Human and Tiefling), draconic ancestry (Dragonborn), and elven lineage (Elf). "+
				"You will return to the character sheet when you finish."),
			style.TypesetDivider(w),

			style.Header.Render("What the details mean"),
			style.Label.Render("Creature Type"),
			p("Almost every playable species is Humanoid. This matters mostly for spells and effects "+
				"that only work on certain kinds of creatures."),
			"",
			style.Label.Render("Size"),
			p("Medium characters are roughly human-sized. Small characters are shorter, which helps "+
				"with squeezing into tight spaces, but some heavy weapons are harder for them to use. "+
				"Species with two size options let you choose."),
			"",
			style.Label.Render("Speed"),
			p("How many feet you can move on your turn. Thirty feet is typical. Thirty-five is slightly faster."),
			"",
			style.Label.Render("Traits"),
			p("The special abilities your species gives you. The list on the right shows only their names; "+
				"the glossary below explains what they do."),
			style.TypesetDivider(w),

			style.Header.Render("Words you will see in traits"),
			style.Label.Render("Darkvision"),
			p("You can see in dim light and darkness for a limited distance, though only in shades of gray."),
			"",
			style.Label.Render("Resistance"),
			p("You take half damage from one type of harm, such as fire or poison."),
			"",
			style.Label.Render("Advantage"),
			p("When you roll a d20, you roll two and keep the higher result."),
			"",
			style.Label.Render("Cantrip"),
			p("A simple spell you can cast as often as you like."),
			"",
			style.Label.Render("Long Rest"),
			p("About eight hours of sleep and light activity. Many abilities recharge after one."),
			style.TypesetDivider(w),

			style.Header.Render("The species"),
			sp("Dragonborn",
				"Descended from dragons. You choose a draconic ancestry, which decides the damage type of your "+
					"Breath Weapon (an attack you exhale in a cone) and gives you resistance to that same damage. "+
					"You also have Darkvision, and at 5th level you can sprout temporary wings and fly. "+
					"A strong pick if you like a clear identity and a ranged attack."),
			sp("Dwarf",
				"Sturdy and hard to kill. You have Darkvision, resistance to poison damage, advantage on "+
					"saving throws against being poisoned, extra hit points at every level, and a knack for "+
					"sensing stonework. A good, simple choice for a first character who wants to be tough."),
			sp("Elf",
				"Graceful and long-lived. You have Darkvision, advantage against being charmed, and training "+
					"in a sense-based skill, and you meditate for four hours instead of sleeping eight. You also "+
					"choose a lineage: Drow, High Elf, or Wood Elf. Each gives you a cantrip or speed bonus now, "+
					"plus spells at levels 3 and 5."),
			sp("Gnome",
				"Small, curious, and magical. You have Darkvision and advantage on Intelligence, Wisdom, and "+
					"Charisma saving throws, which protects you from many mind-affecting spells. Gnomish Lineage "+
					"adds a bit of magic; this app does not ask you to choose one yet."),
			sp("Goliath",
				"Tall and strong, descended from giants. You are a bit faster than most (35 feet), you are "+
					"hard to knock down or restrain thanks to Powerful Build, and at 5th level you can briefly "+
					"grow to Large size. Giant Ancestry gives you a supernatural boon, which this app does not "+
					"ask you to choose yet."),
			sp("Halfling",
				"Small, brave, and lucky. When you roll a 1 on a d20 you get to roll again, you have "+
					"advantage against being frightened, you can slip past larger creatures, and you are "+
					"naturally good at hiding. Great for a character who survives on good fortune."),
			sp("Human",
				"The most adaptable species. You choose Medium or Small size, gain one skill of your choice, "+
					"get an Origin feat of your choice, and earn Heroic Inspiration after each long rest, "+
					"which lets you reroll one die. A flexible choice that fits any class."),
			sp("Orc",
				"Powerful and relentless. You have Darkvision, you can dash as a bonus action and gain "+
					"temporary hit points when you do (Adrenaline Rush), and once per long rest you can drop "+
					"to 1 hit point instead of falling unconscious. Excellent for frontline characters."),
			sp("Tiefling",
				"Touched by the Lower Planes. You have Darkvision and the Thaumaturgy cantrip, and your "+
					"Fiendish Legacy grants resistance to a damage type and spells that grow as you level. "+
					"You choose Medium or Small size."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Choose the one whose story you enjoy most. Species traits are small bonuses, and any "+
				"species can succeed in any class. If you want something easy to play, Human, Dwarf, "+
				"and Halfling have traits that work without much bookkeeping."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		allSpecies := species.GetAll()
		items := make([]shell.MenuItem, 0, len(allSpecies))

		for i, specInfo := range allSpecies {
			items = append(items, shell.MenuItem{
				Label: specInfo.Name,
				Desc:  specInfo.Traits,
				Data:  species.Species(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, selected species.Species) tea.Cmd {
		c.Species = selected
		spec := species.Get(selected)
		c.Size = spec.Size[0]

		if spec.Size.RequiresChoice() {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesSizeSelection)
			}
		} else if c.Species == species.Dragonborn {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesAncestrySelection)
			}
		} else if c.Species == species.Elf {
			return func() tea.Msg {
				return pager.SwitchViewMsg(pager.SpeciesLineageSelection)
			}
		}

		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, selected species.Species, contentWidth int) string {
		info := species.Get(selected)

		// Header
		title := style.Title.Render(info.Name)

		// Stat Grid / Key-Value Blocks
		creatureType := style.Label.Render("Creature Type: ") + info.CreatureType
		size := style.Label.Render("Size: ") + info.Size.GetAll()
		speed := style.Label.Render("Speed: ") + style.Badge.Render(info.Speed.String())

		// Combine stats into a neat vertical block
		statsGroup := lipgloss.JoinVertical(
			lipgloss.Left,
			creatureType,
			size,
			speed,
		)

		// Special traits section
		traits := lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("248")).
			Render(info.Traits)

		// Main flavor description
		desc := style.Desc.Render(info.Desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
			style.TypesetDivider(contentWidth-2),
			statsGroup,
			style.Header.Render("Traits"),
			traits,
		)
	},
}
