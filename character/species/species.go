// Package species
package species

import (
	"fmt"
	"strings"

	draconic "dndwiz/character/species/draconic_ancestry"
	elven "dndwiz/character/species/elven_lineage"
)

// Size and Speed
type (
	Speed int
	Size  []string
)

func (s Size) RequiresChoice() bool {
	return len(s) > 1
}

func (s Speed) String() string {
	return fmt.Sprintf("%d feet", s)
}

func (s Size) GetAll() string {
	return strings.Join(s, " or ")
}

// Species ...
type Species int

const (
	None Species = iota
	Dragonborn
	Dwarf
	Elf
	Gnome
	Goliath
	Halfling
	Human
	Orc
	Tiefling
)

type Info struct {
	Name             string
	CreatureType     string
	Size             Size
	Speed            Speed
	DraconicAncestry draconic.Ancestry
	ElvenLineage     elven.Lineage
	Desc             string
	Traits           string
}

var info = [...]Info{
	None: {
		Name:         "None",
		CreatureType: "None",
		Size:         Size{},
		Speed:        0,
		Desc:         "",
		Traits:       "",
	},
	Dragonborn: {
		Name:         "Dragonborn",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 5-7 feet tall)"},
		Speed:        30,
		Desc:         "The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons.",
		Traits:       "Dragon Ancestry, Breath Weapon, Damage Resistance, Darkvision, Draconic Flight",
	},
	Dwarf: {
		Name:         "Dwarf",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 4-5 feet tall)"},
		Speed:        30,
		Desc:         "Dwarves were raised from the earth in the elder days by a deity of the forge.",
		Traits:       "Darkvision, Dwarven Resilience, Dwarven Toughness, Stonecunning",
	},
	Elf: {
		Name:         "Elf",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 5-6 feet tall)"},
		Speed:        30,
		Desc:         "The elves' curiosity led many of them to explore other planes of existence.",
		Traits:       "Darkvision, Elven Lineage, Fey Ancestry, Keen Senses, Trance",
	},
	Gnome: {
		Name:         "Gnome",
		CreatureType: "Humanoid",
		Size:         Size{"Small (about 3-4 feet tall)"},
		Speed:        30,
		Desc:         "Gnomes are magical folk created by gods of invention, illusions, and life underground.",
		Traits:       "Darkvision, Gnomish Cunning, Gnomish Lineage",
	},
	Goliath: {
		Name:         "Goliath",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 7-8 feet tall)"},
		Speed:        35,
		Desc:         "Goliaths are distant descendants of giants and seek heights above those reached by their ancestors.",
		Traits:       "Giant Ancestry, Large Form, Powerful Build",
	},
	Halfling: {
		Name:         "Halfling",
		CreatureType: "Humanoid",
		Size:         Size{"Small (about 2-3 feet tall)"},
		Speed:        30,
		Desc:         "Halflings possess a brave and adventurous spirit that leads them on journeys of discovery.",
		Traits:       "Brave, Halfling Nimbleness, Luck, Naturally Stealthy",
	},
	Human: {
		Name:         "Human",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 4-7 feet tall)", "Small (about 2-4 feet tall)"},
		Speed:        30,
		Desc:         "Found throughout the multiverse, humans are as varied as they are numerous.",
		Traits:       "Resourceful, Skillful, Versitile",
	},
	Orc: {
		Name:         "Orc",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 6-7 feet tall)"},
		Speed:        30,
		Desc:         "Orcs are equipped with gifths to help them wander great plains, vast caverns, and churning seas.",
		Traits:       "Adrenaline Rush, Darkvision, Relentless Endurance",
	},
	Tiefling: {
		Name:         "Tiefling",
		CreatureType: "Humanoid",
		Size:         Size{"Medium (about 4-7 feet tall", "Small (about 2-4 feet tall)"},
		Speed:        30,
		Desc:         "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there.",
		Traits:       "Darkvision, fiendish Legacy, Otherworldly Presence",
	},
}

func Get(species Species) Info {
	return info[species]
}

func GetAll() []Info {
	return info[1:]
}
