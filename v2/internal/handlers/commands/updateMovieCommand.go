package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type UpdateMovieCommand struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewUpdateMovieCommand(repo repositories.IWriteRepository[entities.Movie]) *UpdateMovieCommand {
	return &UpdateMovieCommand{
		repo: repo,
	}
}

func (handler *UpdateMovieCommand) Handle(movie *entities.Movie) error {
	return handler.repo.Update(movie)
}
