// Package pages/buy_ability_scores.go
package pages

import (
	"fmt"

	"dndwiz/character"
	"dndwiz/character/ability"
	"dndwiz/character/score"
	"dndwiz/tui/pager"
	"dndwiz/tui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func getMinAbilityPoints() score.Ability {
	return score.Ability{8, 8, 8, 8, 8, 8}
}

const abilityPointBuyBudget = 27

type buyModel struct {
	c               *character.Character
	points          int
	selectedAbility ability.Ability
	confVisible     bool
	confIndex       confSel
	width           int
	height          int
}

func NewBuyModel(c *character.Character) tea.Model {
	c.AbilityScore = getMinAbilityPoints()
	return buyModel{
		c:               c,
		points:          abilityPointBuyBudget,
		selectedAbility: ability.Strength,
		confVisible:     false,
		confIndex:       confConfirm,
	}
}

func pointCost(score int) int {
	switch score {
	case 8:
		return 0
	case 9:
		return 1
	case 10:
		return 2
	case 11:
		return 3
	case 12:
		return 4
	case 13:
		return 5
	case 14:
		return 7
	case 15:
		return 9
	}
	return 0
}

func (m *buyModel) abilityIncrease() {
	score := m.c.AbilityScore[m.selectedAbility-1]
	if score >= 15 {
		return
	}

	cost := pointCost(score + 1)

	if m.points >= cost {
		m.points -= cost
		m.c.AbilityScore[m.selectedAbility-1]++
	}

	if m.points == 0 {
		m.confVisible = true
	}
}

func (m *buyModel) abilityDecrease() {
	score := m.c.AbilityScore[m.selectedAbility-1]
	if score <= 8 {
		return
	}

	refund := pointCost(score)
	m.points += refund
	m.c.AbilityScore[m.selectedAbility-1]--

	if m.points > 0 {
		m.confVisible = false
	}
}

func (m buyModel) Init() tea.Cmd {
	return nil
}

func (m buyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			if m.selectedAbility > ability.Strength {
				m.selectedAbility--
			}
		case "right", "l":
			if m.selectedAbility < ability.Charisma {
				m.selectedAbility++
			}
		case "up", "+", "k":
			m.abilityIncrease()
		case "down", "-", "j":
			m.abilityDecrease()
		case "enter":
			if m.points == 0 {
				return m, pager.SwitchTo(pager.Charsheet)
			}
		}
	}
	return m, nil
}

func (m buyModel) View() string {
	title := style.Title.Width(m.width).Align(lipgloss.Center).Render("Ability Points")
	points := fmt.Sprintf("Points: %d", m.points)
	abilityBox := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).Render(style.TypesetScoresPicked(m.c, m.selectedAbility, m.selectedAbility, m.points > 0))

	var bottomSection string
	if m.points == 0 {
		bottomSection = style.Label.Render("Press enter to continue")
	}
	return lipgloss.JoinVertical(lipgloss.Center, title, " ", points, abilityBox, "", bottomSection)
}

func (m buyModel) Help() []style.Hint {
	hints := []style.Hint{{Key: "←/→", Desc: "choose ability"}, {Key: "↑/↓", Desc: "buy/refund"}}
	if m.points == 0 {
		hints = append(hints, style.Hint{Key: "enter", Desc: "continue"})
	}
	return append(hints, style.Hint{Key: "esc", Desc: "back"})
}

func (m buyModel) HelpPage() (string, string, bool) {
	return "Point Buy Help", "You have 27 points to spend. Every score starts at 8 and can go up to 15. Higher scores cost more: 14 costs 7 points and 15 costs 9.\n\n" +
		"Use ←/→ to choose an ability and ↑/↓ to buy or refund points. Press Enter when you've spent all 27.", true
}
