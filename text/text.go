package text

/**
TODO: 폰트
*/

const (
	English = iota
	Korean
)

var language = English

func SetLanguageToEnglish() {
	language = English
}

func SetLanguageToKorean() {
	language = Korean
}

type Text struct {
	English string
	Korean  string
}

func NewText(english string, korean string) *Text {
	return &Text{
		English: english,
		Korean:  korean,
	}
}

func (t Text) Get() string {
	switch language {
	case English:
		return t.English
	case Korean:
		return t.Korean
	default:
		return t.English
	}
}
