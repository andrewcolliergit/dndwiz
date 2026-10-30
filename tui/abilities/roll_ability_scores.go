package abilities

import (
	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/score"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"
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

type rollModel struct {
	c               *character.Character
	rolled          [6]int
	rolledCache     [6]int
	selectedAbility ability.Ability
	selectedRoll    int
	rollBoxVisible  bool
	focus           window
	confIndex       confSel
	width           int
	height          int
}

func (m rollModel) firstEmptyRoll() int {
	for i, roll := range m.rolled {
		if roll == 0 {
			return i
		}
	}
	return -1
}

func (m rollModel) firstAvailableRoll() int {
	for i, roll := range m.rolled {
		if roll > 0 {
			return i
		}
	}
	return -1
}

func (m rollModel) rightAvailableRoll() int {
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

func (m rollModel) leftAvailableRoll() int {
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

func NewRollModel(c *character.Character) tea.Model {
	c.AbilityScore = score.Ability{}
	rolled := utils.RollAllAbilityScores()
	return rollModel{
		c:               c,
		rolled:          rolled,
		rolledCache:     rolled,
		selectedAbility: ability.Strength,
		selectedRoll:    -1,
		rollBoxVisible:  true,
		confIndex:       confConfirm,
	}
}

// Navigation: Ability box

func (m *rollModel) abilityNavRight() {
	if m.selectedAbility < ability.Charisma {
		m.selectedAbility++
	}
}

func (m *rollModel) abilityNavLeft() {
	if m.selectedAbility > ability.Strength {
		m.selectedAbility--
	}
}

func (m *rollModel) abilitySelect() {
	m.focus = winRollBox
	m.selectedRoll = m.firstAvailableRoll()
}

// Navigation: Roll box

func (m *rollModel) rollSelect() {
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

func (m *rollModel) rollNavRight() {
	m.selectedRoll = m.rightAvailableRoll()
}

func (m *rollModel) rollNavLeft() {
	m.selectedRoll = m.leftAvailableRoll()
}

func (m rollModel) Init() tea.Cmd {
	return nil
}

// Navigation: Confirmation menu

func (m *rollModel) confReset() {
	m.c.AbilityScore = score.Ability{}
	m.rolled = m.rolledCache
	m.focus = winAbilitiesBox
	m.rollBoxVisible = true
	m.confIndex = 0
}

func (m *rollModel) confNavLeft() {
	if m.confIndex < numConf {
		m.confIndex++
	}
}

func (m *rollModel) confNavRight() {
	if m.confIndex > 0 {
		m.confIndex--
	}
}

func (m rollModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
					return m, pager.SwitchTo(pager.ViewCharsheet)
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

func (m rollModel) View() string {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Ability Points")

	abilityBox := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).Render(style.RenderAbilityScoresRow(m.c, m.selectedAbility, m.focus == winAbilitiesBox))

	var bottomSection string
	if m.focus == winConf {
		// Render Clear / Confirm menu options
		clearStyle := lipgloss.NewStyle().Padding(0, 2)
		confirmStyle := lipgloss.NewStyle().Padding(0, 2)

		if m.confIndex == 0 {
			clearStyle = clearStyle.Inherit(style.Label)
		} else {
			confirmStyle = confirmStyle.Inherit(style.Label)
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
			Align(lipgloss.Center).Render(style.RenderAbilityRollsRow(&m.rolled, m.selectedRoll, m.focus == winRollBox))
	}

	return lipgloss.JoinVertical(lipgloss.Center, title, " ", abilityBox, "", bottomSection)
}
