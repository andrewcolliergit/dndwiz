package character

type OriginFeat int

const (
	FeatNone OriginFeat = iota
	FeatAlert
	FeatMagicInitiate
	FeatSavageAttacker
	FeatSkilled
)

func (f OriginFeat) String() string {
	switch f {
	case FeatAlert:
		return "Alert"
	case FeatMagicInitiate:
		return "Magic Initiate"
	case FeatSavageAttacker:
		return "Savage Attacker"
	case FeatSkilled:
		return "Skilled"
	}
	return "None"
}
