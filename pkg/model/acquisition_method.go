package model

import (
	"encoding/json"
	"fmt"

	"github.com/shima004/arknights-tag-checker/pkg/config"
)

type AcquisitionMethod int
type AcquisitionMethods []AcquisitionMethod

const (
	AMStandardHeadHunting AcquisitionMethod = iota + 1
	AMKernelHeadHunting
	AMLimitedHeadHuntingCelebration
	AMLimitedHeadHuntingFestival
	AMLimitedHeadHuntingCarnival
	AMPublicRecruitment
	AMCrossOverHeadHunting
	AMCrossOverEvent
	AMEvent
	AMMainStory
	AMIntegratedStrategies
	AMPurchaseCertificateStore
	AMAnniversary
	AMReclamationAlgorithm
	AMContingencyContract
)

var acquisitionMethodMapJa = map[AcquisitionMethod]string{
	AMStandardHeadHunting:           "スタンダードスカウト",
	AMKernelHeadHunting:             "中堅スカウト",
	AMLimitedHeadHuntingCelebration: "リミテッドスカウト[祭]",
	AMLimitedHeadHuntingFestival:    "リミテッドスカウト[宴]",
	AMLimitedHeadHuntingCarnival:    "リミテッドスカウト[遊]",
	AMPublicRecruitment:             "公開求人",
	AMCrossOverHeadHunting:          "コラボスカウト",
	AMCrossOverEvent:                "コラボイベント報酬",
	AMEvent:                         "イベント報酬",
	AMMainStory:                     "メインストーリー",
	AMIntegratedStrategies:          "統合戦略",
	AMPurchaseCertificateStore:      "購買資格証",
	AMAnniversary:                   "周年記念",
	AMReclamationAlgorithm:          "生息演算",
	AMContingencyContract:           "危機契約",
}

var acquisitionMethodMapEn = map[AcquisitionMethod]string{
	AMStandardHeadHunting:           "Standard Head Hunting",
	AMKernelHeadHunting:             "Kernel Head Hunting",
	AMLimitedHeadHuntingCelebration: "Limited Head Hunting[Celebration]",
	AMLimitedHeadHuntingFestival:    "Limited Head Hunting[Festival]",
	AMLimitedHeadHuntingCarnival:    "Limited Head Hunting[Carnival]",
	AMPublicRecruitment:             "Public Recruitment",
	AMCrossOverHeadHunting:          "Cross Over Head Hunting",
	AMCrossOverEvent:                "Cross Over Event Reward",
	AMEvent:                         "Event Reward",
	AMMainStory:                     "Main Story",
	AMIntegratedStrategies:          "Integrated Strategies",
	AMPurchaseCertificateStore:      "Purchase Certificate Store",
	AMAnniversary:                   "Anniversary",
	AMReclamationAlgorithm:          "Reclamation Algorithm",
	AMContingencyContract:           "Contingency Contract",
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
	switch config.DefaultConfig.DataFileLanguage {
	case config.LanguageJapanese:
		return json.Marshal(acquisitionMethodMapJa[am])
	default:
		return json.Marshal(acquisitionMethodMapEn[am])
	}
}

func (am *AcquisitionMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch config.DefaultConfig.DataFileLanguage {
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
