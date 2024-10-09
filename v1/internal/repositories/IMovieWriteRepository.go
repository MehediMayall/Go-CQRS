package repositories

import "github.com/mehedimayall/go-cqrs/internal/entities"

type IMovieWriteRepository interface {
	Add(*entities.Movie) error
	Update(*entities.Movie) error
	Delete(movieId string) error
}
