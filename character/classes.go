package character

import (
	"fmt"
	"strings"

	dice "dndwiz/utils"
)

type Class int

const (
	ClassNone Class = iota
	ClassBarbarian
	ClassBard
	ClassCleric
	ClassDruid
	ClassFighter
	ClassMonk
	ClassPaladin
	ClassRanger
	ClassRogue
	ClassSorcerer
	ClassWarlock
	ClassWizard
)

// Class Skill Proficiencies
type (
	ClassSkills           []Skill
	ClassSkillProficiency struct {
		Has  int
		From ClassSkills
	}
)

func (csp ClassSkillProficiency) String() string {
	prefix := fmt.Sprintf("Choose %d: ", csp.Has)
	var str []string
	for _, skill := range csp.From {
		str = append(str, GetSkill(skill).Name)
	}

	return prefix + strings.Join(str, ", ")
}

func (csp ClassSkillProficiency) RequiresChoice() bool {
	return csp.Has < len(csp.From)
}

type ClassSavingThrow []Ability

func (st ClassSavingThrow) String() string {
	var str []string
	for _, a := range st {
		str = append(str, GetAbility(a).Name)
	}
	return strings.Join(str, ", ")
}

type ClassStartingEquipment []string

func (e ClassStartingEquipment) String() string {
	letter := []string{"A", "B", "C", "D"}
	var sb strings.Builder

	for i, item := range e {
		if i > 0 {
			sb.WriteString(", ")
		}
		if i < len(letter) {
			sb.WriteString("(")
			sb.WriteString(letter[i])
			sb.WriteString(") ")
		}
		sb.WriteString(item)
	}

	return sb.String()
}

type ClassInfo struct {
	Name                string
	PrimaryAbility      ClassAbilities
	HitPointDie         dice.Die
	SavingThrow         ClassSavingThrow
	SkillProficiency    ClassSkillProficiency
	WeaponProficiencies string
	ArmorTraining       string
	ToolProficiencies   string
	StartingEquipment   ClassStartingEquipment
	Desc                string
	StdArray            AbilityScores
}

var classInfo = [...]ClassInfo{
	ClassNone: {
		Name: "None",
	},
	ClassBarbarian: {
		Name:                "Barbarian",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityStrength}},
		HitPointDie:         dice.D12,
		SavingThrow:         ClassSavingThrow{AbilityStrength, AbilityConstitution},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillAnimalHealing, SkillAthletics, SkillNature, SkillPerception, SkillSurvival}},
		WeaponProficiencies: "Simple and Martial weapons",
		ArmorTraining:       "Light and Medium armor and Shields",
		StartingEquipment:   ClassStartingEquipment{"Greataxe, 4 Handaxes, Explorer's Pack, and 15 GP", "75 GP"},
		Desc:                "",
		StdArray:            AbilityScores{15, 13, 14, 10, 12, 8},
	},
	ClassBard: {
		Name:                "Bard",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityCharisma}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityDexterity, AbilityCharisma},
		SkillProficiency:    ClassSkillProficiency{3, AllSkills},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light armor",
		StartingEquipment:   ClassStartingEquipment{"Leather Armor, 2 Daggers, Musical Instrument of your choice, Entertainer's Pack, and 19 GP", "90 GP"},
		Desc:                "",
		StdArray:            AbilityScores{8, 14, 12, 13, 10, 15},
	},
	ClassCleric: {
		Name:                "Cleric",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityWisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityWisdom, AbilityCharisma},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillHistory, SkillInsight, SkillMedicine, SkillPersuasion, SkillReligion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light and Medium armor and Shields",
		StartingEquipment:   ClassStartingEquipment{"Chain Shirt, Shield, Mace, Holy Symbol, Priest's Pack, and 7 GP", "110 GP"},
		Desc:                "",
		StdArray:            AbilityScores{14, 8, 13, 10, 15, 12},
	},
	ClassDruid: {
		Name:                "Druid",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityWisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityIntelligence, AbilityWisdom},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillAnimalHealing, SkillArcana, SkillInsight, SkillMedicine, SkillNature, SkillPerception, SkillReligion, SkillSurvival}},
		WeaponProficiencies: "Simple weapons",
		ToolProficiencies:   "",
		ArmorTraining:       "Light armor and Shields",
		StartingEquipment:   ClassStartingEquipment{"Leather Armor, Shield, Sickle, Duidic Focus (Quarterstaff), Explorer's Pack, Herbalism Kit, and 9 GP", "50 GP"},
		Desc:                "",
		StdArray:            AbilityScores{8, 12, 14, 13, 15, 10},
	},
	ClassFighter: {
		Name:                "Fighter",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityStrength, AbilityDexterity}},
		HitPointDie:         dice.D10,
		SavingThrow:         ClassSavingThrow{AbilityStrength, AbilityConstitution},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillAcrobatics, SkillAnimalHealing, SkillAthletics, SkillHistory, SkillInsight, SkillIntimidation, SkillPersuasion, SkillPerception, SkillSurvival}},
		WeaponProficiencies: "Simple and Martial weapons",
		ArmorTraining:       "Light, Medium and Heavy armor and Shields",
		StartingEquipment:   ClassStartingEquipment{"Chain Mail, Greatsword, Flail, 8 Javelins, Dungeoneer's Pack, 4 GP", "Studded Leather Armor, Scimitar, Shortsword, Longbow, 20 Arrows, Quiver, Duneoneer's Pack, and 11 GP", "155 GP"},
		Desc:                "",
		StdArray:            AbilityScores{15, 14, 13, 8, 10, 12},
	},
	ClassMonk: {
		Name:                "Monk",
		PrimaryAbility:      ClassAbilities{2, Abilities{AbilityDexterity, AbilityWisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityStrength, AbilityDexterity},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillAcrobatics, SkillAthletics, SkillHistory, SkillInsight, SkillReligion, SkillStealth}},
		WeaponProficiencies: "Simple weapons and Martial weapons that have the Light property",
		ToolProficiencies:   "Choose one type of Artisan's Tools or Musical Instrument",
		ArmorTraining:       "None",
		StartingEquipment:   ClassStartingEquipment{"Spear, 5 Daggers, Artisan's Tools or Musical Instrument chosen for the tool proficiency, Explorer's Pack, and 11 GP", "50 GP"},
		Desc:                "",
		StdArray:            AbilityScores{12, 15, 13, 10, 14, 8},
	},
	ClassPaladin: {
		Name:                "Paladin",
		PrimaryAbility:      ClassAbilities{2, Abilities{AbilityStrength, AbilityCharisma}},
		HitPointDie:         dice.D10,
		SavingThrow:         ClassSavingThrow{AbilityWisdom, AbilityCharisma},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillAthletics, SkillInsight, SkillIntimidation, SkillMedicine, SkillPersuasion, SkillReligion}},
		WeaponProficiencies: "Simple and Martial Weapons",
		ArmorTraining:       "Light, Medium, and Heavy armor and Sheilds",
		StartingEquipment:   ClassStartingEquipment{"Chain Mail, Shield, Longsword, 6 Javelins, Holy Symbol, Priest's Pack, and 9 GP", "150 GP"},
		Desc:                "",
		StdArray:            AbilityScores{15, 10, 13, 8, 12, 14},
	},
	ClassRanger: {
		Name:                "Ranger",
		PrimaryAbility:      ClassAbilities{2, Abilities{AbilityDexterity, AbilityWisdom}},
		HitPointDie:         dice.D10,
		SavingThrow:         ClassSavingThrow{AbilityStrength, AbilityDexterity},
		SkillProficiency:    ClassSkillProficiency{3, ClassSkills{SkillAnimalHealing, SkillAthletics, SkillInsight, SkillInvestigation, SkillNature, SkillPerception, SkillStealth, SkillSurvival}},
		WeaponProficiencies: "Simple and Martial Weapons",
		ArmorTraining:       "Light and Medium armor and Sheilds",
		StartingEquipment:   ClassStartingEquipment{"Studded Leather Armor, Scimitar, Shortsword, Longbow, 20 Arrows, Quiver, Druidic Focus (sprig of mistletoe), Explorer's Pack, and 7 GP", "150 GP"},
		Desc:                "",
		StdArray:            AbilityScores{12, 15, 13, 8, 14, 10},
	},
	ClassRogue: {
		Name:                "Rogue",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityDexterity}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityDexterity, AbilityIntelligence},
		SkillProficiency:    ClassSkillProficiency{4, ClassSkills{SkillAcrobatics, SkillAthletics, SkillDeception, SkillInsight, SkillIntimidation, SkillInvestigation, SkillPerception, SkillPersuasion, SkillSleightOfHand, SkillStealth}},
		WeaponProficiencies: "Simple weapons and Martial weapons that have the Finesse or Light property",
		ToolProficiencies:   "Theives' Tools",
		ArmorTraining:       "Light armor",
		StartingEquipment:   ClassStartingEquipment{"Leather Armor, 2 Daggers, Shortsword, Shortbow, 20 Arrows, Quiver, Theives' Tools, Burglar's Pack, and 8 GP", "100GP"},
		Desc:                "",
		StdArray:            AbilityScores{12, 15, 13, 14, 10, 8},
	},
	ClassSorcerer: {
		Name:                "Sorcerer",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityCharisma}},
		HitPointDie:         dice.D6,
		SavingThrow:         ClassSavingThrow{AbilityConstitution, AbilityCharisma},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillArcana, SkillDeception, SkillInsight, SkillIntimidation, SkillPersuasion, SkillReligion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "None",
		StartingEquipment:   ClassStartingEquipment{"Spear, 2 Daggers, Arcane Focus (crystal), Dungeoneer's Pack, and 28 GP", "50 GP"},
		Desc:                "",
		StdArray:            AbilityScores{10, 13, 14, 8, 12, 15},
	},
	ClassWarlock: {
		Name:                "Warlock",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityCharisma}},
		HitPointDie:         dice.D8,
		SavingThrow:         ClassSavingThrow{AbilityWisdom, AbilityCharisma},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillArcana, SkillDeception, SkillHistory, SkillIntimidation, SkillInvestigation, SkillNature, SkillReligion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light armor",
		StartingEquipment:   ClassStartingEquipment{"Leather Armor, Sickle, 2 Daggers, Arcane Focus (orb), Book (occult lore), Scholar's Pack, and 15 GP", "100 GP"},
		Desc:                "",
		StdArray:            AbilityScores{8, 14, 13, 12, 10, 15},
	},
	ClassWizard: {
		Name:                "Wizard",
		PrimaryAbility:      ClassAbilities{1, Abilities{AbilityIntelligence}},
		HitPointDie:         dice.D6,
		SavingThrow:         ClassSavingThrow{AbilityIntelligence, AbilityWisdom},
		SkillProficiency:    ClassSkillProficiency{2, ClassSkills{SkillArcana, SkillHistory, SkillInsight, SkillInvestigation, SkillMedicine, SkillNature, SkillReligion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "None",
		StartingEquipment:   ClassStartingEquipment{"2 Daggers, Arcane Focus (Quarterstaff), Robe, Spellbook, Scholar's Pack, and 5 GP", "55 GP"},
		Desc:                "",
		StdArray:            AbilityScores{8, 12, 13, 15, 14, 10},
	},
}

func GetClass(class Class) ClassInfo {
	return classInfo[class]
}

func GetAllClasses() []ClassInfo {
	return classInfo[1:]
}

func GetHitPointDieString(class Class) string {
	cls := GetClass(class)
	return cls.HitPointDie.String() + " per " + cls.Name + " level"
}
