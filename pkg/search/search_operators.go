package search

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/shima004/arknights-tag-checker/pkg/model"
)

//go:embed operators.json
var defaultOperatorsData []byte

type Filter interface {
	FilterOperator(operator *model.Operator) bool
}

type SearchOperator struct {
	Operators model.Operators
}

func NewSearchOperator() (*SearchOperator, error) {
	operators, err := loadOperator(defaultOperatorsData)
	if err != nil {
		return nil, fmt.Errorf("failed to load operators: %w", err)
	}

	return &SearchOperator{
		Operators: operators,
	}, nil
}

func loadOperator(operatorsData []byte) (model.Operators, error) {
	var operators model.Operators
	err := json.Unmarshal(operatorsData, &operators)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operators: %w", err)
	}

	return operators, nil
}

func (s *SearchOperator) SetOperatorsFromFile(data []byte) error {
	operators, err := loadOperator(data)
	if err != nil {
		return fmt.Errorf("failed to load operators: %w", err)
	}

	s.Operators = operators
	return nil
}

func (s *SearchOperator) SearchOperator(filters ...Filter) model.Operators {
	var result model.Operators
	for _, operator := range s.Operators {
		isValid := true
		for _, filter := range filters {
			if !filter.FilterOperator(operator) {
				isValid = false
				break
			}
		}
		if isValid {
			result = append(result, operator)
		}
	}
	return result
}
