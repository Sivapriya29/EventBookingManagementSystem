package models

import (
	"time"
)

type Event struct {
	ID                string `gorm:"type:uuid;primaryKey"`
	Event_name        string
	Event_description string
	Event_date        time.Time
	Event_time        time.Time
	Event_type        string
	Location          string
	Speaker_name      string
	Organizer_name    string
	Capacity          int
	Per_person_price  float64
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
}
