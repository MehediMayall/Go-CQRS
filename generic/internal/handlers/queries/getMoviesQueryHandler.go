package queries

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMoviesQueryHandler struct {
	repo repositories.IReadRepository[entities.Movie]
}

func NewGetMoviesQuery(repo repositories.IReadRepository[entities.Movie]) *GetMoviesQueryHandler {
	return &GetMoviesQueryHandler{repo}
}

func (handler *GetMoviesQueryHandler) Handle() (*[]entities.Movie, error) {
	return handler.repo.GetAll()
}
