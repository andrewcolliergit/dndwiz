package tui

import (
	"fmt"
	"strings"

	"dndwiz/character"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type charsheetModel struct {
	c                     *character.Character
	title                 string
	menu                  sidebarModel
	textInput             textinput.Model
	isEditingName         bool
	termWidth, termHeight int
}

func (m charsheetModel) Init() tea.Cmd {
	return nil
}

func MakeCharsheetSidebarItems(c *character.Character) SidebarItemList {
	return SidebarItemList{
		{
			title: "Name",
			desc:  c.Name,
			data:  viewNameTextbox,
		},
		{
			title: "Class",
			desc:  character.GetClass(c.Class).Name,
			data:  viewClassSelection,
		},
		{
			title: "Species",
			desc:  character.GetSpecies(c.Species).Name,
			data:  viewSpeciesSelection,
		},
		{
			title: "Background",
			desc:  character.GetBackground(c.Background).Name,
			data:  viewBackgroundSelection,
		},
		{
			title: "Ability Scores",
			desc:  c.AbilityScoreMethod.String(),
			data:  viewAbilityScoresMethodSelection,
		},
		{
			title: "Alignment",
			desc:  character.GetAlignment(c.Alignment).Name,
			data:  viewAlignmentSelection,
		},
	}
}

func NewCharsheetModel(c *character.Character) charsheetModel {
	ti := textinput.New()
	ti.Placeholder = c.Name
	ti.CharLimit = 32
	ti.Width = 24

	items := MakeCharsheetSidebarItems(c)
	sidebar := newSidebar("Edit Character", items, 32, true)

	return charsheetModel{
		c:         c,
		title:     "Character Sheet",
		menu:      sidebar,
		textInput: ti,
	}
}

func (m charsheetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if v, ok := GetItemData[currentView](m.menu); ok {
				if v == viewNameTextbox {
					m.isEditingName = true
					m.textInput.SetValue(m.c.Name)
					m.textInput.Focus()
					return m, textinput.Blink
				}
				return m, switchToView(v)
			}
		}
	}
	m.menu, cmd = m.menu.Update(msg, m.termHeight)
	return m, cmd
}

func headerDisplay(c *character.Character) string {
	return titleStyle.Render(c.Name)
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

	return taglineStyle.Render(output)
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
	return renderAbilityScoresRow(c, character.AbilityNone, false)
}

func primaryAbilitiesDisplay(c *character.Character) string {
	if len(c.PrimaryAbilities) == 0 {
		return ""
	}
	return labelStyle.Render("Primary Abilities: ") + badgeStyle.Render(c.PrimaryAbilities.String())
}

func speciesAttributesDisplay(c *character.Character) string {
	if c.Species == character.SpeciesNone {
		return ""
	}

	spec := character.GetSpecies(c.Species)

	header := subheaderStyle.Render("Species Attributes")
	info := fmt.Sprintf(
		"%s %s\n%s %s\n%s %s",
		labelStyle.Render("Type:"), spec.CreatureType,
		labelStyle.Render("Size:"), c.Size,
		labelStyle.Render("Speed:"), spec.Speed.String(),
	)

	return lipgloss.JoinVertical(lipgloss.Left, header, info)
}

func backgroundDisplay(c *character.Character) string {
	hasBg := c.Background > character.BackgroundNone
	hasAln := c.Alignment > character.AlignmentNone

	if !hasBg && !hasAln {
		return ""
	}

	header := subheaderStyle.Render("Background Highlights")

	var (
		bgStr  string
		alnStr string
	)

	if hasBg {
		bg := character.GetBackground(c.Background)
		var bgAbilities []string
		for _, ability := range character.GetBackground(c.Background).Abilities {
			bgAbilities = append(bgAbilities, badgeStyle.Render(character.GetAbility(ability).Abbr))
		}

		bgAbilitiesStr := strings.Join(bgAbilities, " ")

		bgStr = fmt.Sprintf(
			"%s %s\n%s %s\n%s %s",
			labelStyle.Render("Feat:"), bg.Feat.String(),
			labelStyle.Render("Tool:"), bg.Tool,
			labelStyle.Render("Abilities:"), bgAbilitiesStr,
		)
	}

	if hasAln {
		aln := character.GetAlignment(c.Alignment)
		alnStr = labelStyle.Render("Alignment:") + " " + aln.Name
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

func (m charsheetModel) View() string {
	if m.termWidth == 0 || m.termHeight <= 2 {
		return "Initializing..."
	}

	sidebarWidth := m.menu.width
	contentWidth := m.termWidth - sidebarWidth - 4

	content := displayCharacterOverview(m.c)

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
		titleStyle.Render("Edit Name"),
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
		popupStyle.Render(modalContent),
	)
}

func (m *charsheetModel) RefreshSidebar() {
	items := MakeCharsheetSidebarItems(m.c)
	m.menu.SetItems(items)
}
