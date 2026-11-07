// Package pages/abilities_score_method_selection.go
package pages

import (
	"dndwiz/character"
	"dndwiz/character/class"
	"dndwiz/character/score"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AbilityScoreMethodConfig = shell.ShellConfig[pager.CurrentView]{
	ContentTitle: "Ability Scores Selection Method",
	MenuTitle:    "Allocation Method",
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
			style.Header.Render("What are ability scores?"),
			p("Ability scores are six numbers that describe your character's natural talents: "+
				"how strong, quick, tough, smart, perceptive, and charming they are. Almost "+
				"everything your character attempts in the game, from swinging a sword to "+
				"persuading a guard, is based on one of these six."),
			"",
			p("This screen is where you choose how to generate them. The three methods give "+
				"you the same kind of result, a set of six numbers, but they feel different to use."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each method. Nothing happens until you "+
				"press Enter or Space. Each method then opens its own screen, except Standard "+
				"Array, which applies immediately and returns you to the character sheet."),
			"",
			p("You can come back and switch methods at any time. Switching replaces the scores "+
				"you currently have."),
			style.TypesetDivider(w),

			style.Header.Render("The six abilities"),
			term("Strength (STR)",
				"Physical power. Used for lifting, climbing, and hitting hard with heavy weapons."),
			term("Dexterity (DEX)",
				"Reflexes, balance, and aim. Used for dodging, sneaking, and shooting."),
			term("Constitution (CON)",
				"Health and stamina. Gives you more hit points, so everyone benefits from it."),
			term("Intelligence (INT)",
				"Reasoning and memory. Used for knowledge, investigation, and Wizard spells."),
			term("Wisdom (WIS)",
				"Awareness and intuition. Used for noticing things, reading people, and Cleric "+
					"and Druid spells."),
			term("Charisma (CHA)",
				"Force of personality. Used for persuading, performing, and Bard, Sorcerer, and "+
					"Warlock spells."),
			style.TypesetDivider(w),

			style.Header.Render("Scores and modifiers"),
			p("Each score has a modifier, the bonus or penalty it adds to your rolls. A score of "+
				"10 or 11 gives +0. Every two points above that adds +1, and every two points "+
				"below subtracts 1. For example, 15 gives +2 and 8 gives -1."),
			"",
			p("Starting scores usually fall between 8 and 15. Higher is better, but it is more "+
				"useful to be great at what your class needs than to be average at everything."),
			style.TypesetDivider(w),

			style.Header.Render("The three methods"),
			term("Point Buy",
				"You get 27 points to spend. Every score starts at 8, and raising a score costs "+
					"points: the higher it goes, the more each step costs. Scores cannot go above 15. "+
					"You have full control and no luck is involved, so it is the most predictable "+
					"method. Best if you already know which abilities you want high."),
			term("Standard Array",
				"A ready-made set of six scores (15, 14, 13, 12, 10, 8) arranged for your class, "+
					"so the best numbers land in the abilities your class cares about. It is instant "+
					"and balanced, which makes it the easiest choice for a first character. This "+
					"option appears only after you choose a class, because the arrangement depends "+
					"on it."),
			term("Roll",
				"You roll dice to generate the scores. For each score you roll four six-sided dice "+
					"and add the highest three, then you choose which ability gets each result. "+
					"Scores can come out higher or lower than the other methods, so it adds an "+
					"element of luck. You get only one set of rolls, and they cannot be rerolled."),
			style.TypesetDivider(w),

			style.Header.Render("Which should I pick?"),
			p("If you are new, take the Standard Array. It is quick and gives your class a solid "+
				"start. If you like planning and want control, try Point Buy. If you enjoy "+
				"surprises and your table uses dice, try Roll."),
			"",
			p("You can try the methods in any order. Switching replaces your scores, but Roll "+
				"keeps your original six rolls, so you can return to them."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		items := []shell.MenuItem{
			{
				Label: "Point Buy",
				Desc:  "",
				Data:  pager.AbilityScoresBuy,
			},
		}

		if c.Class != class.None {
			stdarray := class.Get(c.Class).StdArray

			items = append(items, shell.MenuItem{
				Label: "Standard Array",
				Desc:  stdarray.String(),
				Data:  pager.Charsheet,
			})
		}

		// if !c.HasRolledAbilityScores {
		items = append(items, shell.MenuItem{
			Label: "Roll",
			Desc:  "",
			Data:  pager.AbilityScoresRoll,
		})
		//}

		return items
	},

	OnSelect: func(c *character.Character, val pager.CurrentView) tea.Cmd {
		switch val {
		case pager.Charsheet:
			cls := class.Get(c.Class)
			c.AbilityScore = cls.StdArray
			c.AbilityScoreMethod = score.StdArray

		case pager.AbilityScoresRoll:
			prev := c.AbilityScoreMethod
			c.AbilityScoreMethod = score.Roll
			if prev == score.StdArray || prev == score.PointBuy {
				// the other method overwrote c.AbilityScore, so reset the cached roll model
				return tea.Sequence(pager.SwitchTo(val), pager.ResetView())
			}

		case pager.AbilityScoresBuy:
			c.AbilityScoreMethod = score.PointBuy
		}
		return func() tea.Msg {
			return pager.SwitchViewMsg(val)
		}
	},

	Content: func(c *character.Character, view pager.CurrentView, contentWidth int) string {
		var title, desc string

		switch view {

		case pager.AbilityScoresRoll:
			title = "Roll"
			desc = "Roll six random scores, then choose which ability gets each one. " +
				"Best for players who enjoy a bit of chance and want a character shaped by the dice.\n\n" +
				"Your scores may come out higher or lower than the standard array. " +
				"You only get one set of rolls, so you can't roll again. " +
				"You can still switch to Point Buy or Standard Array later, " +
				"but that replaces the scores you assigned. " +
				"If you switch to Point Buy or Standard Array later, " +
				"you can switch back to your rolled scores again if you change your mind, and will be able to reassign them at any time."

		case pager.AbilityScoresBuy:
			title = "Point Buy"
			desc = "You may elect to customize your skill points as you see fit using Point Buy. With this method, you will be given a budget of 27 points to spend on skill point levels, where each level purchased becomes increasingly expensive."

		// Standard Array
		case pager.Charsheet:
			tmpCharacter := character.NewCharacter()
			tmpCharacter.AbilityScore = class.Get(c.Class).StdArray
			box := style.TypesetScores(&tmpCharacter, -1, false)
			title = "Standard Array"
			desc = "Rather than customizing your ability scores, you may elect to take the standard score array for your chosen class. The standard skill point array for the " + class.Get(c.Class).Name + " class is:\n\n" + box
		}

		// Header & Accents
		title = style.Title.Render(title)

		// Body content
		desc = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Render(desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
		)
	},
}
