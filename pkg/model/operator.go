package model

import (
	"fmt"
)

type Operator struct {
	Name               string             `json:"name"`
	Rarity             int                `json:"rarity"`
	Tags               Tags               `json:"tags"`
	AcquisitionMethods AcquisitionMethods `json:"acquisition_methods"`
}

type Operators []*Operator

func (o Operators) IsOnlyContainsRarities(rarities []int) bool {
	for _, operator := range o {
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

func (o Operator) String() string {
	return fmt.Sprintf("%s : ★ %d", o.Name, o.Rarity)
}
