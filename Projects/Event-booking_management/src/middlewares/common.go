package middlewares

import (
	"event-booking/utils/context"
	"event-booking/utils/db"

	"github.com/gin-gonic/gin"
)

func getContext(c *gin.Context) *context.Context {
	return &context.Context{
		Context: c,
		DB:      *db.New(),
	}
}
