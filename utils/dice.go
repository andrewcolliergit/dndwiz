// Package utils/dice.go
package utils

import "math/rand/v2"

type Die int

const (
	D6 Die = iota
	D8
	D10
	D12
	D20
)

func (d Die) String() string {
	switch d {
	case D6:
		return "D6"
	case D8:
		return "D8"
	case D10:
		return "D10"
	case D12:
		return "D12"
	case D20:
		return "D20"
	default:
		return "Invalid"
	}
}

func (d Die) Roll() int {
	var sides int
	switch d {
	case D6:
		sides = 6
	case D8:
		sides = 8
	case D10:
		sides = 10
	case D12:
		sides = 12
	case D20:
		sides = 20
	default:
		return 0
	}
	return rand.IntN(sides) + 1
}

type Roller interface {
	Roll(sides int) (int, error)
}

// Roll return a single die roll between 1 and sides, inclusive.
func Roll(sides int) int {
	return rand.IntN(sides) + 1
}

// RollAbilityScore rolls 4d6, drops the lowest die, and returns
// all four dice plus the final score.
func RollAbilityScore() (dice [4]int, score int) {
	lowest := 7
	for i := range dice {
		dice[i] = Roll(6)
		score += dice[i]
		if dice[i] < lowest {
			lowest = dice[i]
		}
	}
	return dice, score - lowest
}

func RollAllAbilityScores() [6]int {
	var scores [6]int

	for i := range 6 {
		_, roll := RollAbilityScore()
		scores[i] = roll
	}
	return scores
}
