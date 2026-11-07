// Package elvenlineage
package elvenlineage

type Lineage int

const (
	None = iota
	Drow
	HighElf
	WoodElf
)

type Info struct{ Name, Desc, Lvl1, Lvl3, Lvl5 string }

func Get(e Lineage) Info { return info[e] }

func GetAll() []Info { return info[1:] }

var info = [...]Info{
	Drow: {
		Name: "Drow Elf",
		Desc: "The drow trace their heritage to the lightless caverns of the Underdark, where their ancestors adapted to a world without sun. Their descendants are graceful and secretive, at home in darkness and shadow.",
		Lvl1: "The range of your Darkvision increases to 120 feet. You also know the Dancing Lights cantrip.",
		Lvl3: "Faerie Fire",
		Lvl5: "Darkness",
	},
	HighElf: {
		Name: "High Elf",
		Desc: "High elves are scholars and aesthetes, drawn to the study of magic and the refined arts of their ancient courts. Their descendants carry a sharp intellect and an easy familiarity with the arcane.",
		Lvl1: "You know the Prestidigitation cantrip. Whenever you finish a Long Rest, you can replace that cantrip with a different cantrip from the Wizard spell list.",
		Lvl3: "Detect Magic",
		Lvl5: "Misty Step",
	},
	WoodElf: {
		Name: "Wood Elf",
		Desc: "Wood elves make their homes deep in ancient forests, living in step with the rhythms of the wild. Their descendants are swift and watchful, and move through the wilderness as though they were part of it.",
		Lvl1: "Your Speed increases to 35 feet. You also Know the Druidcraft cantrip.",
		Lvl3: "Longstrider",
		Lvl5: "Pass without Trace",
	},
}
