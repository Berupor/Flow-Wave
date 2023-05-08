package etl

import (
	"etl/models/clickhouse"
	"etl/models/mongodb"
	"strings"
)

type Transformer interface {
	TransformReviews(reviews []mongodb.Review) ([]clickhouse.Review, error)
}

type reviewTransformer struct {
}

func NewReviewTransformer() Transformer {
	return &reviewTransformer{}
}

func (rt *reviewTransformer) TransformReviews(reviews []mongodb.Review) ([]clickhouse.Review, error) {
	transformedReviews := make([]clickhouse.Review, len(reviews))

	for i, review := range reviews {
		transformedReview := clickhouse.Review{
			ProductID: review.ProductID,
			PlaceID:   review.PlaceID,
			AuthorID:  review.AuthorID,
			Rating:    review.Rating,
			Review:    cleanText(review.Review),
			Timestamp: review.Timestamp,
		}

		transformedReviews[i] = transformedReview
	}
	return transformedReviews, nil
}

func cleanText(text string) string {
	return strings.ReplaceAll(text, "\n", " ")
}
