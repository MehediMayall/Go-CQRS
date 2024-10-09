package queries

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type GetMoviesQuery struct {
	repo repositories.IMovieReadRepository
}

func NewGetMoviesQuery(repo repositories.IMovieReadRepository) *GetMoviesQuery {
	return &GetMoviesQuery{
		repo: repo,
	}
}

func (handler *GetMoviesQuery) Handle() (*[]entities.Movie, error) {
	return handler.repo.GetAll()
}
