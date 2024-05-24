package handlers

import (
	"event-booking/constants"
	"event-booking/dtos"
	"event-booking/services/events"
	"event-booking/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *context.Context) {
	req := &dtos.EventReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.User.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized to create Event. Only admin can create.",
		})
		return
	}

	err := events.New().CreateEvent(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "event created"})
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := events.New().GetEvent(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrEventNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, event)

}

func UpdateEvent(c *context.Context) {
	id := c.Param("id")
	req := &dtos.EventReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.User.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized to update Event. Only admin can update",
		})
		return
	}

	err := events.New().UpdateEvent(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "event updated"})
}

func DeleteEvent(c *context.Context) {
	id := c.Param("id")

	if c.User.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized to update Event. Only admin can delete",
		})
		return
	}

	err := events.New().DeleteEvent(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "event deleted"})
}

func GetAllEvents(c *gin.Context) {
	events, err := events.New().GetAllEvents(getContext(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
