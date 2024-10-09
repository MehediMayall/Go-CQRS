package queries

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMovieByIdQuery struct {
	repo repositories.IReadRepository[entities.Movie]
}

func NewGetMovieByIdQuery(repo repositories.IReadRepository[entities.Movie]) *GetMovieByIdQuery {
	return &GetMovieByIdQuery{
		repo: repo,
	}
}

func (handler *GetMovieByIdQuery) Handle(movieId string) (*entities.Movie, error) {
	return handler.repo.GetById(movieId)
}
