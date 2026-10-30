// Package ability
package ability

import (
	"strings"

	"dndwiz/character/skill"
)

type (
	Ability   int
	Abilities []Ability
	Skills    []skill.Skill
)

const (
	None Ability = iota
	Strength
	Dexterity
	Constitution
	Intelligence
	Wisdom
	Charisma
)

type AbilityInfo struct {
	Name  string
	Abbr  string
	Desc  string
	Skill Skills
}

var abilityInfo = [...]AbilityInfo{
	Strength: {
		Name:  "Strength",
		Abbr:  "STR",
		Desc:  "Measures physical power, lifting, and carrying.",
		Skill: Skills{skill.Athletics},
	},
	Dexterity: {
		Name:  "Dexterity",
		Abbr:  "DEX",
		Desc:  "Measures reflexes, coordination, and stealth.",
		Skill: Skills{skill.Acrobatics, skill.SleightOfHand, skill.Stealth},
	},
	Constitution: {
		Name:  "Constitution",
		Abbr:  "CON",
		Desc:  "Measures health, stamina, and hit point pools.",
		Skill: nil,
	},
	Intelligence: {
		Name:  "Intelligence",
		Abbr:  "INT",
		Desc:  "Measures mental acuity, recall, and analytical skill.",
		Skill: Skills{skill.Arcana, skill.History, skill.Investigation, skill.Nature, skill.Religion},
	},
	Wisdom: {
		Name:  "Wisdom",
		Abbr:  "WIS",
		Desc:  "Measures awareness, intuition, and mental fortitude.",
		Skill: Skills{skill.AnimalHealing, skill.Insight, skill.Medicine, skill.Perception, skill.Survival},
	},
	Charisma: {
		Name:  "Charisma",
		Abbr:  "CHA",
		Desc:  "Measures force of personality, persuasion, and leadership.",
		Skill: Skills{skill.Deception, skill.Intimidation, skill.Performance, skill.Persuasion},
	},
}

func Get(a Ability) AbilityInfo {
	return abilityInfo[a]
}

func GetAll() []AbilityInfo {
	return abilityInfo[1:]
}

// Class Primary Abilities

func (p Abilities) String() string {
	var str []string
	for _, a := range p {
		str = append(str, Get(a).Abbr)
	}
	return strings.Join(str, ", ")
}
