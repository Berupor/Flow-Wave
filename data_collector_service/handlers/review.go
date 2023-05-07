package handlers

import (
	"errors"
	"net/http"

	"data-collector/models"
	"data-collector/services"

	"github.com/gin-gonic/gin"
)

func CreateReview(service services.EventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review models.Review
		userID, err := getUserID(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := c.BindJSON(&review); err != nil || review.Validate() != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "incorrect review data"})
			return
		}

		review.AuthorID = userID

		err = service.Produce(review)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusCreated, "Message sent")
	}
}

func getUserID(c *gin.Context) (int, error) {
	userIDInterface, exists := c.Get("user_id")

	if !exists {
		return 0, errors.New("user_id not found in context")
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		return 0, errors.New("user_id has invalid type")
	}

	return userID, nil
}
