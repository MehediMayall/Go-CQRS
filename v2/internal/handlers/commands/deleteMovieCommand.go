package commands

import (
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type DeleteMovieCommand struct {
	repo repositories.IWriteRepository
}

func NewDeleteMovieCommand(repo repositories.IWriteRepository) *DeleteMovieCommand {
	return &DeleteMovieCommand{
		repo: repo,
	}
}

func (handler *DeleteMovieCommand) Handle(movieId string) error {
	return handler.repo.Delete(movieId)
}
