package search

import (
	"arknights-tag-checker/pkg/model"
	"encoding/json"
	"io"
	"os"
)

type Filter interface {
	FilterOperator(operator *model.Operator) bool
}

type SearchOperator struct {
	Operators model.Operators
}

func NewSearchOperator(filePath string) (*SearchOperator, error) {
	operators, err := loadOperator(filePath)
	if err != nil {
		return nil, err
	}
	return &SearchOperator{
		Operators: operators,
	}, nil
}

func loadOperator(filePath string) (model.Operators, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var operators model.Operators
	err = json.Unmarshal(byteValue, &operators)
	if err != nil {
		return nil, err
	}

	return operators, nil
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
