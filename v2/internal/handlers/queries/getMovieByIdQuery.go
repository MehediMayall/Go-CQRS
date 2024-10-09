package queries

import (
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMovieByIdQuery struct {
	repo repositories.IReadRepository
}

func NewGetMovieByIdQuery(repo repositories.IReadRepository) *GetMovieByIdQuery {
	return &GetMovieByIdQuery{
		repo: repo,
	}
}

func (handler *GetMovieByIdQuery) Handle(movieId string) (*interface{}, error) {
	return handler.repo.GetById(movieId)
}
