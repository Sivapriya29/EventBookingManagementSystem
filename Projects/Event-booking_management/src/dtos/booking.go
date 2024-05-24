package dtos

import "time"

type BookingReq struct {
	ID                string
	Event_id          string  `json:"event_id" binding:"required"`
	User_id           string  `json:"user_id" binding:"required"`
	Number_of_tickets int     `json:"no_of_tickets" binding:"required"`
	Total_amount      float64 `json:"total_amount" binding:"required"`
}

type BookingRes struct {
	ID                string    `json:"id"`
	Event_id          string    `json:"event_id"`
	EventName         string    `json:"event_name"`
	User_id           string    `json:"user_id"`
	Number_of_tickets int       `json:"no_of_tickets"`
	Total_amount      float64   `json:"total_amount"`
	Created_at        time.Time `json:"created_at"`
}

type UpdateBookingReq struct {
	Number_of_tickets int     `json:"number_of_tickets" binding:"required"`
	Total_amount      float64 `json:"total_amount" binding:"required"`
}
