package modal

import (
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pauseItem struct {
	label string
	cmd   tea.Cmd // nil means not available yet
}

type pause struct {
	items  []pauseItem
	cursor int
}

// Pause opens the pause menu modal.
func Pause() tea.Cmd {
	items := []pauseItem{
		{"Resume", pager.CloseModal()},
		{"Save", nil},   // TODO
		{"Export", nil}, // TODO
		{"Quit to Title", Confirm(Options{
			Title:     "Quit to title?",
			Body:      "Unsaved changes will be lost.",
			OnConfirm: pager.SwitchTo(pager.Title),
		})},
		{"Quit", Confirm(Options{
			Title:     "Quit?",
			Body:      "Are you sure you would like to quit the program?",
			OnConfirm: tea.Quit,
		})},
	}
	return pager.OpenModal(pause{items: items})
}

func (m pause) Init() tea.Cmd { return nil }

func (m pause) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
	case "esc":
		return m, pager.CloseModal()
	case "enter":
		if cmd := m.items[m.cursor].cmd; cmd != nil {
			return m, cmd
		}
	}
	return m, nil
}

func (m pause) View() string {
	rows := []string{style.Title.Render("Paused")}

	for i, it := range m.items {
		switch {
		case it.cmd == nil:
			rows = append(rows, style.Desc.Render("  "+it.label+" (soon)"))
		case i == m.cursor:
			rows = append(rows, style.Label.Render("> "+it.label))
		default:
			rows = append(rows, "  "+it.label)
		}
	}

	return style.Popup.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func (m pause) Help() []style.Hint {
	return []style.Hint{
		{Key: "↑/↓", Desc: "move"},
		{Key: "enter", Desc: "select"},
		{Key: "esc", Desc: "resume"},
	}
}
