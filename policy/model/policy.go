package model

import "fmt"

type Policy struct {
	ID     int
	UserID int
	Type   int
}

func (p *Policy) Save() error {
	return db.Save(p).Error
}

func (_ *Policy) FindBy(m map[string]any) (*Policy, error) {
	for k, v := range m {
		query := fmt.Sprintf("%s = ?", k)
		db = db.Where(query, v)
	}

	var result Policy
	if err := db.First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
