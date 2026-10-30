// Package draconicancestry
package draconicancestry

type Ancestry int

const (
	None Ancestry = iota
	Black
	Blue
	Brass
	Bronze
	Copper
	Gold
	Green
	Red
	Silver
	White
)

type Info struct{ Name, DamageType string }

func Get(d Ancestry) Info { return info[d] }

func GetAll() []Info { return info[1:] }

var info = [...]Info{
	Black: {
		Name:       "Black",
		DamageType: "Acid",
	},
	Blue: {
		Name:       "Blue",
		DamageType: "Lightning",
	},
	Brass: {
		Name:       "Brass",
		DamageType: "Fire",
	},
	Bronze: {
		Name:       "Bronze",
		DamageType: "Lightning",
	},
	Copper: {
		Name:       "Copper",
		DamageType: "Acid",
	},
	Gold: {
		Name:       "Gold",
		DamageType: "Fire",
	},
	Green: {
		Name:       "Green",
		DamageType: "Poison",
	},
	Red: {
		Name:       "Red",
		DamageType: "Fire",
	},
	Silver: {
		Name:       "Silver",
		DamageType: "Cold",
	},
	White: {
		Name:       "White",
		DamageType: "Cold",
	},
}
