package models

type Feedback struct {
	ID       string `gorm:"type:uuid;primaryKey"`
	User_id  string
	Event_id string
	Rating   float64
	Comments string
}
