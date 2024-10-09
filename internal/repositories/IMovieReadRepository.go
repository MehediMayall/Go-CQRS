package repositories

import "github.com/mehedimayall/go-cqrs/internal/entities"

type IMovieReadRepository interface {
	GetMovies() ([]entities.Movie, error)
	GetById(id string) (entities.Movie, error)
}
