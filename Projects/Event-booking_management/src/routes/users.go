package routes

import (
	"event-booking/handlers"
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	router.GET("/users", middlewares.WithAuth(handlers.GetAccountUsingToken))
	router.GET("/access-token", handlers.GetAccessTokenFromRefreshToken)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
