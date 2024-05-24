package handlers

import (
	"event-booking/constants"
	"event-booking/dtos"
	"event-booking/services/users"
	"event-booking/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	req := &dtos.RegisterReq{}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := users.New().Register(getContext(c), req)
	if err == constants.ErrEmailTaken || err == constants.ErrMobileTaken {
		c.JSON(http.StatusGone, gin.H{
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

	c.JSON(http.StatusCreated, gin.H{
		"msg": "created",
	})
}

func Login(c *gin.Context) {
	req := &dtos.LoginReq{}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := users.New().Login(getContext(c), req)
	if err == constants.ErrInvalidEmailOrPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetAccountUsingToken(c *context.Context) {

	c.JSON(http.StatusOK, c.User)
}

func GetAccessTokenFromRefreshToken(c *gin.Context) {
	accessToken, err := users.New().GetAccessTokenFromRefreshToken(getContext(c), c.Query("refresh_token"))
	if err == constants.ErrRefreshTokenExpired {
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
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
