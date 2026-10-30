// Package pager
package pager

import tea "github.com/charmbracelet/bubbletea"

// Views
type (
	CurrentView   int
	SwitchViewMsg CurrentView
)

const (
	ViewQuit CurrentView = iota
	ViewTitleScreen
	ViewCharsheet
	ViewClassSelection
	ViewPrimaryAbilitiesSelection
	ViewSpeciesSizeSelection
	ViewSpeciesSelection
	ViewBackgroundSelection
	ViewNameTextbox
	ViewAlignmentSelection
	ViewDraconicAncestrySelection
	ViewElvenLineageSelection
	ViewEquipmentSelection
	ViewAbilityScoresMethodSelection
	ViewAbilityScoresRoll
	ViewAbilityScoresBuy
)

func SwitchTo(v CurrentView) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg(v)
	}
}
