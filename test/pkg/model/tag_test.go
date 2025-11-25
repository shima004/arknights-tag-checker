package model_test

import (
	"encoding/json"
	"testing"

	"github.com/shima004/arknights-tag-checker/pkg/config"
	"github.com/shima004/arknights-tag-checker/pkg/model"
)

func TestTagString(t *testing.T) {
	tests := []struct {
		name     string
		tag      model.Tag
		language config.Language
		want     string
	}{
		{
			name:     "Vanguard in Japanese",
			tag:      model.TagVanguard,
			language: config.LanguageJapanese,
			want:     "先鋒タイプ",
		},
		{
			name:     "Vanguard in English",
			tag:      model.TagVanguard,
			language: config.LanguageEnglish,
			want:     "Vanguard",
		},
		{
			name:     "TopOperator in Japanese",
			tag:      model.TagTopOperator,
			language: config.LanguageJapanese,
			want:     "上級エリート",
		},
		{
			name:     "TopOperator in English",
			tag:      model.TagTopOperator,
			language: config.LanguageEnglish,
			want:     "Top Operator",
		},
		{
			name:     "Melee in Japanese",
			tag:      model.TagMelee,
			language: config.LanguageJapanese,
			want:     "近距離",
		},
		{
			name:     "Melee in English",
			tag:      model.TagMelee,
			language: config.LanguageEnglish,
			want:     "Melee",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 元の設定を保存
			originalLang := config.DefaultConfig.Language
			defer func() {
				config.DefaultConfig.Language = originalLang
			}()

			// テスト用の言語を設定
			config.DefaultConfig.Language = tt.language

			got := tt.tag.String()
			if got != tt.want {
				t.Errorf("Tag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		json         string
		dataFileLang config.Language
		want         model.Tag
		wantErr      bool
	}{
		{
			name:         "Unmarshal Vanguard in Japanese",
			json:         `"先鋒タイプ"`,
			dataFileLang: config.LanguageJapanese,
			want:         model.TagVanguard,
			wantErr:      false,
		},
		{
			name:         "Unmarshal Vanguard in English",
			json:         `"Vanguard"`,
			dataFileLang: config.LanguageEnglish,
			want:         model.TagVanguard,
			wantErr:      false,
		},
		{
			name:         "Unmarshal TopOperator in Japanese",
			json:         `"上級エリート"`,
			dataFileLang: config.LanguageJapanese,
			want:         model.TagTopOperator,
			wantErr:      false,
		},
		{
			name:         "Unmarshal TopOperator in English",
			json:         `"Top Operator"`,
			dataFileLang: config.LanguageEnglish,
			want:         model.TagTopOperator,
			wantErr:      false,
		},
		{
			name:         "Unmarshal unknown tag",
			json:         `"UnknownTag"`,
			dataFileLang: config.LanguageEnglish,
			want:         0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 元の設定を保存
			originalDataFileLang := config.DefaultConfig.DataFileLanguage
			defer func() {
				config.DefaultConfig.DataFileLanguage = originalDataFileLang
			}()

			// テスト用のデータファイル言語を設定
			config.DefaultConfig.DataFileLanguage = tt.dataFileLang

			var got model.Tag
			err := json.Unmarshal([]byte(tt.json), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tag.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Tag.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagMarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		tag          model.Tag
		dataFileLang config.Language
		want         string
	}{
		{
			name:         "Marshal Vanguard in Japanese",
			tag:          model.TagVanguard,
			dataFileLang: config.LanguageJapanese,
			want:         `"先鋒タイプ"`,
		},
		{
			name:         "Marshal Vanguard in English",
			tag:          model.TagVanguard,
			dataFileLang: config.LanguageEnglish,
			want:         `"Vanguard"`,
		},
		{
			name:         "Marshal TopOperator in Japanese",
			tag:          model.TagTopOperator,
			dataFileLang: config.LanguageJapanese,
			want:         `"上級エリート"`,
		},
		{
			name:         "Marshal TopOperator in English",
			tag:          model.TagTopOperator,
			dataFileLang: config.LanguageEnglish,
			want:         `"Top Operator"`,
		},
		{
			name:         "Marshal Soar in Japanese",
			tag:          model.TagSoar,
			dataFileLang: config.LanguageJapanese,
			want:         `"高空"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 元の設定を保存
			originalDataFileLang := config.DefaultConfig.DataFileLanguage
			defer func() {
				config.DefaultConfig.DataFileLanguage = originalDataFileLang
			}()

			// テスト用のデータファイル言語を設定
			config.DefaultConfig.DataFileLanguage = tt.dataFileLang

			got, err := json.Marshal(tt.tag)
			if err != nil {
				t.Errorf("Tag.MarshalJSON() error = %v", err)
				return
			}
			if string(got) != tt.want {
				t.Errorf("Tag.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestTagsCombinationsTags(t *testing.T) {
	tests := []struct {
		name string
		tags model.Tags
		want int // 組み合わせの数
	}{
		{
			name: "Empty tags",
			tags: model.Tags{},
			want: 0,
		},
		{
			name: "Single tag",
			tags: model.Tags{model.TagVanguard},
			want: 1, // [Vanguard]
		},
		{
			name: "Two tags",
			tags: model.Tags{model.TagVanguard, model.TagMelee},
			want: 3, // [Vanguard], [Melee], [Vanguard, Melee]
		},
		{
			name: "Three tags",
			tags: model.Tags{model.TagVanguard, model.TagMelee, model.TagDPRecovery},
			want: 7, // 2^3 - 1 = 7 (空集合を除く)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tags.CombinationsTags()
			if len(got) != tt.want {
				t.Errorf("Tags.CombinationsTags() returned %d combinations, want %d", len(got), tt.want)
			}
		})
	}
}

func TestTagsCombinationsTagsContent(t *testing.T) {
	tags := model.Tags{model.TagVanguard, model.TagMelee}
	result := tags.CombinationsTags()

	// 期待される組み合わせ
	expected := []model.Tags{
		{model.TagVanguard},
		{model.TagVanguard, model.TagMelee},
		{model.TagMelee},
	}

	if len(result) != len(expected) {
		t.Fatalf("CombinationsTags() returned %d combinations, want %d", len(result), len(expected))
	}

	// 各組み合わせが期待される結果に含まれているか確認
	for _, exp := range expected {
		found := false
		for _, res := range result {
			if tagsEqual(res, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected combination %v not found in result", exp)
		}
	}
}

func tagsEqual(a, b model.Tags) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
