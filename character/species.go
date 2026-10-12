package character

import (
	"fmt"
	"strings"
)

// Size and Speed
type (
	SpeciesSpeed int
	SpeciesSize  []string
)

func (s SpeciesSize) RequiresChoice() bool {
	return len(s) > 1
}

func (s SpeciesSpeed) String() string {
	return fmt.Sprintf("%d feet", s)
}

func (s SpeciesSize) GetAll() string {
	return strings.Join(s, " or ")
}

// Draconic Ancestry

type DraconicAncestry int

const (
	DraconicAncestryNone DraconicAncestry = iota
	DraconicAncestryBlack
	DraconicAncestryBlue
	DraconicAncestryBrass
	DraconicAncestryBronze
	DraconicAncestryCopper
	DraconicAncestryGold
	DraconicAncestryGreen
	DraconicAncestryRed
	DraconicAncestrySilver
	DraconicAncestryWhite
)

type DraconicAncestryInfo struct{ Name, DamageType string }

func GetDraconicAncestry(d DraconicAncestry) DraconicAncestryInfo { return draconicAncestoryInfo[d] }

var draconicAncestoryInfo = [...]DraconicAncestryInfo{
	DraconicAncestryBlack: {
		Name:       "Black",
		DamageType: "Acid",
	},
	DraconicAncestryBlue: {
		Name:       "Blue",
		DamageType: "Lightning",
	},
	DraconicAncestryBrass: {
		Name:       "Brass",
		DamageType: "Fire",
	},
	DraconicAncestryBronze: {
		Name:       "Bronze",
		DamageType: "Lightning",
	},
	DraconicAncestryCopper: {
		Name:       "Copper",
		DamageType: "Acid",
	},
	DraconicAncestryGold: {
		Name:       "Gold",
		DamageType: "Fire",
	},
	DraconicAncestryGreen: {
		Name:       "Green",
		DamageType: "Poison",
	},
	DraconicAncestryRed: {
		Name:       "Red",
		DamageType: "Fire",
	},
	DraconicAncestrySilver: {
		Name:       "Silver",
		DamageType: "Cold",
	},
	DraconicAncestryWhite: {
		Name:       "White",
		DamageType: "Cold",
	},
}

func GetAllDraconicAncestries() []DraconicAncestryInfo {
	return draconicAncestoryInfo[1:]
}

type ElvenLineage int

const (
	ElvenLineageNone = iota
	ElvenLineageDrow
	ElvenLineageHighElf
	ElvenLineageWoodElf
)

type ElvenLineageInfo struct{ Name, Lvl1, Lvl3, Lvl5 string }

func GetElvenLineage(e ElvenLineage) ElvenLineageInfo {
	return elvenLineageInfo[e]
}

var elvenLineageInfo = [...]ElvenLineageInfo{
	ElvenLineageDrow: {
		Name: "Drow Elf",
		Lvl1: "The range of your Darkvision increases to 120 feet. You also know the Dancing Lights cantrip.",
		Lvl3: "Faerie Fire",
		Lvl5: "Darkness",
	},
	ElvenLineageHighElf: {
		Name: "High Elf",
		Lvl1: "You know the Prestidigitation cantrip. Whenever you finish a Long Rest, you can replace that cantrip with a different cantrip from the Wizard spell list.",
		Lvl3: "Detect Magic",
		Lvl5: "Misty Step",
	},
	ElvenLineageWoodElf: {
		Name: "Wood Elf",
		Lvl1: "Your Speed increases to 35 feet. You also Know the Druidcraft cantrip.",
		Lvl3: "Longstrider",
		Lvl5: "Pass without Trace",
	},
}

func GetAllElvenLineages() []ElvenLineageInfo {
	return elvenLineageInfo[1:]
}

// Species ...
type Species int

const (
	SpeciesNone Species = iota
	SpeciesDragonborn
	SpeciesDwarf
	SpeciesElf
	SpeciesGnome
	SpeciesGoliath
	SpeciesHalfling
	SpeciesHuman
	SpeciesOrc
	SpeciesTiefling
)

type SpeciesInfo struct {
	Name             string
	CreatureType     string
	Size             SpeciesSize
	Speed            SpeciesSpeed
	DraconicAncestry DraconicAncestry
	ElvenLineage     ElvenLineage
	Desc             string
	Traits           string
}

var speciesInfo = [...]SpeciesInfo{
	SpeciesNone: {
		Name:             "None",
		CreatureType:     "None",
		Size:             SpeciesSize{},
		Speed:            0,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "",
		Traits:           "",
	},
	SpeciesDragonborn: {
		Name:             "Dragonborn",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 5-7 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons.",
		Traits:           "Dragon Ancestry, Breath Weapon, Damage Resistance, Darkvision, Draconic Flight",
	},
	SpeciesDwarf: {
		Name:             "Dwarf",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 4-5 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Dwarves were raised from the earth in the elder days by a deity of the forge.",
		Traits:           "Darkvision, Dwarven Resilience, Dwarven Toughness, Stonecunning",
	},
	SpeciesElf: {
		Name:             "Elf",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 5-6 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "The elves' curiosity led many of them to explore other planes of existence.",
		Traits:           "Darkvision, Elven Lineage, Fey Ancestry, Keen Senses, Trance",
	},
	SpeciesGnome: {
		Name:             "Gnome",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Small (about 3-4 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Gnomes are magical folk created by gods of invention, illusions, and life underground.",
		Traits:           "Darkvision, Gnomish Cunning, Gnomish Lineage",
	},
	SpeciesGoliath: {
		Name:             "Goliath",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 7-8 feet tall)"},
		Speed:            35,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Goliaths are distant descendants of giants and seek heights above those reached by their ancestors.",
		Traits:           "Giant Ancestry, Large Form, Powerful Build",
	},
	SpeciesHalfling: {
		Name:             "Halfling",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Small (about 2-3 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Halflings possess a brave and adventurous spirit that leads them on journeys of discovery.",
		Traits:           "Brave, Halfling Nimbleness, Luck, Naturally Stealthy",
	},
	SpeciesHuman: {
		Name:             "Human",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 4-7 feet tall)", "Small (about 2-4 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Found throughout the multiverse, humans are as varied as they are numerous.",
		Traits:           "Resourceful, Skillful, Versitile",
	},
	SpeciesOrc: {
		Name:             "Orc",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 6-7 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Orcs are equipped with gifths to help them wander great plains, vast caverns, and churning seas.",
		Traits:           "Adrenaline Rush, Darkvision, Relentless Endurance",
	},
	SpeciesTiefling: {
		Name:             "Tiefling",
		CreatureType:     "Humanoid",
		Size:             SpeciesSize{"Medium (about 4-7 feet tall", "Small (about 2-4 feet tall)"},
		Speed:            30,
		DraconicAncestry: DraconicAncestryNone,
		ElvenLineage:     ElvenLineageNone,
		Desc:             "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there.",
		Traits:           "Darkvision, fiendish Legacy, Otherworldly Presence",
	},
}

func GetSpecies(species Species) SpeciesInfo {
	return speciesInfo[species]
}

func GetAllSpecies() []SpeciesInfo {
	return speciesInfo[1:]
}
