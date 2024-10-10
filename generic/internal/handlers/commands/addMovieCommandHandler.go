package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type AddMovieCommandHandler struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewAddMovieCommand(repo repositories.IWriteRepository[entities.Movie]) AddMovieCommandHandler {
	return AddMovieCommandHandler{repo}
}

func (handler *AddMovieCommandHandler) Handle(movie *entities.Movie) (string, error) {
	handler.repo.Add(movie)
	return movie.Id, nil
}
