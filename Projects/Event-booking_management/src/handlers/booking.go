package handlers

import (
	"event-booking/constants"
	"event-booking/dtos"
	"event-booking/services/bookings"
	"event-booking/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *context.Context) {
	req := &dtos.BookingReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.User.Role != constants.RoleUser {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized to book an event. Only users can book event",
		})
		return
	}

	booking, err := bookings.New().CreateBooking(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Booking created", "booking": booking})
}

func GetBooking(c *gin.Context) {
	id := c.Param("id")

	booking, err := bookings.New().GetBooking(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrBookingNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)

}

func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	err := bookings.New().DeleteBooking(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Booking deleted"})
}

func GetAllBookings(c *gin.Context) {
	bookings, err := bookings.New().GetAllBookings(getContext(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func GetBookingsByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	bookings, err := bookings.New().GetBookingsByUserID(getContext(c), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func GetBookingsByEventID(c *gin.Context) {
	eventID := c.Param("event_id")

	bookings, err := bookings.New().GetBookingsByEventID(getContext(c), eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
