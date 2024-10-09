package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/repositories"
)

type AddMovieCommand struct {
	repo repositories.IMovieWriteRepository
}

func NewAddMovieCommand(repo repositories.IMovieWriteRepository) *AddMovieCommand {
	return &AddMovieCommand{
		repo: repo,
	}
}

func (handler *AddMovieCommand) Handle(movie *entities.Movie) error {
	return handler.repo.Add(movie)
}
