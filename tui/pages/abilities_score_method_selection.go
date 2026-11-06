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
