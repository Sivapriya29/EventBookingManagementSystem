package models

import "time"

type Booking struct {
	ID                string `gorm:"type:uuid;primaryKey"`
	Event_id          string
	Event_name        string
	User_id           string
	Number_of_tickets int
	Total_amount      float64
	Created_at        time.Time
}
