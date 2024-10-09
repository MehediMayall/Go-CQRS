package queries

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMoviesQuery struct {
	repo repositories.IReadRepository[entities.Movie]
}

func NewGetMoviesQuery(repo repositories.IReadRepository[entities.Movie]) *GetMoviesQuery {
	return &GetMoviesQuery{
		repo: repo,
	}
}

func (handler *GetMoviesQuery) Handle() (*[]entities.Movie, error) {
	return handler.repo.GetAll()
}
