package service

import "github.com/ss49919201/diary/policy/model"

func CreatePolicy(userID int, typName string) (*model.Policy, error) {
	typ, err := (&model.Type{}).FindBy(map[string]any{"name": typName})
	if err != nil {
		return nil, err
	}
	policy := model.Policy{
		UserID: userID,
		Type:   typ.ID,
	}
	if err := policy.Save(); err != nil {
		return nil, err
	}
	return &policy, nil
}
