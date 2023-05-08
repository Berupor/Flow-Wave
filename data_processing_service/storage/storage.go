package storage

import (
	"data-processing/models"
)

type Storage interface {
	CreateReview(review models.Review) error
	FindReviews() ([]models.Review, error)
}
