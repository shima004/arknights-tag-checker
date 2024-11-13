package main

import (
	"arknigths-tag-checker/pkg/filter"
	"arknigths-tag-checker/pkg/model"
	"arknigths-tag-checker/pkg/search"
	"fmt"
)

func main() {
	search, err := search.NewSearchOperator("data/operators.json")
	if err != nil {
		panic(err)
	}
	tags := []model.Tag{model.TagVanguard}
	acquisitionMethods := []model.AcquisitionMethod{model.AMPublicRecruitment}
	result := search.SearchOperator(
		true,
		&filter.TagFilter{Tags: tags, IsAnd: false},
		&filter.RarityFilter{Rarities: []int{2}},
		&filter.AcquisitionMethodFilter{AcquisitionMethods: acquisitionMethods, IsAnd: true},
	)
	for k, v := range result {
		fmt.Printf("%d: %s\n", k, v.Name)
	}
}
