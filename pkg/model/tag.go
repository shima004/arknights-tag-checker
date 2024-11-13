package model

import (
	"arknigths-tag-checker/pkg/config"
	"encoding/json"
	"fmt"
)

type Tag int

const (
	TagVanguard Tag = iota
	TagGuard
	TagDefender
	TagSniper
	TagCaster
	TagMedic
	TagSupporter
	TagSpecialist
	TagMelee
	TagRanged
	TagDPS
	TagDefense
	TagSurvival
	TagHealing
	TagSlow
	TagInitial
	TagCrowdControl
	TagSupport
	TagDebuff
	TagSummon
	TagElemental
	TagNuker
	TagRobot
	TagShift
	TagAoE
	TagDPRecovery
	TagFastRedeploy
)

var tagMapJa = map[Tag]string{
	TagVanguard:     "先鋒",
	TagGuard:        "前衛",
	TagDefender:     "重装",
	TagSniper:       "狙撃",
	TagCaster:       "術師",
	TagMedic:        "医療",
	TagSupporter:    "補助",
	TagSpecialist:   "特殊",
	TagMelee:        "近距離",
	TagRanged:       "遠距離",
	TagDPS:          "火力",
	TagDefense:      "防御",
	TagSurvival:     "生存",
	TagHealing:      "治療",
	TagSlow:         "減速",
	TagInitial:      "初期",
	TagCrowdControl: "牽制",
	TagSupport:      "支援",
	TagDebuff:       "弱化",
	TagSummon:       "召喚",
	TagElemental:    "元素",
	TagNuker:        "爆発力",
	TagRobot:        "ロボット",
	TagShift:        "強制移動",
	TagAoE:          "範囲攻撃",
	TagDPRecovery:   "COST回復",
	TagFastRedeploy: "高速再配置",
}

var tagMapEn = map[Tag]string{
	TagVanguard:     "Vanguard",
	TagGuard:        "Guard",
	TagDefender:     "Defender",
	TagSniper:       "Sniper",
	TagCaster:       "Caster",
	TagMedic:        "Medic",
	TagSupporter:    "Supporter",
	TagSpecialist:   "Specialist",
	TagMelee:        "Melee",
	TagRanged:       "Ranged",
	TagDPS:          "DPS",
	TagDefense:      "Defense",
	TagSurvival:     "Survival",
	TagHealing:      "Healing",
	TagSlow:         "Slow",
	TagInitial:      "Initial",
	TagCrowdControl: "Crowd-Control",
	TagSupport:      "Support",
	TagDebuff:       "Debuff",
	TagSummon:       "Summon",
	TagElemental:    "Elemental",
	TagNuker:        "Nuker",
	TagRobot:        "Robot",
	TagShift:        "Shift",
	TagAoE:          "AoE",
	TagDPRecovery:   "DP-Recovery",
	TagFastRedeploy: "Fast Redeploy",
}

func (t Tag) String() string {
	switch config.DefaultConfig.Language {
	case config.LanguageJapanese:
		return tagMapJa[t]
	default:
		return tagMapEn[t]
	}
}

func (t *Tag) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch config.DefaultConfig.Language {
	case config.LanguageJapanese:
		for k, v := range tagMapJa {
			if v == s {
				*t = k
				return nil
			}
		}
	default:
		for k, v := range tagMapEn {
			if v == s {
				*t = k
				return nil
			}
		}
	}
	return fmt.Errorf("unknown tag: %s", s)
}

func (t Tag) MarshalJSON() ([]byte, error) {
	switch config.DefaultConfig.Language {
	case config.LanguageJapanese:
		return json.Marshal(tagMapJa[t])
	default:
		return json.Marshal(tagMapEn[t])
	}
}
