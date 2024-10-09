package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type AddMovieCommand struct {
	repo repositories.IWriteRepository
}

func NewAddMovieCommand(repo repositories.IWriteRepository) AddMovieCommand {
	return AddMovieCommand{
		repo: repo,
	}
}

func (handler *AddMovieCommand) Handle(movie *entities.Movie) error {
	return handler.repo.Add(movie)
}
