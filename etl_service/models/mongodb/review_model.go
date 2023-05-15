package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type WeightSentence struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}

type Review struct {
	ID        primitive.ObjectID `bson:"_id"`
	ProductID int                `json:"product_id"`
	PlaceID   int                `json:"place_id"`
	AuthorID  int                `json:"author_id"`
	Rating    float64            `json:"rating"`
	Review    string             `json:"review"`
	Timestamp time.Time          `json:"timestamp"`

	Keywords  []string         `json:"keywords"`
	Sentences []WeightSentence `json:"sentences"`
	Sentiment float64          `json:"sentiment"`
}
