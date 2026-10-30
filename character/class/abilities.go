package class

import (
	"strings"

	"dndwiz/character/ability"
)

// Class Abilities
type (
	Abilities struct {
		Has  int
		From ability.Abilities
	}
)

func (cpa Abilities) HasOption() bool {
	return cpa.Has < len(cpa.From)
}

func (cpa Abilities) String() string {
	var str []string
	delim := " and "

	if cpa.HasOption() {
		delim = " or "
	}

	for _, a := range cpa.From {
		str = append(str, ability.Get(a).Name)
	}
	return strings.Join(str, delim)
}

func (cpa Abilities) RequiresChoice() bool {
	return cpa.Has < len(cpa.From)
}
