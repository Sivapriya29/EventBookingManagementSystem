package context

import (
	"event-booking/dtos"
	"event-booking/utils/db"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	DB   db.DB
	User *dtos.User
}
