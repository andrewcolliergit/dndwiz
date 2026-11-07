// Package modal/confirm.go
package modal

import (
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Options struct {
	Title, Body  string
	ConfirmLabel string // defaults to "Confirm"
	CancelLabel  string // defaults to "Cancel"
	OnConfirm    tea.Cmd
	OnCancel     tea.Cmd // optional
}

type confirm struct {
	opts   Options
	onYes  bool
	width  int
	height int
}

// Confirm is what callers use.
func Confirm(o Options) tea.Cmd {
	if o.ConfirmLabel == "" {
		o.ConfirmLabel = "Confirm"
	}
	if o.CancelLabel == "" {
		o.CancelLabel = "Cancel"
	}
	return pager.OpenModal(confirm{opts: o})
}

func (m confirm) Init() tea.Cmd { return nil }

func (m confirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right", "h", "l", "tab":
			m.onYes = !m.onYes
		case "y":
			return m, tea.Sequence(pager.CloseModal(), m.opts.OnConfirm)
		case "n", "esc":
			return m, tea.Sequence(pager.CloseModal(), m.opts.OnCancel)
		case "enter":
			if m.onYes {
				return m, tea.Sequence(pager.CloseModal(), m.opts.OnConfirm)
			}
			return m, tea.Sequence(pager.CloseModal(), m.opts.OnCancel)
		}
	}
	return m, nil
}

func (m confirm) View() string {
	yes, no := m.opts.ConfirmLabel, m.opts.CancelLabel
	if m.onYes {
		yes = style.Label.Render("[ " + yes + " ]")
		no = "  " + no + "  "
	} else {
		yes = "  " + yes + "  "
		no = style.Label.Render("[ " + no + " ]")
	}
	body := lipgloss.JoinVertical(lipgloss.Left,
		style.Title.Render(m.opts.Title),
		style.TypesetBody(m.opts.Body, 30),
		"",
		lipgloss.JoinHorizontal(lipgloss.Center, yes, "   ", no),
	)
	return style.Popup.Render(body)
}

func (m confirm) Help() []style.Hint {
	return []style.Hint{
		{Key: "←/→", Desc: "choose"},
		{Key: "enter", Desc: "select"},
		{Key: "y/n", Desc: "yes/no"},
	}
}
