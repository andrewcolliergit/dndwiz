// Package background
package background

import (
	"strings"

	"dndwiz/character/ability"
	"dndwiz/character/feat"
	"dndwiz/character/skill"
)

type Background int

const (
	None Background = iota
	Acolyte
	Criminal
	Sage
	Soldier
)

type Info struct {
	Name      string
	Abilities [3]ability.Ability
	Skills    [2]skill.Skill
	Feat      feat.Feat
	Tool      string
	Equipment [2]string
	Desc      string
}

var info = [...]Info{
	None: {
		Name: "None",
	},
	Acolyte: {
		Name:      "Acolyte",
		Abilities: [3]ability.Ability{ability.Intelligence, ability.Wisdom, ability.Charisma},
		Skills:    [2]skill.Skill{skill.Insight, skill.Religion},
		Feat:      feat.MagicInitiate,
		Tool:      "Caligrapher's Supplies",
		Equipment: [2]string{"Caligrapher's Supplies, Book (prayers), Holy Symbol, Parchment (10 sheets), Robe, 8 GP", "50 GP"},
		Desc:      "You spent your youth in a temple, serving a deity or faith. You acted as an intermediary between the sacred realm and the mortal world, performing rites and offering sacrifices to conduct worship.",
	},
	Criminal: {
		Name:      "Criminal",
		Abilities: [3]ability.Ability{ability.Dexterity, ability.Constitution, ability.Intelligence},
		Skills:    [2]skill.Skill{skill.SleightOfHand, skill.Stealth},
		Feat:      feat.Alert,
		Tool:      "Theives' Tools",
		Equipment: [2]string{"2 Daggers, Theives' Tools, Crowbar, 2 Pouches, Traveler's Clothes, 16 GP", "50 GP"},
		Desc:      "You learned to survive in the underbelly of society, operating outside the boundaries of the law. Whether picking pockets, cracking safes, or smuggling contraband, stealth and street sense are your lifelines.",
	},
	Sage: {
		Name:      "Sage",
		Abilities: [3]ability.Ability{ability.Constitution, ability.Intelligence, ability.Wisdom},
		Skills:    [2]skill.Skill{skill.Arcana, skill.History},
		Feat:      feat.MagicInitiate,
		Tool:      "Caligrapher's Supplies",
		Equipment: [2]string{"Quarterstaff, Caligrapher's Supplies, Book (history), Parchment (8 sheets), Robe, 8 GP", "50 GP"},
		Desc:      "You spent years studying manuscripts, scrolls, and ancient lore in academies or monasteries. Your sharp mind hungers for knowledge, seeking answers to the multiverse's deepest mysteries.",
	},
	Soldier: {
		Name:      "Soldier",
		Abilities: [3]ability.Ability{ability.Strength, ability.Dexterity, ability.Constitution},
		Skills:    [2]skill.Skill{skill.Athletics, skill.Intimidation},
		Feat:      feat.SavageAttacker,
		Tool:      "Gaming Set",
		Equipment: [2]string{"Spear, Shortbow, 20 Arrows, Gaming Set, Healer's Kit, Quiver, Traveler's Clothes, 14 GP", "50 GP"},
		Desc:      "You served in an army, mercenary company, or local militia, learning discipline and martial tactics. You know the brutality of combat, the weight of armor, and the value of standing alongside allies.",
	},
}

func Get(bg Background) Info {
	return info[bg]
}

func GetAll() []Info {
	return info[1:]
}

func GetBackgroundAbilities(bg Background) string {
	bgi := info[bg]
	var abilities []string
	for _, a := range bgi.Abilities {
		ai := ability.Get(a)
		abilities = append(abilities, ai.Name)
	}

	return strings.Join(abilities, ", ")
}

func EquipmentOptions(bg Background) string {
	bgi := info[bg]
	return "(A) " + bgi.Equipment[0] + ", or (B) " + bgi.Equipment[1]
}

func GetSkills(bg Background) string {
	bgi := info[bg]
	var skills []string
	for _, s := range bgi.Skills {
		skills = append(skills, skill.Get(s).Name)
	}

	return strings.Join(skills, ", ")
}
