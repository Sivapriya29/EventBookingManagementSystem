package dtos

import "time"

type EventReq struct {
	ID                string
	Event_name        string    `json:"event_name" binding:"required"`
	Event_description string    `json:"event_description" binding:"required"`
	Event_date        time.Time `json:"event_date" binding:"required"`
	Event_time        time.Time `json:"event_time" binding:"required"`
	Event_type        string    `json:"event_type" binding:"required"`
	Location          string    `json:"location" binding:"required"`
	Speaker_name      string    `json:"speaker_name" binding:"required"`
	Organizer_name    string    `json:"organizer_name" binding:"required"`
	Capacity          int       `json:"capacity" binding:"required"`
	Per_person_price  float64   `json:"per_person_price" binding:"required"`
}

type EventRes struct {
	ID                string    `json:"id"`
	Event_name        string    `json:"event_name"`
	Event_description string    `json:"event_description"`
	Event_date        time.Time `json:"event_date"`
	Event_time        time.Time `json:"event_time"`
	Event_type        string    `json:"event_type"`
	Location          string    `json:"location"`
	Speaker_name      string    `json:"speaker_name"`
	Organizer_name    string    `json:"organizer_name"`
	Capacity          int       `json:"capacity"`
	Per_person_price  float64   `json:"per_person_price"`
}
