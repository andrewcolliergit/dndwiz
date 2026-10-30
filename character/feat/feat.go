// Package feat
package feat

type Feat int

const (
	None Feat = iota
	Alert
	MagicInitiate
	SavageAttacker
	Skilled
)

func (f Feat) String() string {
	switch f {
	case Alert:
		return "Alert"
	case MagicInitiate:
		return "Magic Initiate"
	case SavageAttacker:
		return "Savage Attacker"
	case Skilled:
		return "Skilled"
	}
	return "None"
}
