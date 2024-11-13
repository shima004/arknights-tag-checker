package config

type Language string

const (
	LanguageJapanese Language = "ja"
	LanguageEnglish  Language = "en"
)

var DefaultConfig = Config{
	Language: LanguageJapanese,
}

type Config struct {
	Language Language
}

func SetDefaultConfig(config Config) {
	DefaultConfig = config
}
