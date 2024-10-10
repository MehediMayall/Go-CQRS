package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type UpdateMovieCommandHandler struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewUpdateMovieCommand(repo repositories.IWriteRepository[entities.Movie]) *UpdateMovieCommandHandler {
	return &UpdateMovieCommandHandler{
		repo: repo,
	}
}

func (handler *UpdateMovieCommandHandler) Handle(movie *entities.Movie) error {
	return handler.repo.Update(movie)
}
