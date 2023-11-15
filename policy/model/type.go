package model

import (
	"fmt"
)

type Type struct {
	ID   int
	Name int
}

func (_ *Type) FindBy(m map[string]any) (*Type, error) {
	for k, v := range m {
		query := fmt.Sprintf("%s = ?", k)
		db = db.Where(query, v)
	}

	var result Type
	if err := db.First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
