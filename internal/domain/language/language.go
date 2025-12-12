package language

type LanguageCode string

var (
	En LanguageCode = "en"
	Ja LanguageCode = "ja"
)

var supportedLanguages = []LanguageCode{En, Ja}

type Language string

func (l Language) String() string {
	return string(l)
}

func NewLanguage(s string) (Language, error) {
	for _, lang := range supportedLanguages {
		if string(lang) == s {
			return Language(s), nil
		}
	}
	return Language(""), ErrUnsupportedLanguage
}
