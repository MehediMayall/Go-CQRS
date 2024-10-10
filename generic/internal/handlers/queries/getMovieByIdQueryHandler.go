package queries

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type GetMovieByIdQueryHandler struct {
	repo repositories.IReadRepository[entities.Movie]
}

func NewGetMovieByIdQuery(repo repositories.IReadRepository[entities.Movie]) *GetMovieByIdQueryHandler {
	return &GetMovieByIdQueryHandler{
		repo: repo,
	}
}

func (handler *GetMovieByIdQueryHandler) Handle(movieId string) (*entities.Movie, error) {
	return handler.repo.GetById(movieId)
}
