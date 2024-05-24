package models

import "time"

type Payment struct {
	ID           string `gorm:"type:uuid;primaryKey"`
	Booking_id   string
	User_id      string
	Event_id     string
	Amount       float64
	Card_number  string
	Expiry_month string
	Expiry_year  string
	Cvv          string
	Card_holder  string
	Payment_date time.Time
}
