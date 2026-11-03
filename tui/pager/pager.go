// Package pager
package pager

import tea "github.com/charmbracelet/bubbletea"

type (
	CurrentView   int
	SwitchViewMsg CurrentView
)

const (
	Title CurrentView = iota
	PauseMenu
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

var parent = map[CurrentView]CurrentView{
	// Top-level editors
	NameTextbox:                  Charsheet,
	ClassSelection:               Charsheet,
	SpeciesSelection:             Charsheet,
	BackgroundSelection:          Charsheet,
	AlignmentSelection:           Charsheet,
	AbilityScoresMethodSelection: Charsheet,

	// Sub-pages
	ClassAbilitiesSelection:  ClassSelection,
	SpeciesSizeSelection:     SpeciesSelection,
	SpeciesAncestrySelection: SpeciesSelection,
	SpeciesLineageSelection:  SpeciesSelection,
	AbilityScoresRoll:        AbilityScoresMethodSelection,
	AbilityScoresBuy:         AbilityScoresMethodSelection,
}

func SwitchTo(v CurrentView) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg(v)
	}
}

func Parent(v CurrentView) (CurrentView, bool) {
	p, ok := parent[v]
	return p, ok
}

type (
	GoUpMsg   struct{}
	GoBackMsg struct{}
)

func Up() tea.Cmd {
	return func() tea.Msg { return GoUpMsg{} }
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
