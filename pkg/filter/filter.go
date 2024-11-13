package filter

import (
	"arknigths-tag-checker/pkg/model"
)

type TagFilter struct {
	Tags  []model.Tag
	IsAnd bool
}

func (f *TagFilter) FilterOperator(operator *model.Operator) bool {
	if f.IsAnd {
		return containsAll(operator.Tags, f.Tags)
	}
	return containsAny(operator.Tags, f.Tags)
}

type RarityFilter struct {
	Rarities []int
}

func (f *RarityFilter) FilterOperator(operator *model.Operator) bool {
	for _, rarity := range f.Rarities {
		if operator.Rarity == rarity {
			return true
		}
	}
	return false
}

type AcquisitionMethodFilter struct {
	AcquisitionMethods []model.AcquisitionMethod
	IsAnd              bool
}

func (f *AcquisitionMethodFilter) FilterOperator(operator *model.Operator) bool {
	if f.IsAnd {
		return containsAll(operator.AcquisitionMethods, f.AcquisitionMethods)
	}
	return containsAny(operator.AcquisitionMethods, f.AcquisitionMethods)
}

func containsAll[T comparable](slice []T, items []T) bool {
	for _, item := range items {
		found := false
		for _, s := range slice {
			if s == item {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func containsAny[T comparable](slice []T, items []T) bool {
	for _, item := range items {
		for _, s := range slice {
			if s == item {
				return true
			}
		}
	}
	return false
}
