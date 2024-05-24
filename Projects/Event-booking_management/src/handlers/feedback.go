package handlers

import (
	"event-booking/constants"
	"event-booking/dtos"
	"event-booking/services/feedbacks"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFeedback(c *gin.Context) {
	req := &dtos.FeedbackReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := feedbacks.New().CreateFeedback(getContext(c), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Feedback created"})
}

func GetFeedback(c *gin.Context) {
	id := c.Param("id")

	feedback, err := feedbacks.New().GetFeedback(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if feedback == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrFeedbackNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, feedback)

}

func DeleteFeedback(c *gin.Context) {
	id := c.Param("id")
	err := feedbacks.New().DeleteFeedback(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Feedback deleted"})
}

func GetAllFeedbacks(c *gin.Context) {
	feedbacks, err := feedbacks.New().GetAllFeedbacks(getContext(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, feedbacks)
}

func GetFeedbacksByEventID(c *gin.Context) {
	eventID := c.Param("event_id")

	feedbacks, err := feedbacks.New().GetFeedbacksByEventID(getContext(c), eventID)
	if err != nil {
		if err.Error() == "event_id not found" || err.Error() == "invalid event_id format" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}

func GetFeedbacksByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	feedbacks, err := feedbacks.New().GetFeedbacksByUserID(getContext(c), userID)
	if err != nil {
		if err.Error() == "user_id not found" || err.Error() == "invalid user_id format" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}
