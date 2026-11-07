// Package modal/help.go
package modal

import (
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type helpPage struct {
	title string
	build func(width int) string
	vp    viewport.Model
}

func HelpPage(title string, build func(width int) string) tea.Cmd {
	m := helpPage{title: title, build: build, vp: viewport.New(60, 15)}
	m.resize(80, 24)
	return pager.OpenModal(m)
}

func (m *helpPage) resize(termW, termH int) {
	m.vp.Width = min(max(termW-12, 30), 70)
	m.vp.Height = min(max(termH-10, 5), 20)
	m.vp.SetContent(m.build(m.vp.Width))
}

func (m helpPage) Init() tea.Cmd { return nil }

func (m helpPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.resize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "?", "q", "enter":
			return m, pager.CloseModal()
		}
		var cmd tea.Cmd
		m.vp, cmd = m.vp.Update(msg) // up/down, j/k, pgup/pgdn scroll
		return m, cmd
	}
	return m, nil
}

func (m helpPage) View() string {
	body := lipgloss.JoinVertical(lipgloss.Left,
		style.Title.Render(m.title),
		m.vp.View(),
	)
	return style.Popup.Width(m.vp.Width + 4).Render(body)
}

func (m helpPage) Help() []style.Hint {
	hints := []style.Hint{}
	if m.vp.TotalLineCount() > m.vp.Height {
		hints = append(hints, style.Hint{Key: "↑/↓", Desc: "scroll"})
	}
	return append(hints, style.Hint{Key: "esc", Desc: "close"})
}
