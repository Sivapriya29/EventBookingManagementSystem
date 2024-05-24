package routes

import (
	"event-booking/handlers"

	"github.com/gin-gonic/gin"
)

func feedbackRoutes(router *gin.RouterGroup) {
	router.POST("/feedbacks", handlers.CreateFeedback)
	router.GET("/feedbacks/:id", handlers.GetFeedback)
	router.GET("/event/:event_id/feedbacks", handlers.GetFeedbacksByEventID)
	router.GET("/user/:user_id/feedbacks", handlers.GetFeedbacksByUserID)
	router.DELETE("/feedbacks/:id", handlers.DeleteFeedback)
	router.GET("/feedbacks", handlers.GetAllFeedbacks)
}
