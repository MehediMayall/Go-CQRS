package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type GetMovieByIdQuery struct {
	repo repositories.IMovieReadRepository
}

func NewGetMovieByIdQuery(repo repositories.IMovieReadRepository) *GetMovieByIdQuery {
	return &GetMovieByIdQuery{
		repo: repo,
	}
}

func (handler *GetMovieByIdQuery) Handle(movieId string) (*entities.Movie, error) {
	return handler.repo.GetById(movieId)
}
