package model

import (
	"arknigths-tag-checker/pkg/config"
	"encoding/json"
	"fmt"
)

type AcquisitionMethod int

const (
	AMStandardHeadHunting AcquisitionMethod = iota + 1
	AMKernelHeadHunting
	AMLimitedHeadHuntingCelebration
	AMLimitedHeadHuntingFestival
	AMLimitedHeadHuntingCarnival
	AMPublicRecruitment
	AMCrossOverHeadHunting
	AMEvent
)

var acquisitionMethodMapJa = map[AcquisitionMethod]string{
	AMStandardHeadHunting:           "スタンダードスカウト",
	AMKernelHeadHunting:             "中堅スカウト",
	AMLimitedHeadHuntingCelebration: "リミテッドスカウト[祭]",
	AMLimitedHeadHuntingFestival:    "リミテッドスカウト[宴]",
	AMLimitedHeadHuntingCarnival:    "リミテッドスカウト[遊]",
	AMPublicRecruitment:             "公開求人",
	AMCrossOverHeadHunting:          "コラボスカウト",
	AMEvent:                         "イベント",
}

var acquisitionMethodMapEn = map[AcquisitionMethod]string{
	AMStandardHeadHunting:           "Standard Head Hunting",
	AMKernelHeadHunting:             "Kernel Head Hunting",
	AMLimitedHeadHuntingCelebration: "Limited Head Hunting[Celebration]",
	AMLimitedHeadHuntingFestival:    "Limited Head Hunting[Festival]",
	AMLimitedHeadHuntingCarnival:    "Limited Head Hunting[Carnival]",
	AMPublicRecruitment:             "Public Recruitment",
	AMCrossOverHeadHunting:          "Cross Over Head Hunting",
	AMEvent:                         "Event",
}

func (am AcquisitionMethod) String() string {
	switch config.DefaultConfig.Language {
	case config.LanguageJapanese:
		return acquisitionMethodMapJa[am]
	default:
		return acquisitionMethodMapEn[am]
	}
}

func (am AcquisitionMethod) MarshalJSON() ([]byte, error) {
	return []byte(`"` + am.String() + `"`), nil
}

func (am *AcquisitionMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch config.DefaultConfig.Language {
	case config.LanguageJapanese:
		for k, v := range acquisitionMethodMapJa {
			if v == s {
				*am = k
				return nil
			}
		}
	default:
		for k, v := range acquisitionMethodMapEn {
			if v == s {
				*am = k
				return nil
			}
		}
	}
	return fmt.Errorf("invalid acquisition method: %s", s)
}
