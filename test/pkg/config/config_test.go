package config_test

import (
	"testing"

	"github.com/shima004/arknights-tag-checker/pkg/config"
)

func TestSetDefaultConfig(t *testing.T) {
	if config.DefaultConfig.Language != config.LanguageJapanese {
		t.Errorf("DefaultConfig.Language = %v, want %v", config.DefaultConfig.Language, config.LanguageJapanese)
	}

	c := config.Config{
		Language: config.LanguageEnglish,
	}
	config.SetDefaultConfig(c)
	if config.DefaultConfig.Language != config.LanguageEnglish {
		t.Errorf("DefaultConfig.Language = %v, want %v", config.DefaultConfig.Language, config.LanguageEnglish)
	}
}
