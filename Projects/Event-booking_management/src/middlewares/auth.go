package middlewares

import (
	"event-booking/constants"
	"event-booking/services/users"
	"event-booking/utils/context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuth(next func(*context.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getBearerToken(c)
		ctx := getContext(c)

		user, err := users.New().GetAccountWithToken(ctx, token)
		if err == constants.ErrAccessTokenExpired {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.User = user
		next(ctx)
	}
}

func getBearerToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	strs := strings.Split(bearerToken, " ")
	if len(strs) > 1 {
		if strs[0] != "Bearer" {
			return ""
		}
		return strs[1]
	}
	return ""
}
