package character

import "strings"

type (
	Ability        int
	Abilities      []Ability
	AbilitySkills  []Skill
	ClassAbilities struct {
		Has  int
		From Abilities
	}
)

const (
	AbilityNone Ability = iota
	AbilityStrength
	AbilityDexterity
	AbilityConstitution
	AbilityIntelligence
	AbilityWisdom
	AbilityCharisma
)

type AbilityInfo struct {
	Name  string
	Abbr  string
	Desc  string
	Skill AbilitySkills
}

var abilityInfo = [...]AbilityInfo{
	AbilityStrength: {
		Name:  "Strength",
		Abbr:  "STR",
		Desc:  "Measures physical power, lifting, and carrying.",
		Skill: AbilitySkills{SkillAthletics},
	},
	AbilityDexterity: {
		Name:  "Dexterity",
		Abbr:  "DEX",
		Desc:  "Measures reflexes, coordination, and stealth.",
		Skill: AbilitySkills{SkillAcrobatics, SkillSleightOfHand, SkillStealth},
	},
	AbilityConstitution: {
		Name:  "Constitution",
		Abbr:  "CON",
		Desc:  "Measures health, stamina, and hit point pools.",
		Skill: nil,
	},
	AbilityIntelligence: {
		Name:  "Intelligence",
		Abbr:  "INT",
		Desc:  "Measures mental acuity, recall, and analytical skill.",
		Skill: AbilitySkills{SkillArcana, SkillHistory, SkillInvestigation, SkillNature, SkillReligion},
	},
	AbilityWisdom: {
		Name:  "Wisdom",
		Abbr:  "WIS",
		Desc:  "Measures awareness, intuition, and mental fortitude.",
		Skill: AbilitySkills{SkillAnimalHealing, SkillInsight, SkillMedicine, SkillPerception, SkillSurvival},
	},
	AbilityCharisma: {
		Name:  "Charisma",
		Abbr:  "CHA",
		Desc:  "Measures force of personality, persuasion, and leadership.",
		Skill: AbilitySkills{SkillDeception, SkillIntimidation, SkillPerformance, SkillPersuasion},
	},
}

func GetAbility(a Ability) AbilityInfo {
	return abilityInfo[a]
}

func GetAllAbilities() []AbilityInfo {
	return abilityInfo[1:]
}

// Class Primary Abilities

func (p Abilities) String() string {
	var str []string
	for _, a := range p {
		str = append(str, GetAbility(a).Name)
	}
	return strings.Join(str, ", ")
}

func (cpa ClassAbilities) HasOption() bool {
	return cpa.Has < len(cpa.From)
}

func (cpa ClassAbilities) String() string {
	var str []string
	delim := " and "

	if cpa.HasOption() {
		delim = " or "
	}

	for _, a := range cpa.From {
		str = append(str, GetAbility(a).Name)
	}
	return strings.Join(str, delim)
}

func (cpa ClassAbilities) RequiresChoice() bool {
	return cpa.Has < len(cpa.From)
}
