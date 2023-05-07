package models

import (
	"time"
)

type Review struct {
	ProductID int       `json:"product_id"`
	PlaceID   int       `json:"place_id"`
	AuthorID  int       `json:"author_id"`
	Rating    float64   `json:"rating"`
	Review    string    `json:"review"`
	Timestamp time.Time `json:"timestamp"`
}
