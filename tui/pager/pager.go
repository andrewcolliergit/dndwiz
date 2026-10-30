// Package pager
package pager

import tea "github.com/charmbracelet/bubbletea"

// Views
type (
	CurrentView   int
	SwitchViewMsg CurrentView
)

const (
	Quit CurrentView = iota
	Title
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
	EquipmentSelection
	AbilityScoresMethodSelection
	AbilityScoresRoll
	AbilityScoresBuy
)

func SwitchTo(v CurrentView) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg(v)
	}
}
