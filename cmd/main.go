package main

import (
	"arknigths-tag-checker/pkg/filter"
	"arknigths-tag-checker/pkg/model"
	"arknigths-tag-checker/pkg/search"
	"fmt"
	"strings"
)

func main() {
	search, err := search.NewSearchOperator("data/operators.json")
	if err != nil {
		panic(err)
	}
	tags := []model.Tag{model.TagRanged, model.TagHealing, model.TagSurvival, model.TagAoE, model.TagShift}
	mutationTags := combinationsTags(tags)
	acquisitionMethods := []model.AcquisitionMethod{model.AMPublicRecruitment}
	for _, tags := range mutationTags {
		result := search.SearchOperator(
			true,
			&filter.TagFilter{Tags: tags, IsAnd: true},
			&filter.RarityFilter{Rarities: []int{5, 4, 3}},
			&filter.AcquisitionMethodFilter{AcquisitionMethods: acquisitionMethods, IsAnd: true},
		)
		if len(result) == 0 || !is_only_contains_rarities(result, []int{5, 4}) {
			continue
		}

		s := strings.Builder{}
		s.WriteString(fmt.Sprintf("Tags: %s\n", tags))
		for _, operator := range result {
			s.WriteString(fmt.Sprintf("%s\n", operator.Name))
		}
		fmt.Println(s.String())

	}

}

func is_only_contains_rarities(operators []*model.Operator, rarities []int) bool {
	for _, operator := range operators {
		if !contains(rarities, operator.Rarity) {
			return false
		}
	}
	return true
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func combinationsTags(tags []model.Tag) [][]model.Tag {
	var result [][]model.Tag
	var helper func([]model.Tag, int)
	helper = func(current []model.Tag, start int) {
		if len(current) > 0 {
			temp := make([]model.Tag, len(current))
			copy(temp, current)
			result = append(result, temp)
		}
		for i := start; i < len(tags); i++ {
			helper(append(current, tags[i]), i+1)
		}
	}
	helper([]model.Tag{}, 0)
	return result
}
