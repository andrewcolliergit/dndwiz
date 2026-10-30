// Package elvenlineage
package elvenlineage

type Lineage int

const (
	None = iota
	Drow
	HighElf
	WoodElf
)

type Info struct{ Name, Lvl1, Lvl3, Lvl5 string }

func Get(e Lineage) Info { return info[e] }

func GetAll() []Info { return info[1:] }

var info = [...]Info{
	Drow: {
		Name: "Drow Elf",
		Lvl1: "The range of your Darkvision increases to 120 feet. You also know the Dancing Lights cantrip.",
		Lvl3: "Faerie Fire",
		Lvl5: "Darkness",
	},
	HighElf: {
		Name: "High Elf",
		Lvl1: "You know the Prestidigitation cantrip. Whenever you finish a Long Rest, you can replace that cantrip with a different cantrip from the Wizard spell list.",
		Lvl3: "Detect Magic",
		Lvl5: "Misty Step",
	},
	WoodElf: {
		Name: "Wood Elf",
		Lvl1: "Your Speed increases to 35 feet. You also Know the Druidcraft cantrip.",
		Lvl3: "Longstrider",
		Lvl5: "Pass without Trace",
	},
}
