package routes

import (
	"event-booking/handlers"
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func bookingRoutes(router *gin.RouterGroup) {
	router.POST("/bookings", middlewares.WithAuth(handlers.CreateBooking))
	router.GET("/bookings/:id", handlers.GetBooking)
	router.DELETE("/bookings/:id", handlers.DeleteBooking)
	router.GET("/bookings", handlers.GetAllBookings)
	router.GET("/event/:event_id/bookings", handlers.GetBookingsByEventID)
	router.GET("/user/:user_id/bookings", handlers.GetBookingsByUserID)
}
