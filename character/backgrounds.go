package character

import "strings"

type Background int

const (
	BackgroundNone Background = iota
	BackgroundAcolyte
	BackgroundCriminal
	BackgroundSage
	BackgroundSoldier
)

type BackgroundInfo struct {
	Name      string
	Abilities [3]Ability
	Skills    [2]Skill
	Feat      OriginFeat
	Tool      string
	Equipment [2]string
	Desc      string
}

var backgroundInfo = [...]BackgroundInfo{
	BackgroundNone: {
		Name: "None",
	},
	BackgroundAcolyte: {
		Name:      "Acolyte",
		Abilities: [3]Ability{AbilityIntelligence, AbilityWisdom, AbilityCharisma},
		Skills:    [2]Skill{SkillInsight, SkillReligion},
		Feat:      FeatMagicInitiate,
		Tool:      "Caligrapher's Supplies",
		Equipment: [2]string{"Caligrapher's Supplies, Book (prayers), Holy Symbol, Parchment (10 sheets), Robe, 8 GP", "50 GP"},
		Desc:      "You spent your youth in a temple, serving a deity or faith. You acted as an intermediary between the sacred realm and the mortal world, performing rites and offering sacrifices to conduct worship.",
	},
	BackgroundCriminal: {
		Name:      "Criminal",
		Abilities: [3]Ability{AbilityDexterity, AbilityConstitution, AbilityIntelligence},
		Skills:    [2]Skill{SkillSleightOfHand, SkillStealth},
		Feat:      FeatAlert,
		Tool:      "Theives' Tools",
		Equipment: [2]string{"2 Daggers, Theives' Tools, Crowbar, 2 Pouches, Traveler's Clothes, 16 GP", "50 GP"},
		Desc:      "You learned to survive in the underbelly of society, operating outside the boundaries of the law. Whether picking pockets, cracking safes, or smuggling contraband, stealth and street sense are your lifelines.",
	},
	BackgroundSage: {
		Name:      "Sage",
		Abilities: [3]Ability{AbilityConstitution, AbilityIntelligence, AbilityWisdom},
		Skills:    [2]Skill{SkillArcana, SkillHistory},
		Feat:      FeatMagicInitiate,
		Tool:      "Caligrapher's Supplies",
		Equipment: [2]string{"Quarterstaff, Caligrapher's Supplies, Book (history), Parchment (8 sheets), Robe, 8 GP", "50 GP"},
		Desc:      "You spent years studying manuscripts, scrolls, and ancient lore in academies or monasteries. Your sharp mind hungers for knowledge, seeking answers to the multiverse's deepest mysteries.",
	},
	BackgroundSoldier: {
		Name:      "Soldier",
		Abilities: [3]Ability{AbilityStrength, AbilityDexterity, AbilityConstitution},
		Skills:    [2]Skill{SkillAthletics, SkillIntimidation},
		Feat:      FeatSavageAttacker,
		Tool:      "Gaming Set",
		Equipment: [2]string{"Spear, Shortbow, 20 Arrows, Gaming Set, Healer's Kit, Quiver, Traveler's Clothes, 14 GP", "50 GP"},
		Desc:      "You served in an army, mercenary company, or local militia, learning discipline and martial tactics. You know the brutality of combat, the weight of armor, and the value of standing alongside allies.",
	},
}

func GetBackground(bg Background) BackgroundInfo {
	return backgroundInfo[bg]
}

func GetAllBackgrounds() []BackgroundInfo {
	return backgroundInfo[1:]
}

func GetBackgroundAbilities(bg Background) string {
	bgi := backgroundInfo[bg]
	var abilities []string
	for _, a := range bgi.Abilities {
		ai := GetAbility(a)
		abilities = append(abilities, ai.Name)
	}

	return strings.Join(abilities, ", ")
}

func GetBackgroundEquipmentOptions(bg Background) string {
	bgi := backgroundInfo[bg]
	return "(A) " + bgi.Equipment[0] + ", or (B) " + bgi.Equipment[1]
}

func GetBackgroundSkills(bg Background) string {
	bgi := backgroundInfo[bg]
	var skills []string
	for _, s := range bgi.Skills {
		skills = append(skills, GetSkill(s).Name)
	}

	return strings.Join(skills, ", ")
}
