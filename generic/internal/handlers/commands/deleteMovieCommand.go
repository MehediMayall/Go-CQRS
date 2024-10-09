package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type DeleteMovieCommand struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewDeleteMovieCommand(repo repositories.IWriteRepository[entities.Movie]) *DeleteMovieCommand {
	return &DeleteMovieCommand{
		repo: repo,
	}
}

func (handler *DeleteMovieCommand) Handle(movieId string) error {
	return handler.repo.Delete(movieId)
}
