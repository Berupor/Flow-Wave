package storage

import (
	"data-processing/models"
)

type Storage interface {
	CreateReview(review models.ReviewCreate) error
	FindReviews() ([]models.Review, error)
}
