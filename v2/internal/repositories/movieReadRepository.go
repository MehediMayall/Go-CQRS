package repositories

import (
	"errors"

	"github.com/mehedimayall/go-cqrs/internal/database"
	"github.com/mehedimayall/go-cqrs/internal/entities"
)

type MovieReadRepository struct {
	db *database.InMemoryDB
}

func NewMovieReadRepository(db *database.InMemoryDB) MovieReadRepository {
	return MovieReadRepository{
		db: db,
	}
}

func (r *MovieReadRepository) GetAll() ([]entities.Movie, error) {
	movies := []entities.Movie{}
	for _, value := range r.db.Movies {
		movies = append(movies, value)
	}
	return movies, nil
}

func (r *MovieReadRepository) GetById(movieId string) (entities.Movie, error) {

	movie := r.db.Movies[movieId]
	if &movie == nil {
		return movie, errors.New("There is no movie exists with Id: " + movieId)
	}
	return movie, nil
}
