package filter_test

import (
	"testing"

	"github.com/shima004/arknights-tag-checker/pkg/filter"
	"github.com/shima004/arknights-tag-checker/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestTagFilter_FilterOperator(t *testing.T) {
	t.Parallel()
	t.Run("Test OR tag filter", func(t *testing.T) {
		t.Parallel()
		t.Run("Test signle tag", func(t *testing.T) {
			t.Parallel()
			tagFilter := &filter.TagFilter{
				Tags:  []model.Tag{model.TagVanguard},
				IsAnd: false,
			}
			t.Run("Test if operator has the tag return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagVanguard},
				}
				assert.True(t, tagFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the tag return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagGuard},
				}
				assert.False(t, tagFilter.FilterOperator(operator))
			})
		})

		t.Run("Test multiple tags", func(t *testing.T) {
			t.Parallel()
			tagFilter := &filter.TagFilter{
				Tags:  []model.Tag{model.TagVanguard, model.TagDPRecovery},
				IsAnd: false,
			}
			t.Run("Test if operator has the tags return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagVanguard, model.TagDPS},
				}
				assert.True(t, tagFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the tags return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagGuard, model.TagDPS},
				}
				assert.False(t, tagFilter.FilterOperator(operator))
			})
		})
	})

	t.Run("Test AND tag filter", func(t *testing.T) {
		t.Parallel()
		t.Run("Test signle tag", func(t *testing.T) {
			t.Parallel()
			tagFilter := &filter.TagFilter{
				Tags:  []model.Tag{model.TagVanguard},
				IsAnd: true,
			}
			t.Run("Test if operator has the tag return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagVanguard},
				}
				assert.True(t, tagFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the tag return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagGuard},
				}
				assert.False(t, tagFilter.FilterOperator(operator))
			})
		})

		t.Run("Test multiple tags", func(t *testing.T) {
			t.Parallel()
			tagFilter := &filter.TagFilter{
				Tags:  []model.Tag{model.TagVanguard, model.TagDPRecovery},
				IsAnd: true,
			}
			t.Run("Test if operator has the tags return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagVanguard, model.TagDPRecovery},
				}
				assert.True(t, tagFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the tags return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					Tags: []model.Tag{model.TagVanguard, model.TagDPS},
				}
				assert.False(t, tagFilter.FilterOperator(operator))
			})
		})
	})
}

func TestRarityFilter_FilterOperator(t *testing.T) {
	t.Parallel()
	rarityFilter := &filter.RarityFilter{
		Rarities: []int{2},
	}
	t.Run("Test if operator has the rarity return true", func(t *testing.T) {
		t.Parallel()
		operator := &model.Operator{
			Rarity: 2,
		}
		assert.True(t, rarityFilter.FilterOperator(operator))
	})
	t.Run("Test if operator doesn't have the rarity return false", func(t *testing.T) {
		t.Parallel()
		operator := &model.Operator{
			Rarity: 3,
		}
		assert.False(t, rarityFilter.FilterOperator(operator))
	})
}

func TestAcquisitionMethodFilter_FilterOperator(t *testing.T) {
	t.Parallel()
	t.Run("Test OR acquisition method filter", func(t *testing.T) {
		t.Parallel()
		t.Run("Test single acquisition method", func(t *testing.T) {
			t.Parallel()
			acquisitionMethodFilter := &filter.AcquisitionMethodFilter{
				AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment},
				IsAnd:              false,
			}
			t.Run("Test if operator has the acquisition method return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment},
				}
				assert.True(t, acquisitionMethodFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the acquisition method return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMLimitedHeadHuntingFestival},
				}
				assert.False(t, acquisitionMethodFilter.FilterOperator(operator))
			})
		})

		t.Run("Test multiple acquisition methods", func(t *testing.T) {
			t.Parallel()
			acquisitionMethodFilter := &filter.AcquisitionMethodFilter{
				AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment, model.AMStandardHeadHunting},
				IsAnd:              false,
			}
			t.Run("Test if operator has the acquisition methods return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment, model.AMLimitedHeadHuntingFestival},
				}
				assert.True(t, acquisitionMethodFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the acquisition methods return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMLimitedHeadHuntingFestival},
				}
				assert.False(t, acquisitionMethodFilter.FilterOperator(operator))
			})
		})

	})

	t.Run("Test AND acquisition method filter", func(t *testing.T) {
		t.Parallel()
		t.Run("Test single acquisition method", func(t *testing.T) {
			t.Parallel()
			acquisitionMethodFilter := &filter.AcquisitionMethodFilter{
				AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment},
				IsAnd:              true,
			}
			t.Run("Test if operator has the acquisition method return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment},
				}
				assert.True(t, acquisitionMethodFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the acquisition method return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMLimitedHeadHuntingFestival},
				}
				assert.False(t, acquisitionMethodFilter.FilterOperator(operator))
			})
		})

		t.Run("Test multiple acquisition methods", func(t *testing.T) {
			t.Parallel()
			acquisitionMethodFilter := &filter.AcquisitionMethodFilter{
				AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment, model.AMStandardHeadHunting},
				IsAnd:              true,
			}
			t.Run("Test if operator has the acquisition methods return true", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment, model.AMStandardHeadHunting},
				}
				assert.True(t, acquisitionMethodFilter.FilterOperator(operator))
			})
			t.Run("Test if operator doesn't have the acquisition methods return false", func(t *testing.T) {
				t.Parallel()
				operator := &model.Operator{
					AcquisitionMethods: []model.AcquisitionMethod{model.AMPublicRecruitment},
				}
				assert.False(t, acquisitionMethodFilter.FilterOperator(operator))
			})
		})
	})
}
