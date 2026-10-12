// Package tui/root.go
package tui

import (
	"dndwiz/character"

	tea "github.com/charmbracelet/bubbletea"
)

// Views
type currentView int

const (
	viewCharsheet currentView = iota
	viewClassSelection
	viewPrimaryAbilitiesSelection
	viewSpeciesSizeSelection
	viewSpeciesSelection
	viewBackgroundSelection
	viewNameTextbox
	viewAlignmentSelection
	viewDraconicAncestrySelection
	viewElvenLineageSelection
	viewEquipmentSelection
	viewAbilityScoresMethodSelection
	viewAbilityScoresRoll
	viewAbilityScoresBuy
)

type (
	switchViewMsg currentView
	goBackMsg     struct{}

	Config struct {
		Character *character.Character
	}

	rootModel struct {
		config                       *Config
		state                        currentView
		history                      []currentView
		charsheet                    charsheetModel
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

func switchToView(v currentView) tea.Cmd {
	return func() tea.Msg {
		return switchViewMsg(v)
	}
}

func (m rootModel) Init() tea.Cmd { return nil }

func (m *rootModel) propagateSize(msg tea.WindowSizeMsg) {
	if m.charsheet.c != nil {
		updated, _ := m.charsheet.Update(msg)
		m.charsheet = updated.(charsheetModel)
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
func (m *rootModel) initView(v currentView) {
	m.state = v
	switch v {
	case viewCharsheet:
		m.charsheet.RefreshSidebar()
	case viewPrimaryAbilitiesSelection:
		m.primaryAbilities = NewPrimaryAbilityModel(m.config.Character)
	case viewClassSelection:
		m.classes = NewClassesModel(m.config.Character)
	case viewSpeciesSelection:
		m.species = NewSpeciesModel(m.config.Character)
	case viewBackgroundSelection:
		m.background = NewBackgroundModel(m.config.Character)
	case viewAlignmentSelection:
		m.alignment = NewAlignmentModel(m.config.Character)
	case viewSpeciesSizeSelection:
		m.speciesSize = NewSpeciesSizeModel(m.config.Character)
	case viewDraconicAncestrySelection:
		m.draconicAncestry = NewDraconicAncestryModel(m.config.Character)
	case viewElvenLineageSelection:
		m.elvenLineage = NewElvenLineageModel(m.config.Character)
	case viewAbilityScoresMethodSelection:
		m.abilityScoresMethodSelection = NewAbilityScoresMethodSelectionModel(m.config.Character)
	case viewAbilityScoresRoll:
		m.abilityScoresRoll = NewAbilityScoresRollModel(m.config.Character)
	case viewAbilityScoresBuy:
		m.abilityScoresBuy = NewAbilityScoresBuyModel(m.config.Character)
	}

	if m.termWidth > 0 && m.termHeight > 0 {
		m.propagateSize(tea.WindowSizeMsg{Width: m.termWidth, Height: m.termHeight})
	}
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case switchViewMsg:
		target := currentView(msg)
		if m.state != target {
			m.history = append(m.history, m.state)
		}

		m.initView(target)
		return m, nil
	}

	switch m.state {
	case viewCharsheet:
		var updatedModel tea.Model
		updatedModel, cmd = m.charsheet.Update(msg)
		m.charsheet = updatedModel.(charsheetModel)

	case viewClassSelection:
		m.classes, cmd = m.classes.Update(msg)

	case viewSpeciesSelection:
		m.species, cmd = m.species.Update(msg)

	case viewBackgroundSelection:
		m.background, cmd = m.background.Update(msg)

	case viewAlignmentSelection:
		m.alignment, cmd = m.alignment.Update(msg)

	case viewPrimaryAbilitiesSelection:
		m.primaryAbilities, cmd = m.primaryAbilities.Update(msg)

	case viewSpeciesSizeSelection:
		m.speciesSize, cmd = m.speciesSize.Update(msg)

	case viewDraconicAncestrySelection:
		m.draconicAncestry, cmd = m.draconicAncestry.Update(msg)

	case viewElvenLineageSelection:
		m.elvenLineage, cmd = m.elvenLineage.Update(msg)
	case viewAbilityScoresMethodSelection:
		m.abilityScoresMethodSelection, cmd = m.abilityScoresMethodSelection.Update(msg)
	case viewAbilityScoresRoll:
		m.abilityScoresRoll, cmd = m.abilityScoresRoll.Update(msg)
	case viewAbilityScoresBuy:
		m.abilityScoresBuy, cmd = m.abilityScoresBuy.Update(msg)
	}

	return m, cmd
}

func (m rootModel) View() string {
	switch m.state {
	case viewCharsheet:
		return m.charsheet.View()
	case viewClassSelection:
		return m.classes.View()
	case viewSpeciesSelection:
		return m.species.View()
	case viewBackgroundSelection:
		return m.background.View()
	case viewAlignmentSelection:
		return m.alignment.View()
	case viewPrimaryAbilitiesSelection:
		return m.primaryAbilities.View()
	case viewSpeciesSizeSelection:
		return m.speciesSize.View()
	case viewDraconicAncestrySelection:
		return m.draconicAncestry.View()
	case viewElvenLineageSelection:
		return m.elvenLineage.View()
	case viewAbilityScoresMethodSelection:
		return m.abilityScoresMethodSelection.View()
	case viewAbilityScoresRoll:
		return m.abilityScoresRoll.View()
	case viewAbilityScoresBuy:
		return m.abilityScoresBuy.View()
	default:
		return "UNKNOWN VIEW"
	}
}

func NewProgram(cfg *Config) *tea.Program {
	m := rootModel{
		config:                       cfg,
		charsheet:                    NewCharsheetModel(cfg.Character),
		classes:                      NewClassesModel(cfg.Character),
		species:                      NewSpeciesModel(cfg.Character),
		background:                   NewBackgroundModel(cfg.Character),
		alignment:                    NewAlignmentModel(cfg.Character),
		primaryAbilities:             NewPrimaryAbilityModel(cfg.Character),
		speciesSize:                  NewSpeciesSizeModel(cfg.Character),
		draconicAncestry:             NewDraconicAncestryModel(cfg.Character),
		elvenLineage:                 NewElvenLineageModel(cfg.Character),
		abilityScoresMethodSelection: NewAbilityScoresMethodSelectionModel(cfg.Character),
		abilityScoresRoll:            NewAbilityScoresRollModel(cfg.Character),
		abilityScoresBuy:             NewAbilityScoresBuyModel(cfg.Character),
	}
	return tea.NewProgram(m, tea.WithAltScreen())
}
