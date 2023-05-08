package mongodb

import "time"

type WeightSentence struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}

type Review struct {
	ProductID int       `json:"product_id"`
	PlaceID   int       `json:"place_id"`
	AuthorID  int       `json:"author_id"`
	Rating    float64   `json:"rating"`
	Review    string    `json:"review"`
	Timestamp time.Time `json:"timestamp"`

	Keywords  []string         `json:"keywords"`
	Sentences []WeightSentence `json:"sentences"`
	Sentiment float64          `json:"sentiment"`
}
