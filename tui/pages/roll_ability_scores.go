// Package pages/roll_ability_scores.go
package pages

import (
	"fmt"
	"time"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/score"
	"dndwiz/dice"
	"dndwiz/tui/modal"
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
	m.selectedRoll = -1
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

// hasAssignedScores reports whether any roll has been placed on an ability.
// A zero in the pool means that roll has moved onto the character.
func (m rollModel) hasAssignedScores() bool {
	if m.phase != phaseAssign {
		return false
	}
	for _, r := range m.rolled {
		if r == 0 {
			return true
		}
	}
	return false
}

// handleEsc leaves the screen, showing a notice first if scores were assigned.
// Assigned scores already live on the character, so nothing needs copying.
func (m rollModel) handleEsc() tea.Cmd {
	if !m.hasAssignedScores() {
		return pager.Up()
	}

	body := "Your assigned scores have been saved. " +
		"Choose Roll again to finish assigning the rest."
	if m.complete() {
		body = "Your ability scores have been saved. " +
			"Choose Roll again to swap them or start over."
	}

	return modal.Notice(modal.NoticeOptions{
		Title:   "Progress saved",
		Body:    body,
		OnClose: pager.Up(),
	})
}

// --- Update ---

func (m rollModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case pager.ResetViewMsg:
		if m.phase == phaseAssign {
			m.confReset() // same as pressing Clear
		}

	case spinMsg:
		if m.phase == phaseDice && m.spinning && msg.gen == m.spinGen {
			m.dice[m.dieIndex] = dice.D6.Roll()
			return m, spinTick(m.spinGen)
		}

	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.swapFrom = ability.None // drop any half-finished swap
			return m, m.handleEsc()
		}
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

func (m rollModel) View() string {
	var body string
	switch m.phase {
	case phaseIntro:
		body, _ = m.introView()
	case phaseDice:
		body, _ = m.diceView()
	default:
		body, _ = m.assignView()
	}
	return body
}

func (m rollModel) Help() []style.Hint {
	switch m.phase {
	case phaseIntro:
		_, h := m.introView()
		return h
	case phaseDice:
		_, h := m.diceView()
		return h
	}
	_, h := m.assignView()
	return h
}

func (m rollModel) introView() (string, []style.Hint) {
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
	help := []style.Hint{
		{Key: "←/→", Desc: "choose"},
		{Key: "enter", Desc: "select"},
		{Key: "esc", Desc: "back"},
		{Key: "?", Desc: "help"},
	}
	return body, help
}

func (m rollModel) diceView() (string, []style.Hint) {
	title := style.Title.
		Width(m.width).
		Align(lipgloss.Center).
		Render("Ability Points")

	progress := m.
		center(
			style.Label.
				Render(
					fmt.Sprintf("Rolling score %d of %d", m.rollCount+1, len(m.rolled)),
				),
		)

	active := -1
	if m.spinning {
		active = m.dieIndex
	}
	diceBox := m.center(style.TypesetDice(m.dice, active, m.dropped, true))
	rollsRow := m.center(style.TypesetRollsRow(&m.rolled, m.rollCount, false))

	var status string
	var help []style.Hint
	if m.spinning {
		status = style.Desc.Render(fmt.Sprintf("Press enter to stop die %d", m.dieIndex+1))
		help = []style.Hint{
			{Key: "enter", Desc: "stop die"},
			{Key: "esc", Desc: "back"},
			{Key: "?", Desc: "help"},
		}
	} else {
		status = fmt.Sprintf("Dropped the lowest die. Total: %s", style.Label.Render(fmt.Sprint(m.rolled[m.rollCount])))
		next := "next roll"
		if m.rollCount == len(m.rolled)-1 {
			next = "assign scores"
		}
		help = []style.Hint{
			{Key: "enter", Desc: next},
			{Key: "esc", Desc: "back"},
			{Key: "?", Desc: "help"},
		}
	}

	body := lipgloss.JoinVertical(lipgloss.Center,
		title, " ", progress, "", diceBox, "", m.center(status), "", rollsRow)
	return body, help
}

func (m rollModel) assignView() (string, []style.Hint) {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Ability Points")

	introText := "Pick an ability, then pick a roll to assign to it. " +
		"Assigning over an existing score returns the old one to the pool."
	if m.complete() {
		introText = "All rolls assigned. Pick a score, then another, to swap them. " +
			"Confirm keeps these scores. Clear returns every score to the pool; rolls can't be rerolled."
	}
	intro := style.Desc.Width(m.width).Align(lipgloss.Center).Render(introText)

	// Roll box: always visible, above the scores. Highlighted only when focused.
	rollBox := m.center(style.TypesetRollsRow(&m.rolled, m.selectedRoll, m.focus == winRollBox))

	abilityBox := m.center(style.TypesetScoresPicked(
		m.c, m.selectedAbility, m.swapFrom, m.focus == winAbilitiesBox))

	info := ability.Get(m.selectedAbility)
	detailW := min(max(m.width-8, 30), 64)
	detail := m.center(lipgloss.JoinVertical(lipgloss.Center,
		style.Label.Render(info.Name),
		style.TypesetDesc(info.Desc, detailW, lipgloss.Center),
	))

	// Confirm/Clear menu: always visible, dimmed until it can be used.
	enabled := m.complete()
	inMenu := enabled && m.focus == winConf
	btn := func(label string, selected bool) string {
		s := lipgloss.NewStyle().Padding(0, 2)
		switch {
		case !enabled:
			s = s.Foreground(lipgloss.Color("240")) // grayed out
		case selected:
			s = s.Inherit(style.Label)
		}
		return s.Render("[ " + label + " ]")
	}
	menu := m.center(lipgloss.JoinHorizontal(lipgloss.Center,
		btn("Confirm", inMenu && m.confIndex == confConfirm),
		"   ",
		btn("Clear", inMenu && m.confIndex == confReset),
	))

	var help []style.Hint
	switch {
	case inMenu:
		help = []style.Hint{
			{Key: "←/→", Desc: "choose"},
			{Key: "enter", Desc: "select"},
			{Key: "↑", Desc: "back to scores"},
			{Key: "?", Desc: "help"},
		}
	case enabled:
		enterDesc := "pick a score"
		switch {
		case m.swapFrom == m.selectedAbility:
			enterDesc = "deselect"
		case m.swapFrom != ability.None:
			enterDesc = "swap with " + ability.Get(m.swapFrom).Abbr
		}
		help = []style.Hint{
			{Key: "←/→", Desc: "choose score"},
			{Key: "enter", Desc: enterDesc},
			{Key: "↓", Desc: "menu"},
			{Key: "?", Desc: "help"},
		}
	case m.focus == winRollBox:
		help = []style.Hint{
			{Key: "←/→", Desc: "choose roll"},
			{Key: "enter", Desc: "assign to " + ability.Get(m.selectedAbility).Abbr},
			{Key: "?", Desc: "help"},
		}
	default:
		help = []style.Hint{
			{Key: "←/→", Desc: "choose ability"},
			{Key: "enter", Desc: "pick"},
			{Key: "?", Desc: "help"},
		}
	}

	body := lipgloss.JoinVertical(lipgloss.Center,
		title, intro, " ", rollBox, abilityBox, detail, "", menu)
	return body, help
}

// HelpPage implements the "?" help for the roll screen. The content depends on
// the current phase, and on whether every roll has been assigned.
func (m rollModel) HelpPage() (string, func(int) string, bool) {
	switch m.phase {
	case phaseIntro:
		return "Rolling Help", m.introHelp, true
	case phaseDice:
		return "Rolling Help", m.diceHelp, true
	}
	return "Assigning Help", m.assignHelp, true
}

func helpP(w int) func(string) string {
	return func(s string) string { return style.TypesetBody(s, w) }
}

func helpTerm(w int, name, text string) string {
	return lipgloss.JoinVertical(lipgloss.Left,
		style.Label.Render(name),
		style.TypesetBody(text, w),
		"",
	)
}

// rollBasics explains what rolling is. It's shared by the intro and dice help.
func rollBasics(w int) string {
	p := helpP(w)
	return lipgloss.JoinVertical(lipgloss.Left,
		style.Header.Render("What is rolling?"),
		p("Ability scores are six numbers that describe your character's natural talents. "+
			"Rolling is the traditional way to get them: you throw dice and let chance "+
			"decide. It is the most surprising of the three methods."),
		"",
		p("To make one score, you roll four six-sided dice, throw away the lowest one, and "+
			"add up the other three. That gives a number from 3 to 18, and most results land "+
			"in the middle. You do this six times to get six scores."),
		"",
		p("You get only one set of six, and they cannot be rerolled. Once you reach the "+
			"assigning step, those are your numbers for the rest of this character. You can "+
			"still switch to Point Buy or Standard Array from the method screen, but your "+
			"rolled set stays available if you come back."),
		style.TypesetDivider(w),
	)
}

func (m rollModel) introHelp(w int) string {
	p := helpP(w)
	return lipgloss.JoinVertical(lipgloss.Left,
		rollBasics(w),

		style.Header.Render("Roll or Auto-roll?"),
		helpTerm(w, "Roll",
			"You stop each die yourself. A die spins on screen, and you press Enter to "+
				"stop it. It is just for fun, since the result is still random. Pick this if "+
				"you like the feeling of rolling dice."),
		helpTerm(w, "Auto-roll",
			"Skips the dice and goes straight to assigning. Your six scores are already "+
				"rolled, so nothing is lost, and the result is just as random. Pick this if "+
				"you just want to get on with it."),
		style.TypesetDivider(w),

		style.Header.Render("How to use this screen"),
		p("Press ←/→ (or h/l) to choose Roll or Auto-roll, then press Enter. Esc takes you "+
			"back to the method screen. Nothing is lost by leaving at this stage."),
		"",
		style.Desc.Render("Press Esc or ? to close."),
	)
}

func (m rollModel) diceHelp(w int) string {
	p := helpP(w)
	return lipgloss.JoinVertical(lipgloss.Left,
		rollBasics(w),

		style.Header.Render("How to use this screen"),
		p("The screen shows four dice, labeled d1 to d4, and one of them spins. Press Enter "+
			"to stop it, and the next one starts spinning. After the fourth die stops, the "+
			"lowest is struck through and the other three are added. That total is one of "+
			"your six scores."),
		"",
		p("The title shows which score you are on, for example \"Rolling score 2 of 6.\" The "+
			"row at the bottom collects your results as you go. Press Enter again to start "+
			"the next score. After the sixth, you move on to assigning."),
		"",
		p("Esc takes you back to the method screen. Your six scores are not lost if you do."),
		style.TypesetDivider(w),

		style.Header.Render("A note on fairness"),
		p("It does not matter when you press Enter. The dice are random, so you cannot "+
			"aim for a high number. Spend as long as you like, or none at all."),
		"",
		style.Desc.Render("Press Esc or ? to close."),
	)
}

func (m rollModel) assignHelp(w int) string {
	p := helpP(w)

	body := []string{
		style.Header.Render("What is this step?"),
		p("You now have six numbers, but they do not belong to any ability yet. Here you " +
			"decide which ability gets which number. This is the most important choice " +
			"you make with rolled scores, because the same six numbers can make very " +
			"different characters."),
		"",
		p("A good rule of thumb is to put your highest numbers in the abilities your class " +
			"depends on most. The character sheet lists them as Primary Abilities, so choose " +
			"a class first if you want to know. Constitution helps everyone, since it adds " +
			"hit points."),
		style.TypesetDivider(w),
	}

	if m.complete() {
		body = append(body,
			style.Header.Render("All rolls are assigned"),
			p("Every score now belongs to an ability, and you can keep rearranging them before "+
				"you finish."),
			"",
			helpTerm(w, "Swap two scores",
				"With the cursor on the score box, press Enter on one score to pick it up. It "+
					"turns into a filled block. Move to another score with ←/→ and press Enter "+
					"again, and the two exchange places. Press Enter on the picked score "+
					"itself to put it back without changing anything."),
			helpTerm(w, "Go to the menu",
				"Press ↓ (or j) to reach Confirm and Clear. Any score you had picked up is "+
					"put back first. Press ↑ (or k) to return to the scores."),
			helpTerm(w, "Confirm",
				"Keeps these scores and returns to the character sheet."),
			helpTerm(w, "Clear",
				"Takes every score off the abilities and returns all six to the pool so you "+
					"can assign them from scratch. This does not reroll them. It is the only "+
					"action here that undoes your arrangement."),
		)
	} else {
		body = append(body,
			style.Header.Render("How to use this screen"),
			helpTerm(w, "1. Pick an ability",
				"Use ←/→ (or h/l) to move the highlight along the six abilities, then press "+
					"Enter. The description under the scores tells you what the highlighted "+
					"ability does."),
			helpTerm(w, "2. Pick a roll",
				"Focus moves to the row of unassigned rolls. Use ←/→ to choose one and press "+
					"Enter to assign it. It leaves the pool and appears on the ability."),
			helpTerm(w, "3. Repeat",
				"Continue until all six are assigned. Then the row of rolls disappears and "+
					"you can fine-tune the arrangement."),
			helpTerm(w, "Changing your mind",
				"Assign a different roll to an ability that already has one. The old number "+
					"goes back to the pool, so nothing is lost or duplicated."),
		)
	}

	body = append(body,
		style.TypesetDivider(w),
		style.Header.Render("Reading the scores"),
		helpTerm(w, "Score and modifier",
			"The big number is the score. The smaller number under it is the modifier, the "+
				"bonus or penalty it adds to your rolls. A score of 10 or 11 gives +0, 15 "+
				"gives +2, and 8 gives -1."),
		style.TypesetDivider(w),

		style.Header.Render("Leaving this screen"),
		p("Esc saves the scores you have assigned and shows a notice before returning to the "+
			"method screen. Choose Roll again to continue where you left off. If you pick "+
			"Point Buy or Standard Array in between, those methods replace your assigned "+
			"scores, but choosing Roll afterward gives you back the same six original rolls, "+
			"unassigned."),
		"",
		style.Desc.Render("Press Esc or ? to close."),
	)

	return lipgloss.JoinVertical(lipgloss.Left, body...)
}
