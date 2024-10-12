// In this struct we have two interactions with the database.
// First, we have to retrieve the movie from the database.
// Second, we have to replace the old movie with the new one and update the movie in the database.
// This has to be done in loosely coupled way.

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
