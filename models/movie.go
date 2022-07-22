package models

import (
	"github.com/google/uuid"
)

type Movie struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Year       int       `json:"year"`
	ImdbRating float64   `json:"imdb_rating"`
	Director   string    `json:"director"`
	Released   bool      `json:"released"`
	Runtime    int       `json:"runtime"`
	Genre      string    `json:"genre"`
	Plot       string    `json:"plot"`
	Country    string    `json:"country"`
}

var MapOfMovies []map[string]Movie

func init() {
	MapOfMovies = make([]map[string]Movie, 0)
}
