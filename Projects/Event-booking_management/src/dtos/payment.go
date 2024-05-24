package dtos

import "time"

type PaymentReq struct {
	ID           string
	Booking_id   string  `json:"booking_id" binding:"required"`
	User_id      string  `json:"user_id" binding:"required"`
	Event_id     string  `json:"event_id" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	Card_number  string  `json:"card_number" binding:"required"`
	Expiry_month string  `json:"expiry_month" binding:"required"`
	Expiry_year  string  `json:"expiry_year" binding:"required"`
	Cvv          string  `json:"cvv" binding:"required"`
	Card_holder  string  `json:"card_holder" binding:"required"`
}

type PaymentRes struct {
	ID           string    `json:"id"`
	Booking_id   string    `json:"booking_id"`
	User_id      string    `json:"user_id"`
	Event_id     string    `json:"event_id"`
	Amount       float64   `json:"amount"`
	Payment_date time.Time `json:"payment_date"`
}
