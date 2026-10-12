package tui

import (
	"fmt"

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

func displayCharacterOverview(c *character.Character) string {
	cls := character.GetClass(c.Class)
	spec := character.GetSpecies(c.Species)
	bg := character.GetBackground(c.Background)
	aln := character.GetAlignment(c.Alignment)

	// Ancestry included in species
	var species string
	switch c.Species {
	case character.SpeciesDragonborn:
		species = character.GetDraconicAncestry(c.DraconicAncestry).Name + " " + spec.Name
	case character.SpeciesElf:
		species = character.GetElvenLineage(c.ElvenLineage).Name
	default:
		species = spec.Name
	}

	// Header block
	nameStr := titleStyle.Render(c.Name)
	tagline := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("246")).
		Render(fmt.Sprintf("%s • %s %s", bg.Name, species, cls.Name))

	// Class Details Block
	primaryAbilities := labelStyle.Render("Primary Abilities: ") + badgeStyle.Render(c.PrimaryAbilities.String())

	abilityScores := renderAbilityScoresRow(c, character.AbilityNone, false)

	// Species Traits Block
	speciesDetails := fmt.Sprintf(
		"%s %s  │  %s %s  │  %s %s",
		labelStyle.Render("Type:"), spec.CreatureType,
		labelStyle.Render("Size:"), c.Size,
		labelStyle.Render("Speed:"), spec.Speed.String(),
	)

	// Background Summary Block
	bgDetails := fmt.Sprintf(
		"%s %s\n%s %s\n%s %s\n%s %s",
		labelStyle.Render("Feat:"), bg.Feat.String(),
		labelStyle.Render("Tool:"), bg.Tool,
		labelStyle.Render("Abilities:"), character.GetBackgroundAbilities(c.Background),
		labelStyle.Render("Alignment:"), aln.Name,
	)

	divwidth := 44

	return lipgloss.JoinVertical(
		lipgloss.Left,
		nameStr,
		tagline,
		abilityScores,
		renderDivider(divwidth),
		primaryAbilities,
		renderDivider(divwidth),
		labelStyle.Render("Species Attributes:"),
		speciesDetails,
		renderDivider(divwidth),
		labelStyle.Render("Background Highlights:"),
		bgDetails,
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

	popupStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Width(36).
		Align(lipgloss.Left)

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
