package user

import "github.com/doug-martin/goqu/v9"

type queryService struct{}

type SearchInput struct{}

type SearchOutput struct{}

func (q *queryService) Search(in SearchInput) (SearchOutput, error) {
	out := SearchOutput{}
	_, _, _ = goqu.From("users").ToSQL()
	return out, nil
}
