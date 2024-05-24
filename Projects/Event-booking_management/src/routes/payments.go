package routes

import (
	"event-booking/handlers"

	"github.com/gin-gonic/gin"
)

func paymentRoutes(router *gin.RouterGroup) {
	router.POST("/payments", handlers.CreatePayment)
	router.GET("/payments/:id", handlers.GetPayment)
	router.PUT("/payments/:id", handlers.UpdatePayment)
	router.DELETE("/payments/:id", handlers.DeletePayment)
	router.GET("/payments", handlers.GetAllPayments)
}
