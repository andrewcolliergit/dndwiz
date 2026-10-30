package class

import (
	"fmt"
	"strings"

	"dndwiz/character/ability"
	"dndwiz/character/skill"
)

// Class Skill Proficiencies
type (
	Skills           []skill.Skill
	SkillProficiency struct {
		Has  int
		From Skills
	}
)

func (csp SkillProficiency) String() string {
	prefix := fmt.Sprintf("Choose %d: ", csp.Has)
	var str []string
	for _, s := range csp.From {
		str = append(str, skill.Get(s).Name)
	}

	return prefix + strings.Join(str, ", ")
}

func (csp SkillProficiency) RequiresChoice() bool {
	return csp.Has < len(csp.From)
}

type SavingThrow []ability.Ability

func (st SavingThrow) String() string {
	var str []string
	for _, a := range st {
		str = append(str, ability.Get(a).Name)
	}
	return strings.Join(str, ", ")
}
