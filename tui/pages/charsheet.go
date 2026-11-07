// Package pages/charsheet.go
package pages

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/alignment"
	"dndwiz/character/background"
	"dndwiz/character/class"
	"dndwiz/character/species"
	draconic "dndwiz/character/species/draconic_ancestry"
	elven "dndwiz/character/species/elven_lineage"
	"dndwiz/tui/modal"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var CharsheetConfig = shell.ShellConfig[pager.CurrentView]{
	ContentTitle: "Character Sheet",
	MenuTitle:    "Edit Character",
	Layout:       shell.LayoutSidebar,
	ShowDesc:     true,
	EscDesc:      "options",
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
			style.Header.Render("Welcome to the character sheet"),
			p("A character sheet is a summary of your adventurer: who they are, what they are good "+
				"at, and what they can do in the game. This screen is your home base. Everything "+
				"you choose is shown on the right, and the menu on the left lets you change any part "+
				"of it."),
			"",
			p("You do not need to know the game to use this. Each screen you open has its own "+
				"help page, which you can open by pressing ? on it."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down the menu. Press Enter or Space on an item to edit it. "+
				"When you finish, you return here and the sheet updates to show your choices."),
			"",
			p("Press Esc to open the pause menu, where you can resume, or quit to the title "+
				"screen or the program. Save and Export are listed but not available yet, so "+
				"unsaved characters are lost when you quit."),
			"",
			p("If the sheet is longer than the screen, scroll it with Shift+↑/↓."),
			style.TypesetDivider(w),

			style.Header.Render("The menu items"),
			term("Name",
				"Opens a box where you type your character's name. Press Enter to keep it, or "+
					"Esc to cancel."),
			term("Class",
				"What your character does: fight, cast spells, sneak, or heal. This is the "+
					"biggest choice, and it sets your primary ability. Fighters are asked to choose "+
					"between Strength and Dexterity."),
			term("Species",
				"What kind of person your character is: human, elf, dwarf, and so on. It sets "+
					"your size and speed. Some species ask a follow-up question about size, "+
					"ancestry, or lineage."),
			term("Background",
				"What your character did before adventuring. It gives you a feat, a tool, "+
					"and three abilities."),
			term("Ability Scores",
				"Six numbers that describe your strengths and weaknesses. You can roll for them, "+
					"spend points to buy them, or take a standard set suited to your class. The "+
					"Standard Array option appears once you have chosen a class."),
			term("Alignment",
				"A short description of your character's moral outlook. It is a guide for "+
					"roleplaying and does not change your abilities."),
			style.TypesetDivider(w),

			style.Header.Render("Reading the sheet"),
			term("Tagline",
				"The line under your name that sums you up, such as your background, species, "+
					"and class."),
			term("Ability scores",
				"The six boxes (STR, DEX, CON, INT, WIS, CHA). The large number is the score and "+
					"the smaller number below it is its modifier, the bonus or penalty it gives to "+
					"your rolls."),
			term("Primary Abilities",
				"The scores that matter most for your class. Put your best numbers here."),
			term("Species Attributes",
				"Your creature type, size, and speed (how many feet you can move on your turn)."),
			term("Background Highlights",
				"Your background's feat, tool, and abilities, along with your alignment."),
			p("Sections appear as you make choices. A new character starts with a short welcome "+
				"message instead of a full sheet."),
			style.TypesetDivider(w),

			style.Header.Render("A suggested order"),
			p("There is no required order, and you can change anything at any time. If you are "+
				"new, this order works well:"),
			"",
			p("1. Class: it affects the most. If you are unsure, read its help page first."),
			p("2. Species: pick a story you like."),
			p("3. Background: this adds a little history."),
			p("4. Ability Scores: choosing a class first unlocks the Standard Array, an easy "+
				"option for beginners."),
			p("5. Alignment and Name: these are about personality, so do them last."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure where to start?"),
			p("Open Class, read about two or three, and pick the one that sounds like fun. "+
				"Everything else can follow from that, and you can always come back and change "+
				"your mind."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		return []shell.MenuItem{
			{
				Label: "Name",
				Desc:  c.Name,
				Data:  pager.NameInput,
			},
			{
				Label: "Class",
				Desc:  class.Get(c.Class).Name,
				Data:  pager.ClassSelection,
			},
			{
				Label: "Species",
				Desc:  species.Get(c.Species).Name,
				Data:  pager.SpeciesSelection,
			},
			{
				Label: "Background",
				Desc:  background.Get(c.Background).Name,
				Data:  pager.BackgroundSelection,
			},
			{
				Label: "Ability Scores",
				Desc:  c.AbilityScoreMethod.String(),
				Data:  pager.AbilityScoresMethodSelection,
			},
			{
				Label: "Alignment",
				Desc:  alignment.Get(c.Alignment).Name,
				Data:  pager.AlignmentSelection,
			},
		}
	},

	OnSelect: func(c *character.Character, selected pager.CurrentView) tea.Cmd {
		if selected == pager.AbilityScoresMethodSelection && c.Class == class.None {
			return modal.Choice(modal.ChoiceOptions{
				Title: "No class yet",
				Body: "You haven't chosen a class, so the Standard Array won't be offered. " +
					"It depends on your class. Pick a class first to unlock it.",
				Options: []modal.ChoiceOption{
					{Label: "Pick a class", Cmd: pager.SwitchTo(pager.ClassSelection)},
					{Label: "Continue", Cmd: pager.SwitchTo(pager.AbilityScoresMethodSelection)},
					{Label: "Back"},
				},
			})
		}

		return func() tea.Msg {
			return pager.SwitchViewMsg(selected)
		}
	},

	Content: func(c *character.Character, selected pager.CurrentView, contentWidth int) string {
		return displayCharacterOverview(c, contentWidth)
	},
}

// --- Helper Formatting Functions ---

func headerDisplay(c *character.Character) string {
	return style.Title.Render(c.Name)
}

func taglineDisplay(c *character.Character) string {
	var species_ string
	switch c.Species {
	case species.None:
		// nothing chosen yet
	case species.Dragonborn:
		if c.DraconicAncestry != draconic.None {
			species_ = draconic.Get(c.DraconicAncestry).Name + " "
		}
		species_ += species.Get(c.Species).Name
	case species.Elf:
		if c.ElvenLineage != elven.None {
			species_ = elven.Get(c.ElvenLineage).Name
		} else {
			species_ = species.Get(c.Species).Name
		}
	default:
		species_ = species.Get(c.Species).Name
	}

	var parts []string
	if c.Background != background.None {
		parts = append(parts, background.Get(c.Background).Name)
	}

	// Species and class read as one phrase ("Gold Dragonborn Wizard").
	var identity []string
	if species_ != "" {
		identity = append(identity, species_)
	}
	if c.Class != class.None {
		identity = append(identity, class.Get(c.Class).Name)
	}
	if len(identity) > 0 {
		parts = append(parts, strings.Join(identity, " "))
	}

	return style.Tagline.Render(strings.Join(parts, " • "))
}

func hasAbilityScores(c *character.Character) bool {
	for _, score := range c.AbilityScore {
		if score > 8 {
			return true
		}
	}
	return false
}

func abilityScoresDisplay(c *character.Character) string {
	if !hasAbilityScores(c) {
		return ""
	}
	return style.TypesetScores(c, ability.None, false)
}

func primaryAbilitiesDisplay(c *character.Character) string {
	if len(c.PrimaryAbilities) == 0 {
		return ""
	}
	return style.Label.Render("Primary Abilities: ") + style.Badge.Render(c.PrimaryAbilities.String())
}

func speciesAttributesDisplay(c *character.Character) string {
	if c.Species == species.None {
		return ""
	}

	spec := species.Get(c.Species)

	header := style.Header.Render("Species Attributes")
	info := fmt.Sprintf(
		"%s %s\n%s %s\n%s %s",
		style.Label.Render("Type:"), spec.CreatureType,
		style.Label.Render("Size:"), c.Size,
		style.Label.Render("Speed:"), spec.Speed.String(),
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, info)
}

func backgroundDisplay(c *character.Character) string {
	hasBg := c.Background > background.None
	hasAln := c.Alignment > alignment.None

	if !hasBg && !hasAln {
		return ""
	}

	header := style.Header.Render("Background Highlights")

	var bgStr, alnStr string

	if hasBg {
		bg := background.Get(c.Background)
		var bgAbilities []string
		for _, a := range bg.Abilities {
			bgAbilities = append(bgAbilities, style.Badge.Render(ability.Get(a).Abbr))
		}

		bgAbilitiesStr := strings.Join(bgAbilities, " ")

		bgStr = fmt.Sprintf(
			"%s %s\n%s %s\n%s %s",
			style.Label.Render("Feat:"), bg.Feat.String(),
			style.Label.Render("Tool:"), bg.Tool,
			style.Label.Render("Abilities:"), bgAbilitiesStr,
		)
	}

	if hasAln {
		aln := alignment.Get(c.Alignment)
		alnStr = style.Label.Render("Alignment:") + " " + aln.Name
	}

	if hasBg && !hasAln {
		return lipgloss.JoinVertical(lipgloss.Left, header, bgStr)
	}
	if !hasBg && hasAln {
		return lipgloss.JoinVertical(lipgloss.Left, header, alnStr)
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, bgStr, alnStr)
}

func displayCharacterOverview(c *character.Character, w int) string {
	blankCharacter := c.Class == class.None &&
		c.Species == species.None &&
		c.Background == background.None &&
		c.Alignment == alignment.None &&
		!hasAbilityScores(c)

	if blankCharacter {
		flavor := "The tavern fire crackles, the ink is still wet, and somewhere out in the dark " +
			"a story is waiting for someone brave, clever, or foolish enough to step into it.\n\n" +
			"That someone has no name yet. No past, no calling, no scars. " +
			"They are whatever you decide they will be."

		nudge := "Choose an item from the menu to begin. Start with Class if you like, or follow " +
			"whatever sparks your imagination. Nothing is final, and every choice can be changed " +
			"later. Press ? on any screen if you want a guide."

		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerDisplay(c),
			style.Desc.Render(flavor),
			style.TypesetDivider(w),
			style.Body.Render(nudge),
		)
	} else {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerDisplay(c),
			taglineDisplay(c),
			abilityScoresDisplay(c),
			primaryAbilitiesDisplay(c),
			speciesAttributesDisplay(c),
			backgroundDisplay(c),
		)
	}
}
