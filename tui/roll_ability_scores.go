package tui

import (
	"dndwiz/character"
	"dndwiz/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	window  int
	confSel int
)

const (
	winAbilitiesBox window = iota
	winRollBox
	winConf
)

const (
	confConfirm confSel = iota
	confReset
	numConf
)

type abilityScoresRollModel struct {
	c               *character.Character
	rolled          [6]int
	rolledCache     [6]int
	selectedAbility character.Ability
	selectedRoll    int
	rollBoxVisible  bool
	focus           window
	confIndex       confSel
	width           int
	height          int
}

func (m abilityScoresRollModel) firstEmptyRoll() int {
	for i, roll := range m.rolled {
		if roll == 0 {
			return i
		}
	}
	return -1
}

func (m abilityScoresRollModel) firstAvailableRoll() int {
	for i, roll := range m.rolled {
		if roll > 0 {
			return i
		}
	}
	return -1
}

func (m abilityScoresRollModel) rightAvailableRoll() int {
	if m.selectedRoll < 0 {
		return m.firstAvailableRoll()
	}
	for i := m.selectedRoll + 1; i < len(m.rolled); i++ {
		if m.rolled[i] > 0 {
			return i
		}
	}
	return m.selectedRoll
}

func (m abilityScoresRollModel) leftAvailableRoll() int {
	if m.selectedRoll < 0 {
		return m.firstAvailableRoll()
	}
	for i := m.selectedRoll - 1; i >= 0; i-- {
		if m.rolled[i] > 0 {
			return i
		}
	}
	return m.selectedRoll
}

func NewAbilityScoresRollModel(c *character.Character) tea.Model {
	c.AbilityScore = character.AbilityScores{}
	rolled := utils.RollAllAbilityScores()
	return abilityScoresRollModel{
		c:               c,
		rolled:          rolled,
		rolledCache:     rolled,
		selectedAbility: character.AbilityStrength,
		selectedRoll:    -1,
		rollBoxVisible:  true,
		confIndex:       confConfirm,
	}
}

// Navigation: Ability box

func (m *abilityScoresRollModel) abilityNavRight() {
	if m.selectedAbility < character.AbilityCharisma {
		m.selectedAbility++
	}
}

func (m *abilityScoresRollModel) abilityNavLeft() {
	if m.selectedAbility > character.AbilityStrength {
		m.selectedAbility--
	}
}

func (m *abilityScoresRollModel) abilitySelect() {
	m.focus = winRollBox
	m.selectedRoll = m.firstAvailableRoll()
}

// Navigation: Roll box

func (m *abilityScoresRollModel) rollSelect() {
	m.focus = winAbilitiesBox

	if m.c.AbilityScore[m.selectedAbility-1] > 0 {
		m.rolled[m.firstEmptyRoll()] = m.c.AbilityScore[m.selectedAbility-1]
	}

	m.c.AbilityScore[m.selectedAbility-1] = m.rolled[m.selectedRoll]
	m.rolled[m.selectedRoll] = 0
	m.selectedRoll = -1
	if m.firstAvailableRoll() == -1 {
		m.focus = winConf
		m.rollBoxVisible = false
	}
}

func (m *abilityScoresRollModel) rollNavRight() {
	m.selectedRoll = m.rightAvailableRoll()
}

func (m *abilityScoresRollModel) rollNavLeft() {
	m.selectedRoll = m.leftAvailableRoll()
}

func (m abilityScoresRollModel) Init() tea.Cmd {
	return nil
}

// Navigation: Confirmation menu

func (m *abilityScoresRollModel) confReset() {
	m.c.AbilityScore = character.AbilityScores{}
	m.rolled = m.rolledCache
	m.focus = winAbilitiesBox
	m.rollBoxVisible = true
	m.confIndex = 0
}

func (m *abilityScoresRollModel) confNavLeft() {
	if m.confIndex < numConf {
		m.confIndex++
	}
}

func (m *abilityScoresRollModel) confNavRight() {
	if m.confIndex > 0 {
		m.confIndex--
	}
}

func (m abilityScoresRollModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch m.focus {
		case winAbilitiesBox:
			switch msg.String() {
			case "enter":
				m.abilitySelect()
			case "left", "h":
				m.abilityNavLeft()
			case "right", "l":
				m.abilityNavRight()
			}
		case winRollBox:
			switch msg.String() {
			case "enter":
				m.rollSelect()
			case "left", "h":
				m.rollNavLeft()
			case "right", "l":
				m.rollNavRight()
			}
		case winConf:
			switch msg.String() {
			case "enter":
				if m.confIndex == confConfirm {
					return m, switchToView(viewCharsheet)
				}
				m.confReset()
			case "l", "right":
				m.confNavRight()
			case "h", "left":
				m.confNavLeft()
			}
		}
	}
	return m, nil
}

func (m abilityScoresRollModel) View() string {
	title := titleStyle.Width(m.width).Align(lipgloss.Center).Render("Ability Points")

	abilityBox := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).Render(renderAbilityScoresRow(m.c, m.selectedAbility, m.focus == winAbilitiesBox))

	var bottomSection string
	if m.focus == winConf {
		// Render Clear / Confirm menu options
		clearStyle := lipgloss.NewStyle().Padding(0, 2)
		confirmStyle := lipgloss.NewStyle().Padding(0, 2)

		if m.confIndex == 0 {
			clearStyle = clearStyle.Inherit(labelStyle)
		} else {
			confirmStyle = confirmStyle.Inherit(labelStyle)
		}

		menuRow := lipgloss.JoinHorizontal(
			lipgloss.Center,
			clearStyle.Render("[ Confirm ]"),
			"   ",
			confirmStyle.Render("[ Clear ]"),
		)

		bottomSection = lipgloss.NewStyle().
			Width(m.width).
			Align(lipgloss.Center).
			Render(lipgloss.JoinVertical(lipgloss.Center, "Are you finished?", menuRow))
	} else {
		bottomSection = lipgloss.NewStyle().
			Width(m.width).
			Align(lipgloss.Center).Render(renderAbilityRollsRow(&m.rolled, m.selectedRoll, m.focus == winRollBox))
	}

	return lipgloss.JoinVertical(lipgloss.Center, title, " ", abilityBox, "", bottomSection)
}
