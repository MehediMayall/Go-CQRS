package entities

import "time"

type EntityBase struct {
	Id        string    `json:"id"`
	CreatedOn time.Time `json:"created"`
	UpdatedOn time.Time `json:"updated"`
}
