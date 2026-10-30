// Package charactersheet
package charactersheet

import (
	"fmt"
	"strings"

	"dndwiz/character"
	"dndwiz/tui/menu"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	c                     *character.Character
	title                 string
	menu                  menu.Menu
	textInput             textinput.Model
	isEditingName         bool
	termWidth, termHeight int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func MakeCharsheetSidebarItems(c *character.Character) menu.List {
	return menu.List{
		{
			Header: "Name",
			Desc:   c.Name,
			Data:   pager.ViewNameTextbox,
		},
		{
			Header: "Class",
			Desc:   character.GetClass(c.Class).Name,
			Data:   pager.ViewClassSelection,
		},
		{
			Header: "Species",
			Desc:   character.GetSpecies(c.Species).Name,
			Data:   pager.ViewSpeciesSelection,
		},
		{
			Header: "Background",
			Desc:   character.GetBackground(c.Background).Name,
			Data:   pager.ViewBackgroundSelection,
		},
		{
			Header: "Ability Scores",
			Desc:   c.AbilityScoreMethod.String(),
			Data:   pager.ViewAbilityScoresMethodSelection,
		},
		{
			Header: "Alignment",
			Desc:   character.GetAlignment(c.Alignment).Name,
			Data:   pager.ViewAlignmentSelection,
		},
	}
}

func NewCharsheetModel(c *character.Character) Model {
	ti := textinput.New()
	ti.Placeholder = c.Name
	ti.CharLimit = 32
	ti.Width = 24

	items := MakeCharsheetSidebarItems(c)
	sidebar := menu.New("Edit Character", items, 32, true)

	return Model{
		c:         c,
		title:     "Character Sheet",
		menu:      sidebar,
		textInput: ti,
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// When in text editor
	if m.isEditingName {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if name := m.textInput.Value(); name != "" {
					m.c.Name = name
					m.RefreshSidebar() // Update sidebar description
				}
				m.isEditingName = false
				m.textInput.Blur()
				return m, nil

			case "esc":
				m.isEditingName = false
				m.textInput.Blur()
				return m, nil
			}
		}
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	// Sidebar navigation
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if v, ok := menu.GetItemData[pager.CurrentView](m.menu); ok {
				if v == pager.ViewNameTextbox {
					m.isEditingName = true
					m.textInput.SetValue(m.c.Name)
					m.textInput.Focus()
					return m, textinput.Blink
				}
				return m, pager.SwitchTo(v)
			}
		}
	}
	m.menu, cmd = m.menu.Update(msg, m.termHeight)
	return m, cmd
}

func headerDisplay(c *character.Character) string {
	return style.Title.Render(c.Name)
}

func taglineDisplay(c *character.Character) string {
	cls := character.GetClass(c.Class)
	spec := character.GetSpecies(c.Species)
	bg := character.GetBackground(c.Background)

	var species string

	switch c.Species {
	case character.SpeciesDragonborn:
		species = character.GetDraconicAncestry(c.DraconicAncestry).Name + " " + spec.Name
	case character.SpeciesElf:
		species = character.GetElvenLineage(c.ElvenLineage).Name
	default:
		species = spec.Name
	}

	var output string

	if c.Background != character.BackgroundNone {
		output = bg.Name
		if c.Species != character.SpeciesNone && c.Class != character.ClassNone {
			output += " • "
		}
	}
	if c.Species != character.SpeciesNone {
		output += species + " "
	}
	if c.Class != character.ClassNone {
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
	return style.RenderAbilityScoresRow(c, character.AbilityNone, false)
}

func primaryAbilitiesDisplay(c *character.Character) string {
	if len(c.PrimaryAbilities) == 0 {
		return ""
	}
	return style.Label.Render("Primary Abilities: ") + style.Badge.Render(c.PrimaryAbilities.String())
}

func speciesAttributesDisplay(c *character.Character) string {
	if c.Species == character.SpeciesNone {
		return ""
	}

	spec := character.GetSpecies(c.Species)

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
	hasBg := c.Background > character.BackgroundNone
	hasAln := c.Alignment > character.AlignmentNone

	if !hasBg && !hasAln {
		return ""
	}

	header := style.Header.Render("Background Highlights")

	var (
		bgStr  string
		alnStr string
	)

	if hasBg {
		bg := character.GetBackground(c.Background)
		var bgAbilities []string
		for _, ability := range character.GetBackground(c.Background).Abilities {
			bgAbilities = append(bgAbilities, style.Badge.Render(character.GetAbility(ability).Abbr))
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
		aln := character.GetAlignment(c.Alignment)
		alnStr = style.Label.Render("Alignment:") + " " + aln.Name
	}

	if hasBg && !hasAln {
		return lipgloss.JoinVertical(lipgloss.Left, header, bgStr)
	}
	if !hasBg && hasAln {
		return lipgloss.JoinVertical(lipgloss.Left, header, alnStr)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left, header, bgStr, alnStr)
	}
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
	text := lipgloss.NewStyle().Align(lipgloss.Center).Render("As you fill in the details of your new character, this character sheet will reflect the traits and choises you have selected.")
	return lipgloss.JoinVertical(lipgloss.Center, header, text)
}

func (m Model) View() string {
	if m.termWidth == 0 || m.termHeight <= 2 {
		return "Initializing..."
	}
	// Blank character placeholder text
	blankCharacter := m.c.Class == character.ClassNone &&
		m.c.Species == character.SpeciesNone &&
		m.c.Background == character.BackgroundNone &&
		m.c.Alignment == character.AlignmentNone &&
		!hasAbilityScores(m.c)

	sidebarWidth := m.menu.Width
	contentWidth := m.termWidth - sidebarWidth - 4

	content := displayCharacterOverview(m.c)

	if blankCharacter {
		content = blankCharacterGreeter()
	}
	contentStyle := lipgloss.NewStyle().
		Width(contentWidth).
		Height(m.termHeight - 2).
		PaddingLeft(2)

	baseView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.menu.View(),
		contentStyle.Render(content),
	)

	if !m.isEditingName {
		return baseView
	}

	modalContent := lipgloss.JoinVertical(
		lipgloss.Center,
		style.Title.Render("Edit Name"),
		"",
		m.textInput.View(),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("press enter to save • esc to cancel"),
	)

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		style.Popup.Render(modalContent),
	)
}

func (m *Model) RefreshSidebar() {
	items := MakeCharsheetSidebarItems(m.c)
	m.menu.SetItems(items)
}
