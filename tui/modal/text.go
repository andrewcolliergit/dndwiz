// Package modal/text.go
package modal

import (
	"dndwiz/character"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputConfig struct {
	C           *character.Character
	Title       string
	Value       string
	Placeholder string
	Width       int
	EnterHelp   string
}

type textInputModel struct {
	cfg TextInputConfig
	ti  textinput.Model
}

func Text(cfg TextInputConfig) tea.Cmd {
	ti := textinput.New()
	ti.SetValue(cfg.Value)
	ti.Placeholder = cfg.Placeholder
	ti.Width = max(cfg.Width, 30)
	ti.Focus()
	ti.Prompt = "> "
	return pager.OpenModal(textInputModel{
		cfg: cfg,
		ti:  ti,
	})
}

func (m textInputModel) Init() tea.Cmd {
	return nil
}

func (m textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.ti.Blur()
			m.cfg.C.Name = m.ti.Value()
			return m, pager.CloseModal()
		case "esc":
			m.ti.Blur()
			return m, pager.CloseModal()
		default:
			m.ti, cmd = m.ti.Update(msg)
		}
	}
	return m, cmd
}

func (m textInputModel) View() string {
	return style.TypesetTextBox(
		m.cfg.Title,
		m.ti.View(),
		style.Hint{Key: "enter", Desc: m.cfg.EnterHelp},
		style.Hint{Key: "esc", Desc: "cancel"},
	)
}
