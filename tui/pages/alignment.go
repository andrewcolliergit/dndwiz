// Package pages/alignment
package pages

import (
	"dndwiz/character"
	"dndwiz/character/alignment"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AlignmentSelectionConfig = shell.ShellConfig[alignment.Alignment]{
	ContentTitle: "Alignments",
	MenuTitle:    "Select Alignment",
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
			style.Header.Render("What is an alignment?"),
			p("Alignment is a short description of your character's moral outlook and how they "+
				"behave. It is a guide for roleplaying, not a rule that controls what you can do. "+
				"You can use it to figure out how your character would react to a tough choice."),
			"",
			p("It does not change your abilities or your class. Your character can act against their "+
				"alignment now and then, and many groups let it shift over the course of a story."),
			style.TypesetDivider(w),

			style.Header.Render("How to use this screen"),
			p("Move the cursor up and down to read about each alignment. Nothing is chosen until "+
				"you press Enter or Space. After you choose, you return to the character sheet. You "+
				"can come back and change your alignment at any time."),
			style.TypesetDivider(w),

			style.Header.Render("The two scales"),
			p("Every alignment combines two ideas."),
			"",
			term("Good, Neutral, Evil",
				"How much your character cares about other people. Good characters try to help and "+
					"protect others. Evil characters are willing to harm others for their own gain. "+
					"Neutral falls in between, doing what seems reasonable."),
			term("Lawful, Neutral, Chaotic",
				"How your character feels about rules and order. Lawful characters respect laws, "+
					"tradition, and promises. Chaotic characters value freedom and follow their own "+
					"instincts. Neutral falls in between."),
			p("Combining the two scales gives nine alignments. A character who is both neutral "+
				"on each scale is just called Neutral."),
			style.TypesetDivider(w),

			style.Header.Render("The alignments"),
			term("Lawful Good (LG)",
				"Does the right thing as society expects. A knight who protects the innocent and "+
					"keeps their word."),
			term("Neutral Good (NG)",
				"Does the best they can to help others, working within rules but not bound by them. "+
					"A kind healer who helps whoever needs it."),
			term("Chaotic Good (CG)",
				"Follows their conscience and has little regard for what others expect. A rebel who "+
					"robs a cruel tax collector to feed the poor."),
			term("Lawful Neutral (LN)",
				"Follows law, tradition, or a personal code above all. A judge or soldier who "+
					"follows orders and a strict discipline."),
			term("Neutral (N)",
				"Avoids taking sides and does whatever seems best at the moment. Someone who "+
					"would rather live their life than debate right and wrong."),
			term("Chaotic Neutral (CN)",
				"Follows their whims and values personal freedom above everything. A wanderer who "+
					"lives by their wits."),
			term("Lawful Evil (LE)",
				"Takes what they want through order, tradition, or loyalty. A noble who exploits "+
					"the people while scheming for more power."),
			term("Neutral Evil (NE)",
				"Is untroubled by the harm they cause while chasing their own desires. A criminal "+
					"who robs and kills as they please."),
			term("Chaotic Evil (CE)",
				"Acts with arbitrary violence, driven by hatred or bloodlust. A villain who spreads "+
					"havoc. These are usually villains, so check with your group before choosing it."),
			style.TypesetDivider(w),

			style.Header.Render("Not sure what to pick?"),
			p("Ask yourself two questions. Does my character care more about helping others or "+
				"themselves? Do they prefer rules or freedom? If you are unsure, Neutral Good is a "+
				"safe, flexible choice that fits almost any group, and you can change it later."),
			"",
			style.Desc.Render("Press Esc or ? to close."),
		)
	},
	GetItems: func(c *character.Character) []shell.MenuItem {
		alignments := alignment.GetAll()
		items := make([]shell.MenuItem, 0, len(alignments))

		for i, alnInfo := range alignments {
			items = append(items, shell.MenuItem{
				Label: alnInfo.Name,
				Desc:  alnInfo.Desc,
				Data:  alignment.Alignment(i + 1),
			})
		}
		return items
	},

	OnSelect: func(c *character.Character, val alignment.Alignment) tea.Cmd {
		c.Alignment = val
		return func() tea.Msg {
			return pager.SwitchViewMsg(pager.Charsheet)
		}
	},

	Content: func(c *character.Character, a alignment.Alignment, contentWidth int) string {
		aln := alignment.Get(a)
		title := style.Title.Render(aln.Name)
		desc := style.Desc.Render(aln.Desc)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			desc,
		)
	},
}
