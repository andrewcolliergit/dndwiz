// Package pager
package pager

import tea "github.com/charmbracelet/bubbletea"

type (
	GoBackMsg     struct{}
	CurrentView   int
	SwitchViewMsg CurrentView
)

const (
	Title CurrentView = iota
	Charsheet
	ClassSelection
	ClassAbilitiesSelection
	SpeciesSizeSelection
	SpeciesSelection
	BackgroundSelection
	NameTextbox
	AlignmentSelection
	SpeciesAncestrySelection
	SpeciesLineageSelection
	AbilityScoresMethodSelection
	AbilityScoresRoll
	AbilityScoresBuy
)

func SwitchTo(v CurrentView) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg(v)
	}
}

func Back() tea.Cmd {
	return func() tea.Msg { return GoBackMsg{} }
}

type (
	OpenModalMsg  struct{ Model tea.Model }
	CloseModalMsg struct{}
)

func OpenModal(m tea.Model) tea.Cmd {
	return func() tea.Msg { return OpenModalMsg{Model: m} }
}

func CloseModal() tea.Cmd {
	return func() tea.Msg { return CloseModalMsg{} }
}
