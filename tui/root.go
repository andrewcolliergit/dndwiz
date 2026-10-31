// Package tui
package tui

import (
	"dndwiz/character"
	"dndwiz/tui/content"
	"dndwiz/tui/pager"
	"dndwiz/tui/shell"

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
		viewCache             map[pager.CurrentView]tea.Model
		termWidth, termHeight int
	}
)

func (m model) Init() tea.Cmd { return nil }

// Helper to initialize/refresh views based on state
func (m *model) initView(v pager.CurrentView) {
	m.state = v

	if cachedModel, exists := m.viewCache[v]; exists {
		m.currentModel = cachedModel
	} else {
		switch v {
		case pager.Title:
			m.currentModel = shell.NewShell(m.config.Character, content.TitleConfig)
		case pager.Charsheet:
			m.currentModel = shell.NewShell(m.config.Character, content.CharsheetConfig)
		case pager.NameTextbox:
			m.currentModel = shell.NewShell(m.config.Character, content.NameConfig)
		case pager.ClassSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.ClassSelectionConfig)
		case pager.ClassAbilitiesSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.PrimaryAbilityConfig)
		case pager.SpeciesSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.SpeciesSelectionConfig)
		case pager.SpeciesSizeSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.SpeciesSizeConfig)
		case pager.SpeciesAncestrySelection:
			m.currentModel = shell.NewShell(m.config.Character, content.DraconicAncestryConfig)
		case pager.SpeciesLineageSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.ElvenLineageConfig)
		case pager.BackgroundSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.BackgroundSelectionConfig)
		case pager.AlignmentSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.AlignmentSelectionConfig)
		case pager.AbilityScoresMethodSelection:
			m.currentModel = shell.NewShell(m.config.Character, content.AbilityScoreMethodConfig)
		case pager.AbilityScoresRoll:
			m.currentModel = content.NewRollModel(m.config.Character)
		case pager.AbilityScoresBuy:
			m.currentModel = content.NewBuyModel(m.config.Character)
		}

		if m.viewCache == nil {
			m.viewCache = make(map[pager.CurrentView]tea.Model)
		}
	}

	m.viewCache[v] = m.currentModel

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
		return m.back()

	case pager.SwitchViewMsg:
		return m.switchView(pager.CurrentView(msg))
	}

	if m.currentModel != nil {
		m.currentModel, cmd = m.currentModel.Update(msg)

		if m.viewCache != nil {
			m.viewCache[m.state] = m.currentModel
		}
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

// -- Helper Methods --

func (m *model) propagateSize(msg tea.WindowSizeMsg) {
	if m.currentModel != nil {
		m.currentModel, _ = m.currentModel.Update(msg)
	}
}

func (m *model) back() (tea.Model, tea.Cmd) {
	if len(m.history) == 0 {
		return m, nil
	}
	prevState := m.history[len(m.history)-1]
	m.history = m.history[:len(m.history)-1]
	m.initView(prevState)
	return m, nil
}

func (m *model) switchView(page pager.CurrentView) (tea.Model, tea.Cmd) {
	if m.state != page {
		m.history = append(m.history, m.state)
	}

	m.initView(page)
	return m, nil
}
