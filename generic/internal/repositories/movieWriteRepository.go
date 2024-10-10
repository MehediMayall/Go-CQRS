// MovieWriteRepository has only one responsibility that is to write data to the database.
// It is used to transfer data from application to database in more generic ways.

package repositories

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mehedimayall/go-cqrs/internal/database"
	"github.com/mehedimayall/go-cqrs/internal/entities"
)

type MovieWriteRepository struct {
	db *database.InMemoryDB
}

func NewMovieWriteRepository(db *database.InMemoryDB) *MovieWriteRepository {
	return &MovieWriteRepository{db}
}

// INSERT
func (repo *MovieWriteRepository) Add(movie *entities.Movie) error {
	movie.Id = uuid.NewString()
	movie.CreatedOn = time.Now()

	repo.db.Movies[movie.Id] = *movie
	log.Println(repo.db)
	return nil
}

// UPDATE
func (repo *MovieWriteRepository) Update(movie *entities.Movie) error {
	movie.UpdatedOn = time.Now()
	repo.db.Movies[movie.Id] = *movie
	return nil
}

// DELETE
func (repo *MovieWriteRepository) Delete(movieId string) error {
	delete(repo.db.Movies, movieId)
	return nil
}
