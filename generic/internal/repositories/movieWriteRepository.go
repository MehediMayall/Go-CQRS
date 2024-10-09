package repositories

import (
	"log"

	"github.com/google/uuid"
	"github.com/mehedimayall/go-cqrs/internal/database"
	"github.com/mehedimayall/go-cqrs/internal/entities"
)

type MovieWriteRepository struct {
	db *database.InMemoryDB
}

func NewMovieWriteRepository(db *database.InMemoryDB) *MovieWriteRepository {
	return &MovieWriteRepository{
		db: db,
	}
}

// INSERT
func (repo *MovieWriteRepository) Add(movie *entities.Movie) error {
	movie.Id = uuid.NewString()
	repo.db.Movies[movie.Id] = *movie
	log.Println(repo.db)
	return nil
}

// UPDATE
func (repo *MovieWriteRepository) Update(movie *entities.Movie) error {
	repo.db.Movies[movie.Id] = *movie
	return nil
}

// DELETE
func (repo *MovieWriteRepository) Delete(movieId string) error {
	delete(repo.db.Movies, movieId)
	return nil
}
