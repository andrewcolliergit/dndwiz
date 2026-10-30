// Package tui
package tui

import (
	"dndwiz/character"
	"dndwiz/tui/abilities"
	"dndwiz/tui/alignment"
	"dndwiz/tui/background"
	charactersheet "dndwiz/tui/character_sheet"
	"dndwiz/tui/classes"
	"dndwiz/tui/pager"
	"dndwiz/tui/species"
	"dndwiz/tui/title"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	goBackMsg struct{}

	Config struct {
		Character *character.Character
	}

	model struct {
		config                       *Config
		state                        pager.CurrentView
		history                      []pager.CurrentView
		titleScreen                  tea.Model
		charsheet                    tea.Model
		classes                      tea.Model
		species                      tea.Model
		background                   tea.Model
		alignment                    tea.Model
		primaryAbilities             tea.Model
		speciesSize                  tea.Model
		draconicAncestry             tea.Model
		elvenLineage                 tea.Model
		abilityScoresMethodSelection tea.Model
		abilityScoresRoll            tea.Model
		abilityScoresBuy             tea.Model
		termWidth, termHeight        int
	}
)

func (m model) Init() tea.Cmd { return nil }

func (m *model) propagateSize(msg tea.WindowSizeMsg) {
	if m.titleScreen != nil {
		m.titleScreen, _ = m.titleScreen.Update(msg)
	}
	if m.charsheet != nil {
		m.charsheet, _ = m.charsheet.Update(msg)
	}
	if m.classes != nil {
		m.classes, _ = m.classes.Update(msg)
	}
	if m.species != nil {
		m.species, _ = m.species.Update(msg)
	}
	if m.background != nil {
		m.background, _ = m.background.Update(msg)
	}
	if m.alignment != nil {
		m.alignment, _ = m.alignment.Update(msg)
	}
	if m.primaryAbilities != nil {
		m.primaryAbilities, _ = m.primaryAbilities.Update(msg)
	}
	if m.speciesSize != nil {
		m.speciesSize, _ = m.speciesSize.Update(msg)
	}
	if m.draconicAncestry != nil {
		m.draconicAncestry, _ = m.draconicAncestry.Update(msg)
	}
	if m.elvenLineage != nil {
		m.elvenLineage, _ = m.elvenLineage.Update(msg)
	}
	if m.abilityScoresMethodSelection != nil {
		m.abilityScoresMethodSelection, _ = m.abilityScoresMethodSelection.Update(msg)
	}
	if m.abilityScoresRoll != nil {
		m.abilityScoresRoll, _ = m.abilityScoresRoll.Update(msg)
	}
	if m.abilityScoresBuy != nil {
		m.abilityScoresBuy, _ = m.abilityScoresBuy.Update(msg)
	}
}

// Helper to initialize/refresh views based on state
func (m *model) initView(v pager.CurrentView) {
	m.state = v
	switch v {
	case pager.Title:
		m.titleScreen = title.NewTitleScreenModel()
	case pager.Charsheet:
		//m.charsheet.RefreshSidebar()
		m.charsheet = charactersheet.NewCharsheetModel(m.config.Character)
	case pager.ClassAbilitiesSelection:
		m.primaryAbilities = classes.NewPrimaryAbilityModel(m.config.Character)
	case pager.ClassSelection:
		m.classes = classes.NewModel(m.config.Character)
	case pager.SpeciesSelection:
		m.species = species.NewModel(m.config.Character)
	case pager.SpeciesSizeSelection:
		m.speciesSize = species.NewSizeModel(m.config.Character)
	case pager.SpeciesAncestrySelection:
		m.draconicAncestry = species.NewDragonModel(m.config.Character)
	case pager.SpeciesLineageSelection:
		m.elvenLineage = species.NewElfModel(m.config.Character)
	case pager.BackgroundSelection:
		m.background = background.NewModel(m.config.Character)
	case pager.AlignmentSelection:
		m.alignment = alignment.NewModel(m.config.Character)
	case pager.AbilityScoresMethodSelection:
		m.abilityScoresMethodSelection = abilities.NewOptionModel(m.config.Character)
	case pager.AbilityScoresRoll:
		m.abilityScoresRoll = abilities.NewRollModel(m.config.Character)
	case pager.AbilityScoresBuy:
		m.abilityScoresBuy = abilities.NewBuyModel(m.config.Character)
	}

	if m.termWidth > 0 && m.termHeight > 0 {
		m.propagateSize(tea.WindowSizeMsg{Width: m.termWidth, Height: m.termHeight})
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return goBackMsg{} }
		}

	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		m.propagateSize(msg)

	case goBackMsg:
		if len(m.history) == 0 {
			return m, nil
		}
		prevState := m.history[len(m.history)-1]
		m.history = m.history[:len(m.history)-1]
		m.initView(prevState)
		return m, nil

	case pager.SwitchViewMsg:
		target := pager.CurrentView(msg)
		if m.state != target {
			m.history = append(m.history, m.state)
		}

		m.initView(target)
		return m, nil
	}

	switch m.state {
	case pager.Title:
		m.titleScreen, cmd = m.titleScreen.Update(msg)

	case pager.Charsheet:
		m.charsheet, cmd = m.charsheet.Update(msg)

	case pager.ClassSelection:
		m.classes, cmd = m.classes.Update(msg)

	case pager.SpeciesSelection:
		m.species, cmd = m.species.Update(msg)

	case pager.BackgroundSelection:
		m.background, cmd = m.background.Update(msg)

	case pager.AlignmentSelection:
		m.alignment, cmd = m.alignment.Update(msg)

	case pager.ClassAbilitiesSelection:
		m.primaryAbilities, cmd = m.primaryAbilities.Update(msg)

	case pager.SpeciesSizeSelection:
		m.speciesSize, cmd = m.speciesSize.Update(msg)

	case pager.SpeciesAncestrySelection:
		m.draconicAncestry, cmd = m.draconicAncestry.Update(msg)

	case pager.SpeciesLineageSelection:
		m.elvenLineage, cmd = m.elvenLineage.Update(msg)
	case pager.AbilityScoresMethodSelection:
		m.abilityScoresMethodSelection, cmd = m.abilityScoresMethodSelection.Update(msg)
	case pager.AbilityScoresRoll:
		m.abilityScoresRoll, cmd = m.abilityScoresRoll.Update(msg)
	case pager.AbilityScoresBuy:
		m.abilityScoresBuy, cmd = m.abilityScoresBuy.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case pager.Title:
		return m.titleScreen.View()
	case pager.Charsheet:
		return m.charsheet.View()
	case pager.ClassSelection:
		return m.classes.View()
	case pager.SpeciesSelection:
		return m.species.View()
	case pager.BackgroundSelection:
		return m.background.View()
	case pager.AlignmentSelection:
		return m.alignment.View()
	case pager.ClassAbilitiesSelection:
		return m.primaryAbilities.View()
	case pager.SpeciesSizeSelection:
		return m.speciesSize.View()
	case pager.SpeciesAncestrySelection:
		return m.draconicAncestry.View()
	case pager.SpeciesLineageSelection:
		return m.elvenLineage.View()
	case pager.AbilityScoresMethodSelection:
		return m.abilityScoresMethodSelection.View()
	case pager.AbilityScoresRoll:
		return m.abilityScoresRoll.View()
	case pager.AbilityScoresBuy:
		return m.abilityScoresBuy.View()
	default:
		return "UNKNOWN VIEW"
	}
}

func NewProgram(cfg *Config) *tea.Program {
	m := model{
		config:                       cfg,
		state:                        pager.Title,
		titleScreen:                  title.NewTitleScreenModel(),
		charsheet:                    charactersheet.NewCharsheetModel(cfg.Character),
		classes:                      classes.NewModel(cfg.Character),
		primaryAbilities:             classes.NewPrimaryAbilityModel(cfg.Character),
		species:                      species.NewModel(cfg.Character),
		speciesSize:                  species.NewSizeModel(cfg.Character),
		draconicAncestry:             species.NewDragonModel(cfg.Character),
		elvenLineage:                 species.NewElfModel(cfg.Character),
		background:                   background.NewModel(cfg.Character),
		alignment:                    alignment.NewModel(cfg.Character),
		abilityScoresMethodSelection: abilities.NewOptionModel(cfg.Character),
		abilityScoresRoll:            abilities.NewRollModel(cfg.Character),
		abilityScoresBuy:             abilities.NewBuyModel(cfg.Character),
	}
	return tea.NewProgram(m, tea.WithAltScreen())
}
