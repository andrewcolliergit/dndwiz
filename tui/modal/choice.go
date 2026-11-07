// Package modal/choice.go
package modal

import (
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ChoiceOption struct {
	Label string
	Cmd   tea.Cmd // runs after the modal closes; nil just closes it
}

type ChoiceOptions struct {
	Title, Body string
	Options     []ChoiceOption // Esc closes the modal without running any option
}

type choice struct {
	opts   ChoiceOptions
	cursor int
}

func Choice(o ChoiceOptions) tea.Cmd {
	return pager.OpenModal(choice{opts: o})
}

func (m choice) Init() tea.Cmd { return nil }

func (m choice) pick(i int) (tea.Model, tea.Cmd) {
	return m, tea.Sequence(pager.CloseModal(), m.opts.Options[i].Cmd)
}

func (m choice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok || len(m.opts.Options) == 0 {
		return m, nil
	}

	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.opts.Options)-1 {
			m.cursor++
		}
	case "enter", " ":
		return m.pick(m.cursor)
	case "esc":
		return m, pager.CloseModal()
	}
	return m, nil
}

func (m choice) View() string {
	rows := []string{
		style.Title.Render(m.opts.Title),
		style.TypesetBody(m.opts.Body, 30),
		"",
	}
	for i, o := range m.opts.Options {
		if i == m.cursor {
			rows = append(rows, style.Label.Render("> "+o.Label))
		} else {
			rows = append(rows, "  "+o.Label)
		}
	}
	return style.Popup.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func (m choice) Help() []style.Hint {
	return []style.Hint{
		{Key: "↑/↓", Desc: "move"},
		{Key: "enter", Desc: "select"},
		{Key: "esc", Desc: "back"},
	}
}
