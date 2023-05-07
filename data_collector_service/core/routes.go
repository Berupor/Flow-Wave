package core

import (
	"data-collector/handlers"
	"data-collector/middlewares"
	"data-collector/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(eventService *services.EventService, jwtSecret string) *gin.Engine {
	router := gin.Default()

	secured := router.Group("api")
	secured.Use(middlewares.JWTAuthMiddleware(jwtSecret))
	{
		secured.POST("/review", handlers.CreateReview(*eventService))
	}

	return router
}
