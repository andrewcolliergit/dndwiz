package character

import (
	"errors"
	"strconv"
	"strings"
)

type (
	AbilityScores               [6]int // Str, Dex, Con, Int, Wis, Cha
	AbilityScoreSelectionMethod int
)

const (
	AbilityScoreMethodNone = iota
	AbilityScoreMethodStdArray
	AbilityScoreMethodPointBuy
	AbilityScoreMethodRoll
)

func (sm AbilityScoreSelectionMethod) String() string {
	switch sm {
	case AbilityScoreMethodNone:
		return "None"
	case AbilityScoreMethodPointBuy:
		return "Point Buy"
	case AbilityScoreMethodStdArray:
		return "Standard Array"
	case AbilityScoreMethodRoll:
		return "Rolled"
	default:
		return "Invalid Ability Score Method"
	}
}

func (as AbilityScores) String() string {
	var arr []string

	for i, a := range GetAllAbilities() {
		score := strconv.Itoa(as[i])
		arr = append(arr, a.Abbr+" "+score)
	}
	result := strings.Join(arr, ", ")
	return "[" + result + "]"
}

func AbilityScorePointCost(score int) (int, error) {
	if score < 9 {
		return 0, errors.New("invalid score")
	} else if score > 14 {
		return 9, nil
	} else {
		return score - 8, nil
	}
}

func AbilityModifier(score int) (int, error) {
	if score < 3 || score > 20 {
		return 0, errors.New("invalid score")
	}
	return (score - 10) / 2, nil
}
