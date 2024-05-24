package models

import (
	"time"
)

type User struct {
	ID         string `gorm:"type:uuid;primaryKey"`
	First_name string
	Last_name  string
	Email      string
	Password   string
	Mobile     string
	Role       string
	Created_at time.Time
}
