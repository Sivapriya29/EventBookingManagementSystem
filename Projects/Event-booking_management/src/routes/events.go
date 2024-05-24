package routes

import (
	"event-booking/handlers"
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func eventRoutes(router *gin.RouterGroup) {
	router.POST("/events", middlewares.WithAuth(handlers.CreateEvent))
	router.GET("/events/:id", handlers.GetEvent)
	router.PUT("/events/:id", middlewares.WithAuth(handlers.UpdateEvent))
	router.DELETE("/events/:id", middlewares.WithAuth(handlers.DeleteEvent))
	router.GET("/events", handlers.GetAllEvents)
}
