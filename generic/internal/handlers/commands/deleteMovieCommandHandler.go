package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type DeleteMovieCommandHandler struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewDeleteMovieCommand(repo repositories.IWriteRepository[entities.Movie]) *DeleteMovieCommandHandler {
	return &DeleteMovieCommandHandler{repo}
}

func (handler *DeleteMovieCommandHandler) Handle(movieId string) error {
	return handler.repo.Delete(movieId)
}
