// Package pages/roll_ability_scores.go
package pages

import (
	"fmt"
	"strings"
	"time"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/score"
	"dndwiz/dice"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	window      int
	confSel     int
	rollPhase   int
	introChoice int
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

const (
	phaseIntro rollPhase = iota
	phaseDice
	phaseAssign
)

const (
	introRoll introChoice = iota
	introAuto
	numIntro
)

const spinInterval = 70 * time.Millisecond

// spinMsg drives the spinning die. gen lets us ignore ticks from a stale chain.
type spinMsg struct{ gen int }

func spinTick(gen int) tea.Cmd {
	return tea.Tick(spinInterval, func(time.Time) tea.Msg { return spinMsg{gen: gen} })
}

// sumDropLowest returns the sum of the three highest dice and the index of the dropped die.
func sumDropLowest(d [4]int) (sum, dropped int) {
	total := 0
	for i, v := range d {
		total += v
		if v < d[dropped] {
			dropped = i
		}
	}
	return total - d[dropped], dropped
}

type rollModel struct {
	c        *character.Character
	phase    rollPhase
	swapFrom ability.Ability
	width    int
	height   int

	// intro
	introIndex introChoice

	// dice phase
	dice      [4]int
	dieIndex  int // which die is currently spinning (4 == all settled)
	dropped   int // index of dropped die, -1 if none yet
	rollCount int // which of the 6 scores is being rolled
	spinning  bool
	spinGen   int

	// assign phase
	rolled          [6]int
	rolledCache     [6]int
	selectedAbility ability.Ability
	selectedRoll    int
	rollBoxVisible  bool
	focus           window
	confIndex       confSel
}

func NewRollModel(c *character.Character) tea.Model {
	c.AbilityScore = score.Ability{}
	return rollModel{
		c:               c,
		phase:           phaseIntro,
		introIndex:      introRoll,
		dropped:         -1,
		rolled:          c.RolledAbilityScores,
		rolledCache:     c.RolledAbilityScores,
		selectedAbility: ability.Strength,
		selectedRoll:    -1,
		rollBoxVisible:  true,
		confIndex:       confConfirm,
	}
}

func (m rollModel) Init() tea.Cmd { return nil }

// --- Roll helpers (assign phase) ---

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
		m.rollBoxVisible = false
		m.confIndex = confConfirm
	}
}

func (m *rollModel) rollNavRight() { m.selectedRoll = m.rightAvailableRoll() }
func (m *rollModel) rollNavLeft()  { m.selectedRoll = m.leftAvailableRoll() }

// Navigation: Confirmation menu

func (m *rollModel) confReset() {
	m.c.AbilityScore = score.Ability{}
	m.rolled = m.rolledCache
	m.focus = winAbilitiesBox
	m.rollBoxVisible = true
	m.confIndex = confConfirm
	m.swapFrom = ability.None
}

func (m *rollModel) confNavLeft() {
	if m.confIndex > 0 {
		m.confIndex--
	}
}

func (m *rollModel) confNavRight() {
	if m.confIndex < numConf-1 {
		m.confIndex++
	}
}

// --- Phase transitions ---

func (m *rollModel) enterAssign() {
	m.phase = phaseAssign
	m.focus = winAbilitiesBox
	m.selectedRoll = -1
	m.rollBoxVisible = true
	m.confIndex = confConfirm
}

func (m *rollModel) beginManualRolling() tea.Cmd {
	m.phase = phaseDice
	m.rolled = [6]int{}
	m.rollCount = 0
	return m.beginSet()
}

// beginSet clears the dice and starts spinning the first one.
func (m *rollModel) beginSet() tea.Cmd {
	m.dice = [4]int{}
	m.dieIndex = 0
	m.dropped = -1
	return m.startSpin()
}

func (m *rollModel) startSpin() tea.Cmd {
	m.spinning = true
	m.spinGen++ // invalidates any previous tick chain
	m.dice[m.dieIndex] = dice.D6.Roll()
	return spinTick(m.spinGen)
}

func (m *rollModel) stopDie() tea.Cmd {
	m.dice[m.dieIndex] = dice.D6.Roll()
	m.dieIndex++

	if m.dieIndex < len(m.dice) {
		return m.startSpin()
	}

	// All four settled: drop lowest, record the sum.
	m.spinning = false
	m.spinGen++
	sum, dropped := sumDropLowest(m.dice)
	m.dropped = dropped
	m.rolled[m.rollCount] = sum
	return nil
}

func (m *rollModel) nextSet() tea.Cmd {
	m.rollCount++
	if m.rollCount >= len(m.rolled) {
		m.c.RolledAbilityScores = m.rolled
		m.rolledCache = m.rolled
		m.c.AbilityScore = score.Ability{}
		m.enterAssign()
		return nil
	}
	return m.beginSet()
}

// complete reports whether every roll has been assigned.
func (m rollModel) complete() bool { return !m.rollBoxVisible }

// swapSelect picks a score, swaps it with a second one, or deselects it.
func (m *rollModel) swapSelect() {
	switch m.swapFrom {
	case ability.None:
		m.swapFrom = m.selectedAbility
	case m.selectedAbility:
		m.swapFrom = ability.None // pressing enter on the picked score cancels
	default:
		i, j := m.swapFrom-1, m.selectedAbility-1
		m.c.AbilityScore[i], m.c.AbilityScore[j] = m.c.AbilityScore[j], m.c.AbilityScore[i]
		m.swapFrom = ability.None
	}
}

// --- Update ---

func (m rollModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case spinMsg:
		if m.phase == phaseDice && m.spinning && msg.gen == m.spinGen {
			m.dice[m.dieIndex] = dice.D6.Roll()
			return m, spinTick(m.spinGen)
		}

	case tea.KeyMsg:
		switch m.phase {
		case phaseIntro:
			return m.updateIntro(msg)
		case phaseDice:
			return m.updateDice(msg)
		case phaseAssign:
			return m.updateAssign(msg)
		}
	}
	return m, nil
}

func (m rollModel) updateIntro(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "left", "h":
		if m.introIndex > 0 {
			m.introIndex--
		}
	case "right", "l":
		if m.introIndex < numIntro-1 {
			m.introIndex++
		}
	case "enter":
		if m.introIndex == introRoll {
			cmd := m.beginManualRolling()
			return m, cmd
		}
		m.enterAssign() // auto-roll: scores are already rolled
	}
	return m, nil
}

func (m rollModel) updateDice(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() != "enter" {
		return m, nil
	}
	if m.spinning {
		cmd := m.stopDie()
		return m, cmd
	}
	cmd := m.nextSet()
	return m, cmd
}

func (m rollModel) updateAssign(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.focus {
	case winAbilitiesBox:
		switch msg.String() {
		case "enter":
			if m.complete() {
				m.swapSelect()
			} else {
				m.abilitySelect()
			}
		case "left", "h":
			m.abilityNavLeft()
		case "right", "l":
			m.abilityNavRight()
		case "down", "j":
			if m.complete() {
				m.swapFrom = ability.None
				m.focus = winConf
				m.confIndex = confConfirm
			}
		}
	case winConf:
		switch msg.String() {
		case "enter":
			if m.confIndex == confConfirm {
				return m, pager.SwitchTo(pager.Charsheet)
			}
			m.confReset()
		case "l", "right":
			m.confNavRight()
		case "h", "left":
			m.confNavLeft()
		case "up", "k":
			m.focus = winAbilitiesBox
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

	}
	return m, nil
}

// --- View ---

func (m rollModel) center(s string) string {
	return lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(s)
}

// withHelp pins the help line to the bottom row of the terminal.
func (m rollModel) withHelp(body, help string) string {
	gap := max(m.height-lipgloss.Height(body), 1)
	return body + strings.Repeat("\n", gap) + m.center(help)
}

func (m rollModel) View() string {
	var body, help string
	switch m.phase {
	case phaseIntro:
		body, help = m.introView()
	case phaseDice:
		body, help = m.diceView()
	default:
		body, help = m.assignView()
	}
	return m.withHelp(body, help)
}

func (m rollModel) introView() (string, string) {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Roll Ability Scores")

	text := "You will roll six scores. Each score is made by rolling four six-sided dice " +
		"and adding up the highest three.\n\n" +
		"Choose Roll to roll each die yourself: a die spins until you press enter to stop it. " +
		"Choose Auto-roll to have all six scores rolled for you.\n\n" +
		"Either way, you will then assign the six scores to your abilities."
	desc := m.center(style.TypesetBody(text, min(max(m.width-8, 30), 64)))

	opt := func(label string, selected bool) string {
		s := lipgloss.NewStyle().Padding(0, 2)
		if selected {
			return s.Inherit(style.Label).Render("[ " + label + " ]")
		}
		return s.Render("[ " + label + " ]")
	}
	options := m.center(lipgloss.JoinHorizontal(lipgloss.Center,
		opt("Roll", m.introIndex == introRoll),
		"   ",
		opt("Auto-roll", m.introIndex == introAuto),
	))

	body := lipgloss.JoinVertical(lipgloss.Center, title, " ", desc, "", options)
	help := style.TypesetHelp(
		style.Hint{Key: "←/→", Desc: "choose"},
		style.Hint{Key: "enter", Desc: "select"},
		style.Hint{Key: "esc", Desc: "back"},
	)
	return body, help
}

func (m rollModel) diceView() (string, string) {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Ability Points")
	progress := m.center(style.Label.Render(fmt.Sprintf("Rolling score %d of %d", m.rollCount+1, len(m.rolled))))

	active := -1
	if m.spinning {
		active = m.dieIndex
	}
	diceBox := m.center(style.TypesetDice(m.dice, active, m.dropped, true))
	rollsRow := m.center(style.TypesetRollsRow(&m.rolled, m.rollCount, false))

	var status, help string
	if m.spinning {
		status = style.Desc.Render(fmt.Sprintf("Press enter to stop die %d", m.dieIndex+1))
		help = style.TypesetHelp(
			style.Hint{Key: "enter", Desc: "stop die"},
			style.Hint{Key: "esc", Desc: "back"},
		)
	} else {
		status = fmt.Sprintf("Dropped the lowest die. Total: %s", style.Label.Render(fmt.Sprint(m.rolled[m.rollCount])))
		next := "next roll"
		if m.rollCount == len(m.rolled)-1 {
			next = "assign scores"
		}
		help = style.TypesetHelp(
			style.Hint{Key: "enter", Desc: next},
			style.Hint{Key: "esc", Desc: "back"},
		)
	}

	body := lipgloss.JoinVertical(lipgloss.Center,
		title, " ", progress, "", diceBox, "", m.center(status), "", rollsRow)
	return body, help
}

func (m rollModel) assignView() (string, string) {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Ability Points")

	introText := "Pick an ability, then pick a roll to assign to it. " +
		"Assigning over an existing score returns the old one to the pool."
	if m.complete() {
		introText = "All rolls assigned. Pick a score, then another, to swap them. " +
			"Confirm keeps these scores. Clear returns every score to the pool; rolls can't be rerolled."
	}
	intro := style.Desc.Width(m.width).Align(lipgloss.Center).Render(introText)

	abilityBox := m.center(style.TypesetScoresPicked(
		m.c, m.selectedAbility, m.swapFrom, m.focus == winAbilitiesBox))

	var bottomSection, help string

	if m.complete() {
		btn := func(label string, selected bool) string {
			s := lipgloss.NewStyle().Padding(0, 2)
			if selected {
				s = s.Inherit(style.Label)
			}
			return s.Render("[ " + label + " ]")
		}
		inMenu := m.focus == winConf
		menuRow := lipgloss.JoinHorizontal(lipgloss.Center,
			btn("Confirm", inMenu && m.confIndex == confConfirm),
			"   ",
			btn("Clear", inMenu && m.confIndex == confReset),
		)
		bottomSection = m.center(lipgloss.JoinVertical(lipgloss.Center, "Are you finished?", menuRow))

		if inMenu {
			help = style.TypesetHelp(
				style.Hint{Key: "←/→", Desc: "choose"},
				style.Hint{Key: "enter", Desc: "select"},
				style.Hint{Key: "↑", Desc: "back to scores"},
			)
		} else {
			enterDesc := "pick a score"
			switch {
			case m.swapFrom == m.selectedAbility:
				enterDesc = "deselect"
			case m.swapFrom != ability.None:
				enterDesc = "swap with " + ability.Get(m.swapFrom).Abbr
			}
			help = style.TypesetHelp(
				style.Hint{Key: "←/→", Desc: "choose score"},
				style.Hint{Key: "enter", Desc: enterDesc},
				style.Hint{Key: "↓", Desc: "menu"},
			)
		}
	} else if m.focus == winRollBox {
		bottomSection = m.center(style.TypesetRollsRow(&m.rolled, m.selectedRoll, true))
		help = style.TypesetHelp(
			style.Hint{Key: "←/→", Desc: "choose roll"},
			style.Hint{Key: "enter", Desc: "assign to " + ability.Get(m.selectedAbility).Abbr},
		)
	} else {
		bottomSection = m.center(style.TypesetRollsRow(&m.rolled, m.selectedRoll, false))
		help = style.TypesetHelp(
			style.Hint{Key: "←/→", Desc: "choose ability"},
			style.Hint{Key: "enter", Desc: "pick"},
		)
	}

	body := lipgloss.JoinVertical(lipgloss.Center, title, intro, " ", abilityBox, "", bottomSection)
	return body, help
}
