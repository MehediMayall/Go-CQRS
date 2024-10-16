package repositories

import "github.com/mehedimayall/go-cqrs/internal/entities"

type IMovieReadRepository interface {
	GetAll() (*[]entities.Movie, error)
	GetById(id string) (*entities.Movie, error)
}
