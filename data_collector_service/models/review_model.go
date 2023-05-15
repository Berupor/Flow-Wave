package models

import (
	"errors"
	"fmt"
	"strings"
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

func (r *Review) Validate() error {
	fmt.Println(r.Timestamp)
	if r.ProductID <= 0 {
		return errors.New("product_id must be greater than 0")
	}
	if r.PlaceID <= 0 {
		return errors.New("place_id must be greater than 0")
	}
	if r.Rating < 1 || r.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	if len(strings.TrimSpace(r.Review)) < 1 {
		return errors.New("review must contain at least 1 symbol")
	}
	if r.Timestamp.IsZero() {
		return errors.New("timestamp cannot be empty")
	}

	return nil
}
