package routes

import (
	"event-booking/handlers"

	"github.com/gin-gonic/gin"
)

func pingRoutes(router *gin.Engine) {
	router.GET("/ping", handlers.Ping)
}
