// Package modal/notice.go
package modal

import (
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NoticeOptions struct {
	Title, Body string
	ButtonLabel string  // defaults to "OK"
	OnClose     tea.Cmd // runs after the modal closes
}

type notice struct{ opts NoticeOptions }

func Notice(o NoticeOptions) tea.Cmd {
	if o.ButtonLabel == "" {
		o.ButtonLabel = "OK"
	}
	return pager.OpenModal(notice{opts: o})
}

func (m notice) Init() tea.Cmd { return nil }

func (m notice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "enter", "esc", " ", "y":
			return m, tea.Sequence(pager.CloseModal(), m.opts.OnClose)
		}
	}
	return m, nil
}

func (m notice) View() string {
	body := lipgloss.JoinVertical(lipgloss.Left,
		style.Title.Render(m.opts.Title),
		style.TypesetBody(m.opts.Body, 30),
		"",
		style.Label.Render("[ "+m.opts.ButtonLabel+" ]"),
	)
	return style.Popup.Render(body)
}
