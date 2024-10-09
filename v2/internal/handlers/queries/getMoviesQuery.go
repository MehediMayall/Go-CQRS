package queries

import (
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMoviesQuery struct {
	repo repositories.IReadRepository
}

func NewGetMoviesQuery(repo repositories.IReadRepository) *GetMoviesQuery {
	return &GetMoviesQuery{
		repo: repo,
	}
}

func (handler *GetMoviesQuery) Handle() (*[]interface{}, error) {
	return handler.repo.GetAll()
}
