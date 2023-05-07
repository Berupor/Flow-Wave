package handlers

import (
	"fmt"
	"net/http"

	"data-collector/models"
	"data-collector/services"

	"github.com/gin-gonic/gin"
)

func CreateReview(service services.EventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review models.Review

		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(review)

		err := service.Produce(review)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "Message sent")
	}
}
