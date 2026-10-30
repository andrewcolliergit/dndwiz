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
		config                *Config
		state                 pager.CurrentView
		history               []pager.CurrentView
		currentModel          tea.Model
		termWidth, termHeight int
	}
)

func (m model) Init() tea.Cmd { return nil }

func (m *model) propagateSize(msg tea.WindowSizeMsg) {
	if m.currentModel != nil {
		m.currentModel, _ = m.currentModel.Update(msg)
	}
}

// Helper to initialize/refresh views based on state
func (m *model) initView(v pager.CurrentView) {
	m.state = v
	switch v {
	case pager.Title:
		m.currentModel = title.NewTitleScreenModel()
	case pager.Charsheet:
		m.currentModel = charactersheet.NewCharsheetModel(m.config.Character)
	case pager.ClassAbilitiesSelection:
		m.currentModel = classes.NewPrimaryAbilityModel(m.config.Character)
	case pager.ClassSelection:
		m.currentModel = classes.NewModel(m.config.Character)
	case pager.SpeciesSelection:
		m.currentModel = species.NewModel(m.config.Character)
	case pager.SpeciesSizeSelection:
		m.currentModel = species.NewSizeModel(m.config.Character)
	case pager.SpeciesAncestrySelection:
		m.currentModel = species.NewDragonModel(m.config.Character)
	case pager.SpeciesLineageSelection:
		m.currentModel = species.NewElfModel(m.config.Character)
	case pager.BackgroundSelection:
		m.currentModel = background.NewModel(m.config.Character)
	case pager.AlignmentSelection:
		m.currentModel = alignment.NewModel(m.config.Character)
	case pager.AbilityScoresMethodSelection:
		m.currentModel = abilities.NewOptionModel(m.config.Character)
	case pager.AbilityScoresRoll:
		m.currentModel = abilities.NewRollModel(m.config.Character)
	case pager.AbilityScoresBuy:
		m.currentModel = abilities.NewBuyModel(m.config.Character)
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

	if m.currentModel != nil {
		m.currentModel, cmd = m.currentModel.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if m.currentModel != nil {
		return m.currentModel.View()
	}
	return "UNKNOWN VIEW"
}

func NewProgram(cfg *Config) *tea.Program {
	m := model{
		config: cfg,
		state:  pager.Title,
	}
	m.initView(pager.Title)
	return tea.NewProgram(m, tea.WithAltScreen())
}
