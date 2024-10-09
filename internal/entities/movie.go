package entities

import "time"

type Movie struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Director     string    `json:"director"`
	ReleasedDate time.Time `json:"releaseddate"`
	Country      string    `json:"country"`
	Rating       float32   `json:"rating"`
}
