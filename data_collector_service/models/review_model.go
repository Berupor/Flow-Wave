package models

import "time"

type Review struct {
	Product_id int       `json:"product_id"`
	Place_id   int       `json:"place_id"`
	Author_id  int       `json:"author_id"`
	Rating     float64   `json:"rating"`
	Review     string    `json:"review"`
	Timestamp  time.Time `json:"timestamp"`
}
