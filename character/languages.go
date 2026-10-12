package character

type Language int

const (
	CommonLang Language = iota
	CommonSignLang
	DraconicLang
	DwarvishLang
	ElvishLang
	GiantLang
	GnomishLang
	GoblinLang
	HalflingLang
	OrcLang
)

// GetLanguage returns a language based on a 1d12
func GetLanguage(roll int) Language {
	switch roll {
	case 1:
		return CommonSignLang
	case 2:
		return DraconicLang
	case 3:
		return DwarvishLang
	case 4:
		return DwarvishLang
	case 5:
		return ElvishLang
	case 6:
		return ElvishLang
	case 7:
		return GiantLang
	case 8:
		return GnomishLang
	case 9:
		return GoblinLang
	case 10:
		return HalflingLang
	case 11:
		return HalflingLang
	case 12:
		return OrcLang
	default:
		return CommonLang
	}
}
