package main

import (
	"fmt"

	"github.com/shima004/arknights-tag-checker/pkg/filter"
	"github.com/shima004/arknights-tag-checker/pkg/model"
	"github.com/shima004/arknights-tag-checker/pkg/search"
)

func main() {
	search, err := search.NewSearchOperator()
	if err != nil {
		panic(err)
	}
	// Filter by tags, rarity, and acquisition method
	tags := model.Tags{
		model.TagDPS, model.TagDefense, model.TagHealing, model.TagVanguard, model.TagSlow,
	}
	rarities := []int{5, 4, 3}
	acquisitionMethods := []model.AcquisitionMethod{model.AMPublicRecruitment}
	// Search for all combinations
	mutationTags := tags.CombinationsTags()
	for _, tags := range mutationTags {
		operators := search.SearchOperator(
			&filter.TagFilter{Tags: tags, IsAnd: true},
			&filter.RarityFilter{Rarities: rarities},
			&filter.AcquisitionMethodFilter{AcquisitionMethods: acquisitionMethods, IsAnd: true},
		)

		if len(operators) != 0 {
			fmt.Printf("%v : %v\n", tags, operators)
		}
	}
}
