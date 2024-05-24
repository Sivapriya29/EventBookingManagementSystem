package routes

import (
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	pingRoutes(router)
	v1 := router.Group("/v1")
	v1.Use(middlewares.CORSMiddleware())
	userRoutes(v1)
	eventRoutes(v1)
	bookingRoutes(v1)
	feedbackRoutes(v1)
	paymentRoutes(v1)

	return router
}
