package model_test

import (
	"testing"

	"github.com/shima004/arknights-tag-checker/pkg/model"
)

func TestAcquisitionMethod_MarshalJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		method  model.AcquisitionMethod
		want    string
		wantErr bool
	}{
		{
			name:    "StandardHeadHunting",
			method:  model.AMStandardHeadHunting,
			want:    `"スタンダードスカウト"`,
			wantErr: false,
		},
		{
			name:    "KernelHeadHunting",
			method:  model.AMKernelHeadHunting,
			want:    `"中堅スカウト"`,
			wantErr: false,
		},
		{
			name:    "LimitedHeadHuntingCelebration",
			method:  model.AMLimitedHeadHuntingCelebration,
			want:    `"リミテッドスカウト[祭]"`,
			wantErr: false,
		},
		{
			name:    "LimitedHeadHuntingFestival",
			method:  model.AMLimitedHeadHuntingFestival,
			want:    `"リミテッドスカウト[宴]"`,
			wantErr: false,
		},
		{
			name:    "LimitedHeadHuntingCarvinal",
			method:  model.AMLimitedHeadHuntingCarnival,
			want:    `"リミテッドスカウト[遊]"`,
			wantErr: false,
		},
		{
			name:    "PublicRecruitment",
			method:  model.AMPublicRecruitment,
			want:    `"公開求人"`,
			wantErr: false,
		},
		{
			name:    "CrossOverHeadHunting",
			method:  model.AMCrossOverHeadHunting,
			want:    `"コラボスカウト"`,
			wantErr: false,
		},
		{
			name:    "CrossOverEvent",
			method:  model.AMCrossOverEvent,
			want:    `"コラボイベント報酬"`,
			wantErr: false,
		},
		{
			name:    "Event",
			method:  model.AMEvent,
			want:    `"イベント報酬"`,
			wantErr: false,
		},
		{
			name:    "MainStory",
			method:  model.AMMainStory,
			want:    `"メインストーリー"`,
			wantErr: false,
		},
		{
			name:    "IntegratedStrategies",
			method:  model.AMIntegratedStrategies,
			want:    `"統合戦略"`,
			wantErr: false,
		},
		{
			name:    "PurchaseCertificateStore",
			method:  model.AMPurchaseCertificateStore,
			want:    `"購買資格証"`,
			wantErr: false,
		},
		{
			name:    "Anniversary",
			method:  model.AMAnniversary,
			want:    `"周年記念"`,
			wantErr: false,
		},
		{
			name:    "ReclamationAlgorithm",
			method:  model.AMReclamationAlgorithm,
			want:    `"生息演算"`,
			wantErr: false,
		},
		{
			name:    "ContingencyContract",
			method:  model.AMContingencyContract,
			want:    `"危機契約"`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := tt.method.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("AcquisitionMethod.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(b) != tt.want {
				t.Errorf("AcquisitionMethod.MarshalJSON() = %v, want %v", string(b), tt.want)
			}
		})
	}
}

func TestAcquisitionMethod_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		stringMethod string
		want         model.AcquisitionMethod
		wantErr      bool
	}{
		{
			name:         "StandardHeadHunting",
			stringMethod: `"スタンダードスカウト"`,
			want:         model.AMStandardHeadHunting,
			wantErr:      false,
		},
		{
			name:         "KernelHeadHunting",
			stringMethod: `"中堅スカウト"`,
			want:         model.AMKernelHeadHunting,
			wantErr:      false,
		},
		{
			name:         "LimitedHeadHuntingCelebration",
			stringMethod: `"リミテッドスカウト[祭]"`,
			want:         model.AMLimitedHeadHuntingCelebration,
			wantErr:      false,
		},
		{
			name:         "LimitedHeadHuntingFestival",
			stringMethod: `"リミテッドスカウト[宴]"`,
			want:         model.AMLimitedHeadHuntingFestival,
			wantErr:      false,
		},
		{
			name:         "LimitedHeadHuntingCarvinal",
			stringMethod: `"リミテッドスカウト[遊]"`,
			want:         model.AMLimitedHeadHuntingCarnival,
			wantErr:      false,
		},
		{
			name:         "PublicRecruitment",
			stringMethod: `"公開求人"`,
			want:         model.AMPublicRecruitment,
			wantErr:      false,
		},
		{
			name:         "CrossOverHeadHunting",
			stringMethod: `"コラボスカウト"`,
			want:         model.AMCrossOverHeadHunting,
			wantErr:      false,
		},
		{
			name:         "CrossOverEvent",
			stringMethod: `"コラボイベント報酬"`,
			want:         model.AMCrossOverEvent,
			wantErr:      false,
		},
		{
			name:         "Event",
			stringMethod: `"イベント報酬"`,
			want:         model.AMEvent,
			wantErr:      false,
		},
		{
			name:         "MainStory",
			stringMethod: `"メインストーリー"`,
			want:         model.AMMainStory,
			wantErr:      false,
		},
		{
			name:         "IntegratedStrategies",
			stringMethod: `"統合戦略"`,
			want:         model.AMIntegratedStrategies,
			wantErr:      false,
		},
		{
			name:         "PurchaseCertificateStore",
			stringMethod: `"購買資格証"`,
			want:         model.AMPurchaseCertificateStore,
			wantErr:      false,
		},
		{
			name:         "Anniversary",
			stringMethod: `"周年記念"`,
			want:         model.AMAnniversary,
			wantErr:      false,
		},
		{
			name:         "ReclamationAlgorithm",
			stringMethod: `"生息演算"`,
			want:         model.AMReclamationAlgorithm,
			wantErr:      false,
		},
		{
			name:         "ContingencyContract",
			stringMethod: `"危機契約"`,
			want:         model.AMContingencyContract,
			wantErr:      false,
		},
		{
			name:         "Invalid",
			stringMethod: `""`,
			want:         model.AcquisitionMethod(0),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var method model.AcquisitionMethod
			err := method.UnmarshalJSON([]byte(tt.stringMethod))
			if (err != nil) != tt.wantErr {
				t.Errorf("AcquisitionMethod.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if method != tt.want {
				t.Errorf("AcquisitionMethod.UnmarshalJSON() = %v, want %v", method, tt.want)
			}
		})
	}
}
