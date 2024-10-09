package database

import "github.com/mehedimayall/go-cqrs/internal/entities"

type InMemoryDB struct {
	Movies map[string]entities.Movie
}

func NewDB() *InMemoryDB {
	return &InMemoryDB{
		Movies: map[string]entities.Movie{},
	}
}
