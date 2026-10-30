// Package class
package class

import (
	"dndwiz/character/ability"
	"dndwiz/character/score"
	"dndwiz/character/skill"
	dice "dndwiz/utils"
)

type Class int

const (
	None Class = iota
	Barbarian
	Bard
	Cleric
	Druid
	Fighter
	Monk
	Paladin
	Ranger
	Rogue
	Sorcerer
	Warlock
	Wizard
)

type Info struct {
	Name                string
	PrimaryAbility      Abilities
	HitPointDie         dice.Die
	SavingThrow         SavingThrow
	SkillProficiency    SkillProficiency
	WeaponProficiencies string
	ArmorTraining       string
	ToolProficiencies   string
	StartingEquipment   Equipment
	Desc                string
	StdArray            score.Ability
}

var classInfo = [...]Info{
	None: {
		Name: "None",
	},
	Barbarian: {
		Name:                "Barbarian",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Strength}},
		HitPointDie:         dice.D12,
		SavingThrow:         SavingThrow{ability.Strength, ability.Constitution},
		SkillProficiency:    SkillProficiency{2, Skills{skill.AnimalHealing, skill.Athletics, skill.Nature, skill.Perception, skill.Survival}},
		WeaponProficiencies: "Simple and Martial weapons",
		ArmorTraining:       "Light and Medium armor and Shields",
		StartingEquipment:   Equipment{"Greataxe, 4 Handaxes, Explorer's Pack, and 15 GP", "75 GP"},
		Desc:                "Driven by a primal fury, these unstoppable warriors transform raw rage into overwhelming martial force.",
		StdArray:            score.Ability{15, 13, 14, 10, 12, 8},
	},
	Bard: {
		Name:                "Bard",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Charisma}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Dexterity, ability.Charisma},
		SkillProficiency:    SkillProficiency{3, skill.All},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light armor",
		StartingEquipment:   Equipment{"Leather Armor, 2 Daggers, Musical Instrument of your choice, Entertainer's Pack, and 19 GP", "90 GP"},
		Desc:                "Master performers who weave magic directly into speech and song to inspire allies, manipulate minds, and alter reality.",
		StdArray:            score.Ability{8, 14, 12, 13, 10, 15},
	},
	Cleric: {
		Name:                "Cleric",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Wisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Wisdom, ability.Charisma},
		SkillProficiency:    SkillProficiency{2, Skills{skill.History, skill.Insight, skill.Medicine, skill.Persuasion, skill.Religion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light and Medium armor and Shields",
		StartingEquipment:   Equipment{"Chain Shirt, Shield, Mace, Holy Symbol, Priest's Pack, and 7 GP", "110 GP"},
		Desc:                "Divine conduits who channel the sacred power of their gods to heal the wounded, protect the innocent, and smite darkness.",
		StdArray:            score.Ability{14, 8, 13, 10, 15, 12},
	},
	Druid: {
		Name:                "Druid",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Wisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Intelligence, ability.Wisdom},
		SkillProficiency:    SkillProficiency{2, Skills{skill.AnimalHealing, skill.Arcana, skill.Insight, skill.Medicine, skill.Nature, skill.Perception, skill.Religion, skill.Survival}},
		WeaponProficiencies: "Simple weapons",
		ToolProficiencies:   "",
		ArmorTraining:       "Light armor and Shields",
		StartingEquipment:   Equipment{"Leather Armor, Shield, Sickle, Duidic Focus (Quarterstaff), Explorer's Pack, Herbalism Kit, and 9 GP", "50 GP"},
		Desc:                "Guardians of the wild who draw magic from nature itself, commanding elemental forces and shifting into fierce beasts.",
		StdArray:            score.Ability{8, 12, 14, 13, 15, 10},
	},
	Fighter: {
		Name:                "Fighter",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Strength, ability.Dexterity}},
		HitPointDie:         dice.D10,
		SavingThrow:         SavingThrow{ability.Strength, ability.Constitution},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Acrobatics, skill.AnimalHealing, skill.Athletics, skill.History, skill.Insight, skill.Intimidation, skill.Persuasion, skill.Perception, skill.Survival}},
		WeaponProficiencies: "Simple and Martial weapons",
		ArmorTraining:       "Light, Medium and Heavy armor and Shields",
		StartingEquipment:   Equipment{"Chain Mail, Greatsword, Flail, 8 Javelins, Dungeoneer's Pack, 4 GP", "Studded Leather Armor, Scimitar, Shortsword, Longbow, 20 Arrows, Quiver, Duneoneer's Pack, and 11 GP", "155 GP"},
		Desc:                "Unmatched masters of martial combat who turn weapon proficiency and tactical discipline into a lethal art form.",
		StdArray:            score.Ability{15, 14, 13, 8, 10, 12},
	},
	Monk: {
		Name:                "Monk",
		PrimaryAbility:      Abilities{2, ability.Abilities{ability.Dexterity, ability.Wisdom}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Strength, ability.Dexterity},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Acrobatics, skill.Athletics, skill.History, skill.Insight, skill.Religion, skill.Stealth}},
		WeaponProficiencies: "Simple weapons and Martial weapons that have the Light property",
		ToolProficiencies:   "Choose one type of Artisan's Tools or Musical Instrument",
		ArmorTraining:       "None",
		StartingEquipment:   Equipment{"Spear, 5 Daggers, Artisan's Tools or Musical Instrument chosen for the tool proficiency, Explorer's Pack, and 11 GP", "50 GP"},
		Desc:                "Agile martial artists who harness inner ki energy to strike with blistering speed and execute extraordinary physical feats.",
		StdArray:            score.Ability{12, 15, 13, 10, 14, 8},
	},
	Paladin: {
		Name:                "Paladin",
		PrimaryAbility:      Abilities{2, ability.Abilities{ability.Strength, ability.Charisma}},
		HitPointDie:         dice.D10,
		SavingThrow:         SavingThrow{ability.Wisdom, ability.Charisma},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Athletics, skill.Insight, skill.Intimidation, skill.Medicine, skill.Persuasion, skill.Religion}},
		WeaponProficiencies: "Simple and Martial Weapons",
		ArmorTraining:       "Light, Medium, and Heavy armor and Sheilds",
		StartingEquipment:   Equipment{"Chain Mail, Shield, Longsword, 6 Javelins, Holy Symbol, Priest's Pack, and 9 GP", "150 GP"},
		Desc:                "Holy warriors bound by sacred oaths, combining martial prowess with divine smites to stand against darkness.",
		StdArray:            score.Ability{15, 10, 13, 8, 12, 14},
	},
	Ranger: {
		Name:                "Ranger",
		PrimaryAbility:      Abilities{2, ability.Abilities{ability.Dexterity, ability.Wisdom}},
		HitPointDie:         dice.D10,
		SavingThrow:         SavingThrow{ability.Strength, ability.Dexterity},
		SkillProficiency:    SkillProficiency{3, Skills{skill.AnimalHealing, skill.Athletics, skill.Insight, skill.Investigation, skill.Nature, skill.Perception, skill.Stealth, skill.Survival}},
		WeaponProficiencies: "Simple and Martial Weapons",
		ArmorTraining:       "Light and Medium armor and Sheilds",
		StartingEquipment:   Equipment{"Studded Leather Armor, Scimitar, Shortsword, Longbow, 20 Arrows, Quiver, Druidic Focus (sprig of mistletoe), Explorer's Pack, and 7 GP", "150 GP"},
		Desc:                "Deadly wilderness scouts and monster hunters who combine precise martial skill with nature-based magic.",
		StdArray:            score.Ability{12, 15, 13, 8, 14, 10},
	},
	Rogue: {
		Name:                "Rogue",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Dexterity}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Dexterity, ability.Intelligence},
		SkillProficiency:    SkillProficiency{4, Skills{skill.Acrobatics, skill.Athletics, skill.Deception, skill.Insight, skill.Intimidation, skill.Investigation, skill.Perception, skill.Persuasion, skill.SleightOfHand, skill.Stealth}},
		WeaponProficiencies: "Simple weapons and Martial weapons that have the Finesse or Light property",
		ToolProficiencies:   "Theives' Tools",
		ArmorTraining:       "Light armor",
		StartingEquipment:   Equipment{"Leather Armor, 2 Daggers, Shortsword, Shortbow, 20 Arrows, Quiver, Theives' Tools, Burglar's Pack, and 8 GP", "100GP"},
		Desc:                "Cunning specialists who strike from the shadows with pinpoint precision, exploiting weaknesses with stealth and agility.",
		StdArray:            score.Ability{12, 15, 13, 14, 10, 8},
	},
	Sorcerer: {
		Name:                "Sorcerer",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Charisma}},
		HitPointDie:         dice.D6,
		SavingThrow:         SavingThrow{ability.Constitution, ability.Charisma},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Arcana, skill.Deception, skill.Insight, skill.Intimidation, skill.Persuasion, skill.Religion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "None",
		StartingEquipment:   Equipment{"Spear, 2 Daggers, Arcane Focus (crystal), Dungeoneer's Pack, and 28 GP", "50 GP"},
		Desc:                "Spellcasters whose magic is an innate, bloodborne force, allowing them to shape and alter the arcane through raw willpower.",
		StdArray:            score.Ability{10, 13, 14, 8, 12, 15},
	},
	Warlock: {
		Name:                "Warlock",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Charisma}},
		HitPointDie:         dice.D8,
		SavingThrow:         SavingThrow{ability.Wisdom, ability.Charisma},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Arcana, skill.Deception, skill.History, skill.Intimidation, skill.Investigation, skill.Nature, skill.Religion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "Light armor",
		StartingEquipment:   Equipment{"Leather Armor, Sickle, 2 Daggers, Arcane Focus (orb), Book (occult lore), Scholar's Pack, and 15 GP", "100 GP"},
		Desc:                "Seekers of forbidden knowledge who bargain with extraplanar entities to wield dark, eldritch powers.",
		StdArray:            score.Ability{8, 14, 13, 12, 10, 15},
	},
	Wizard: {
		Name:                "Wizard",
		PrimaryAbility:      Abilities{1, ability.Abilities{ability.Intelligence}},
		HitPointDie:         dice.D6,
		SavingThrow:         SavingThrow{ability.Intelligence, ability.Wisdom},
		SkillProficiency:    SkillProficiency{2, Skills{skill.Arcana, skill.History, skill.Insight, skill.Investigation, skill.Medicine, skill.Nature, skill.Religion}},
		WeaponProficiencies: "Simple weapons",
		ArmorTraining:       "None",
		StartingEquipment:   Equipment{"2 Daggers, Arcane Focus (Quarterstaff), Robe, Spellbook, Scholar's Pack, and 5 GP", "55 GP"},
		Desc:                "Scholarly masters of the arcane who unravel the laws of reality through intense study and meticulous spellcraft.",
		StdArray:            score.Ability{8, 12, 13, 15, 14, 10},
	},
}

func Get(class Class) Info {
	return classInfo[class]
}

func GetAll() []Info {
	return classInfo[1:]
}

func GetHitPointDieString(class Class) string {
	cls := Get(class)
	return cls.HitPointDie.String() + " per " + cls.Name + " level"
}
