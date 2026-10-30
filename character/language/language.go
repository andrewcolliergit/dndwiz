// Package language
package language

type Language int

const (
	Common Language = iota
	CommonSign
	Draconic
	Dwarvish
	Elvish
	Giant
	Gnomish
	Goblin
	Halfling
	Orc
)

// GetLanguage returns a language based on a 1d12
func GetLanguage(roll int) Language {
	switch roll {
	case 1:
		return CommonSign
	case 2:
		return Draconic
	case 3:
		return Dwarvish
	case 4:
		return Dwarvish
	case 5:
		return Elvish
	case 6:
		return Elvish
	case 7:
		return Giant
	case 8:
		return Gnomish
	case 9:
		return Goblin
	case 10:
		return Halfling
	case 11:
		return Halfling
	case 12:
		return Orc
	default:
		return Common
	}
}
