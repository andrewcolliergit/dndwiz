// Package alignment
package alignment

type Alignment int

const (
	None Alignment = iota
	LawfulGood
	NeutralGood
	ChaoticGood
	LawfulNeutral
	Neutral
	ChaoticNeutral
	LawfulEvil
	NeutralEvil
	ChaoticEvil
)

type Info struct {
	Name string
	Abbr string
	Desc string
}

var info = [...]Info{
	None: {
		Name: "None",
	},
	LawfulGood: {
		Name: "Lawful Good",
		Abbr: "LG",
		Desc: "Lawful Good creatures endeavor to do the right thing as expected by society. Someone who fights injustice and protects the innocent without hesitation is probably Lawful Good.",
	},
	NeutralGood: {
		Name: "Neutral Good",
		Abbr: "NG",
		Desc: "Neutral Good creatures do the best they can, working within rules but not feeling bound by them. A kindly person who helps others according to their needs is probably Neutral Good.",
	},
	ChaoticGood: {
		Name: "Chaotic Good",
		Abbr: "CG",
		Desc: "Chaotic Good creatures act as their conscience directs with little regard for what others expect. A rebel who waylays a cruel baron's tax collectors and uses the stolen money to help the poor is probably Chaotic Good.",
	},
	LawfulNeutral: {
		Name: "Lawful Neutral",
		Abbr: "LN",
		Desc: "Lawful Neutral individuals act in accordance with law, tradition, or personal code. Someone who follows a disciplined rule of life -- and isn't swayed either by the demands of those in need or by the temptations of evil -- is probably Lawful Neutral.",
	},
	Neutral: {
		Name: "Neutral",
		Abbr: "N",
		Desc: "Neutral is the alighment of those who prefer to avoid moral questions and don't take sides, doing what seems best at the time. Someone who's bored by moreal debate is probably Neutral.",
	},
	ChaoticNeutral: {
		Name: "Chaotic Neutral",
		Abbr: "CN",
		Desc: "Chaotic Neurtral creatures follow their whims, valuing their personal freedom above all else. A scoundrel who wanders the land living by their wits is probably Chaotic Neutral.",
	},
	LawfulEvil: {
		Name: "Lawful Evil",
		Abbr: "LE",
		Desc: "Lawful Evil creatures methodically take what they want within the limits of a code of tradition, loyalty, or order. An aristocrat exploiting citizens while scheming for power is probably Lawful Evil.",
	},
	NeutralEvil: {
		Name: "Neutral Evil",
		Abbr: "NE",
		Desc: "Neutral Evil is the alignment of those who are untroubled by the harm they cause as they pursue their desires. A criminal who robs and murders as they please is probably Neutral Evil.",
	},
	ChaoticEvil: {
		Name: "Chaotic Evil",
		Abbr: "CE",
		Desc: "Chaotic Evil creatures act with arbitrary violence, spurred by their hatred or bloodlust. A villain pursuing schemes of vengeance and havoc is probably Chaotic Evil.",
	},
}

func Get(a Alignment) Info {
	return info[a]
}

func GetAll() []Info {
	return info[1:]
}
