package config

type Language string

const (
	LanguageJapanese Language = "ja"
	LanguageEnglish  Language = "en"
)

var DefaultConfig = Config{
	DataFileLanguage: LanguageJapanese,
	Language:         LanguageJapanese,
}

type Config struct {
	DataFileLanguage Language
	Language         Language
}

func SetDefaultConfig(config Config) {
	DefaultConfig = config
}
