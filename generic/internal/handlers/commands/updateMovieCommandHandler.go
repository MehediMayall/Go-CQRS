package commands

import (
	"github.com/mehedimayall/go-cqrs/internal/entities"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type UpdateMovieCommandHandler struct {
	repo     repositories.IWriteRepository[entities.Movie]
	readRepo repositories.IReadRepository[entities.Movie]
}

func NewUpdateMovieCommand(
	repo repositories.IWriteRepository[entities.Movie],
	readRepo repositories.IReadRepository[entities.Movie]) *UpdateMovieCommandHandler {

	return &UpdateMovieCommandHandler{repo, readRepo}
}

func (handler *UpdateMovieCommandHandler) Handle(movie *entities.Movie) error {
	existingMovie, err := handler.readRepo.GetById(movie.Id)
	if err != nil {
		return err
	}

	existingMovie.Name = movie.Name
	existingMovie.Director = movie.Director
	existingMovie.Country = movie.Country
	existingMovie.Rating = movie.Rating

	return handler.repo.Update(existingMovie)
}
