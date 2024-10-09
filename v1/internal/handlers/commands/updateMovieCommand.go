package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type UpdateMovieCommand struct {
	repo repositories.IMovieWriteRepository
}

func NewUpdateMovieCommand(repo repositories.IMovieWriteRepository) *UpdateMovieCommand {
	return &UpdateMovieCommand{
		repo: repo,
	}
}

func (handler *UpdateMovieCommand) Handle(movie *entities.Movie) error {
	return handler.repo.Update(movie)
}
