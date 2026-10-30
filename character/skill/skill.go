// Package skill
package skill

type Skill int

const (
	None Skill = iota
	Acrobatics
	AnimalHealing
	Arcana
	Athletics
	Deception
	History
	Insight
	Intimidation
	Investigation
	Medicine
	Nature
	Perception
	Performance
	Persuasion
	Religion
	SleightOfHand
	Stealth
	Survival
)

var All = []Skill{
	Acrobatics,
	AnimalHealing,
	Arcana,
	Athletics,
	Deception,
	History,
	Insight,
	Intimidation,
	Investigation,
	Medicine,
	Nature,
	Perception,
	Performance,
	Persuasion,
	Religion,
	SleightOfHand,
	Stealth,
	Survival,
}

type Info struct {
	Name string
}

var skillInfo = [...]Info{
	None: {
		Name: "None",
	},
	Acrobatics: {
		Name: "Acrobatics",
	},
	AnimalHealing: {
		Name: "Animal Healing",
	},
	Arcana: {
		Name: "Arcana",
	},
	Athletics: {
		Name: "Athletics",
	},
	Deception: {
		Name: "Deception",
	},
	History: {
		Name: "History",
	},
	Insight: {
		Name: "Insight",
	},
	Intimidation: {
		Name: "Intimidation",
	},
	Investigation: {
		Name: "Investigation",
	},
	Medicine: {
		Name: "Medicine",
	},
	Nature: {
		Name: "Nature",
	},
	Perception: {
		Name: "Perception",
	},
	Performance: {
		Name: "Performance",
	},
	Persuasion: {
		Name: "Persuasion",
	},
	Religion: {
		Name: "Religion",
	},
	SleightOfHand: {
		Name: "Sleight of Hand",
	},
	Stealth: {
		Name: "Stealth",
	},
	Survival: {
		Name: "Survival",
	},
}

func Get(s Skill) Info {
	return skillInfo[s]
}

func GetAll() []Info {
	return skillInfo[1:]
}
