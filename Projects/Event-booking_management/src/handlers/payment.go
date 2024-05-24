package handlers

import (
	"event-booking/constants"
	"event-booking/dtos"
	"event-booking/services/payments"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	req := &dtos.PaymentReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := payments.New().CreatePayment(getContext(c), req)
	// if err == constants.ErrBookingTaken {
	// 	c.JSON(http.StatusGone, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "payment created"})
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := payments.New().GetPayment(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrPaymentNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)

}

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	req := &dtos.PaymentReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := payments.New().UpdatePayment(getContext(c), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "payment updated"})
}

func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	err := payments.New().DeletePayment(getContext(c), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "payment deleted"})
}

func GetAllPayments(c *gin.Context) {
	payments, err := payments.New().GetAllPayments(getContext(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}
