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

type Info struct{ Name, DamageType, Desc string }

func Get(d Ancestry) Info { return info[d] }

func GetAll() []Info { return info[1:] }

var info = [...]Info{
	Black: {
		Name:       "Black",
		DamageType: "Acid",
		Desc:       "Black dragons haunt foul swamps and ruined places, and their scales are said to carry the sting of corrosive acid. Cruel and spiteful by nature, their descendants spit acid in a sharp line.",
	},
	Blue: {
		Name:       "Blue",
		DamageType: "Lightning",
		Desc:       "Blue dragons rule the desert wastes and coastal dunes, where they crackle with the energy of storms. Their descendants carry a vain, commanding streak and unleash lightning in a line.",
	},
	Brass: {
		Name:       "Brass",
		DamageType: "Fire",
		Desc:       "Brass dragons are talkative creatures of the sun-baked deserts, fond of conversation and riddles. Their descendants are often sociable and quick-witted, with a breath of searing flame.",
	},
	Bronze: {
		Name:       "Bronze",
		DamageType: "Lightning",
		Desc:       "Bronze dragons patrol coastlines and take an active role in protecting the innocent from tyranny. Their descendants are brave and honorable, wielding lightning in a line.",
	},
	Copper: {
		Name:       "Copper",
		DamageType: "Acid",
		Desc:       "Copper dragons are playful tricksters of rocky hills and canyons, known for jokes and clever puzzles. Their descendants share that mischievous humor and breathe a line of acid.",
	},
	Gold: {
		Name:       "Gold",
		DamageType: "Fire",
		Desc:       "Gold dragons are wise, noble guardians who seek to uphold justice and order. Their descendants carry that dignified bearing and exhale a cone of golden fire.",
	},
	Green: {
		Name:       "Green",
		DamageType: "Poison",
		Desc:       "Green dragons lurk in old forests, scheming and manipulating from the shadows of the canopy. Their descendants are sly and calculating, and their breath is a cloud of poison.",
	},
	Red: {
		Name:       "Red",
		DamageType: "Fire",
		Desc:       "Red dragons are the proudest and most fearsome of the chromatic dragons, ruling from volcanic peaks. Their descendants burn with ambition and breathe a cone of fire.",
	},
	Silver: {
		Name:       "Silver",
		DamageType: "Cold",
		Desc:       "Silver dragons are kind and gentle, often taking on friendly mortal forms to walk among other peoples. Their descendants are warm and compassionate, with a breath of freezing cold.",
	},
	White: {
		Name:       "White",
		DamageType: "Cold",
		Desc:       "White dragons are savage, primal hunters of frozen tundra and glaciers. Their descendants are fierce and instinctive, exhaling bitter cold.",
	},
}
