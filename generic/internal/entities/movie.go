package entities

type Movie struct {
	EntityBase
	Name     string  `json:"name"`
	Director string  `json:"director"`
	Country  string  `json:"country"`
	Rating   float32 `json:"rating"`
}
