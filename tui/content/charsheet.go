// Package content
package content

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
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var CharsheetConfig = shell.ShellConfig[pager.CurrentView]{
	Title:     "Character Sheet",
	MenuTitle: "Edit Character",
	Layout:    shell.LayoutSidebar,
	Width:     32,
	ShowDesc:  true,

	GetItems: func(c *character.Character) shell.MenuItemList {
		return shell.MenuItemList{
			{
				Label: "Name",
				Desc:  c.Name,
				Data:  pager.NameTextbox,
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
		return func() tea.Msg {
			return pager.SwitchViewMsg(selected)
		}
	},

	Content: func(c *character.Character, selected pager.CurrentView, width int) string {
		blankCharacter := c.Class == class.None &&
			c.Species == species.None &&
			c.Background == background.None &&
			c.Alignment == alignment.None &&
			!hasAbilityScores(c)

		if blankCharacter {
			return blankCharacterGreeter()
		}

		return displayCharacterOverview(c)
	},
}

// --- Helper Formatting Functions ---

func headerDisplay(c *character.Character) string {
	return style.Title.Render(c.Name)
}

func taglineDisplay(c *character.Character) string {
	cls := class.Get(c.Class)
	spec := species.Get(c.Species)
	bg := background.Get(c.Background)

	var s string
	switch c.Species {
	case species.Dragonborn:
		s = draconic.Get(c.DraconicAncestry).Name + " " + spec.Name
	case species.Elf:
		s = elven.Get(c.ElvenLineage).Name
	default:
		s = spec.Name
	}

	var output string
	if c.Background != background.None {
		output = bg.Name
		if c.Species != species.None && c.Class != class.None {
			output += " • "
		}
	}
	if c.Species != species.None {
		output += s + " "
	}
	if c.Class != class.None {
		output += cls.Name
	}

	return style.Tagline.Render(output)
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
	return style.RenderAbilityScoresRow(c, ability.None, false)
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

func displayCharacterOverview(c *character.Character) string {
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

func blankCharacterGreeter() string {
	header := style.Title.Inherit(style.Title).Align(lipgloss.Center).Render("NEW CHARACTER")
	text := lipgloss.NewStyle().Align(lipgloss.Center).Render("As you fill in the details of your new character, this character sheet will reflect the traits and choices you have selected.")
	return lipgloss.JoinVertical(lipgloss.Center, header, text)
}
