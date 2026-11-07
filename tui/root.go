// Package tui
package tui

import (
	"dndwiz/character"
	dialog "dndwiz/tui/modal"
	"dndwiz/tui/pager"
	content "dndwiz/tui/pages"
	"dndwiz/tui/shell"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	Config struct {
		Character *character.Character
	}

	model struct {
		config                *Config
		state                 pager.CurrentView
		currentModel          tea.Model
		modal                 tea.Model
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
		m.propagateSize()
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case pager.OpenModalMsg:
		m.modal = msg.Model
		if m.termWidth > 0 {
			m.modal, _ = m.modal.Update(m.bodySize())
		}
		return m, m.modal.Init()

	case pager.CloseModalMsg:
		m.modal = nil
		return m, func() tea.Msg { return pager.RefreshMsg{} }

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if m.modal != nil {
			m.modal, cmd = m.modal.Update(msg)
			return m, cmd
		}

		if msg.String() == "?" {
			if hp, ok := m.currentModel.(helpPager); ok {
				if title, body, ok := hp.HelpPage(); ok {
					return m, dialog.HelpPage(title, body)
				}
			}
			return m, nil
		}

		if msg.String() == "esc" {
			switch m.state {
			case pager.Charsheet:
				return m, dialog.Pause()
			case pager.AbilityScoresRoll:
				// the roll model decides: notice first, or leave immediately
			default:
				return m, pager.Up()
			}
		}

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		m.propagateSize()
		return m, nil

	case pager.GoUpMsg:
		return m.up()

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
	bodyH := max(m.termHeight-helpHeight, 1)

	var body string
	var src tea.Model
	switch {
	case m.modal != nil:
		body, src = style.ModalStyle(m.termWidth, bodyH, m.modal.View()), m.modal
	case m.currentModel != nil:
		body, src = m.currentModel.View(), m.currentModel
	default:
		return "UNKNOWN VIEW"
	}

	var hints []style.Hint
	if h, ok := src.(helper); ok {
		hints = h.Help()
	}

	// Pad the body to full height so the bar is pinned to the last row.
	body = lipgloss.NewStyle().Height(bodyH).Render(body)
	return lipgloss.JoinVertical(lipgloss.Left, body, style.TypesetHelpBar(m.termWidth, hints))
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

// func (m *model) propagateSize(msg tea.WindowSizeMsg) {
// 	if m.currentModel != nil {
// 		m.currentModel, _ = m.currentModel.Update(msg)
// 	}
// }

const helpHeight = 1

// helper is implemented by any view or modal that wants hints in the bar.
type helper interface{ Help() []style.Hint }

// helpPager is implemented by views that have a help page.
type helpPager interface {
	HelpPage() (title string, build func(width int) string, ok bool)
}

// bodySize is what children see: the terminal minus the help bar.
func (m model) bodySize() tea.WindowSizeMsg {
	return tea.WindowSizeMsg{Width: m.termWidth, Height: max(m.termHeight-helpHeight, 1)}
}

func (m *model) propagateSize() {
	if m.currentModel != nil {
		m.currentModel, _ = m.currentModel.Update(m.bodySize())
	}
	if m.modal != nil {
		m.modal, _ = m.modal.Update(m.bodySize())
	}
}

func (m *model) up() (tea.Model, tea.Cmd) {
	p, ok := pager.Parent(m.state)
	if !ok {
		return m, nil
	}
	m.initView(p)
	return m, nil
}

func (m *model) switchView(page pager.CurrentView) (tea.Model, tea.Cmd) {
	m.initView(page)
	return m, nil
}
