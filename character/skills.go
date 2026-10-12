package character

type Skill int

const (
	SkillNone Skill = iota
	SkillAcrobatics
	SkillAnimalHealing
	SkillArcana
	SkillAthletics
	SkillDeception
	SkillHistory
	SkillInsight
	SkillIntimidation
	SkillInvestigation
	SkillMedicine
	SkillNature
	SkillPerception
	SkillPerformance
	SkillPersuasion
	SkillReligion
	SkillSleightOfHand
	SkillStealth
	SkillSurvival
)

var AllSkills = []Skill{
	SkillAcrobatics,
	SkillAnimalHealing,
	SkillArcana,
	SkillAthletics,
	SkillDeception,
	SkillHistory,
	SkillInsight,
	SkillIntimidation,
	SkillInvestigation,
	SkillMedicine,
	SkillNature,
	SkillPerception,
	SkillPerformance,
	SkillPersuasion,
	SkillReligion,
	SkillSleightOfHand,
	SkillStealth,
	SkillSurvival,
}

type SkillInfo struct {
	Name string
}

var skillInfo = [...]SkillInfo{
	SkillNone: {
		Name: "None",
	},
	SkillAcrobatics: {
		Name: "Acrobatics",
	},
	SkillAnimalHealing: {
		Name: "Animal Healing",
	},
	SkillArcana: {
		Name: "Arcana",
	},
	SkillAthletics: {
		Name: "Athletics",
	},
	SkillDeception: {
		Name: "Deception",
	},
	SkillHistory: {
		Name: "History",
	},
	SkillInsight: {
		Name: "Insight",
	},
	SkillIntimidation: {
		Name: "Intimidation",
	},
	SkillInvestigation: {
		Name: "Investigation",
	},
	SkillMedicine: {
		Name: "Medicine",
	},
	SkillNature: {
		Name: "Nature",
	},
	SkillPerception: {
		Name: "Perception",
	},
	SkillPerformance: {
		Name: "Performance",
	},
	SkillPersuasion: {
		Name: "Persuasion",
	},
	SkillReligion: {
		Name: "Religion",
	},
	SkillSleightOfHand: {
		Name: "Sleight of Hand",
	},
	SkillStealth: {
		Name: "Stealth",
	},
	SkillSurvival: {
		Name: "Survival",
	},
}

func GetSkill(s Skill) SkillInfo {
	return skillInfo[s]
}

func GetAllSkills() []SkillInfo {
	return skillInfo[1:]
}
