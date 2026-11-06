// Package score/score.go
package score

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"dndwiz/character/ability"
)

type (
	Ability [6]int // Str, Dex, Con, Int, Wis, Cha
	Method  int
)

const (
	None = iota
	StdArray
	PointBuy
	Roll
)

func (sm Method) String() string {
	switch sm {
	case None:
		return "None"
	case PointBuy:
		return "Point Buy"
	case StdArray:
		return "Standard Array"
	case Roll:
		return "Rolled"
	default:
		return "Invalid Ability Score Method"
	}
}

func (as Ability) String() string {
	var arr []string

	for i, a := range ability.GetAll() {
		score := strconv.Itoa(as[i])
		arr = append(arr, a.Abbr+" "+score)
	}
	result := strings.Join(arr, ", ")
	return "[" + result + "]"
}

func Cost(score int) (int, error) {
	if score < 9 {
		return 0, errors.New("invalid score")
	} else if score > 14 {
		return 9, nil
	} else {
		return score - 8, nil
	}
}

func Modifier(score int) (int, error) {
	if score < 3 || score > 20 {
		return 0, errors.New("invalid score")
	}
	return int(math.Floor(float64(score-10) / 2)), nil
}

// ModifierString formats a score's modifier as "+2", "+0", or "-1".
// It returns "" for unassigned or invalid scores.
func ModifierString(score int) string {
	m, err := Modifier(score)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%+d", m)
}
