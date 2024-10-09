package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type DeleteMovieCommand struct {
	repo repositories.IMovieWriteRepository
}

func NewDeleteMovieCommand(repo repositories.IMovieWriteRepository) *DeleteMovieCommand {
	return &DeleteMovieCommand{
		repo: repo,
	}
}

func (handler *DeleteMovieCommand) Handle(movieId string) error {
	return handler.repo.Delete(movieId)
}
