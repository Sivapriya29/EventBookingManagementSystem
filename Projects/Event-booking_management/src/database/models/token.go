package models

import "time"

type AccessToken struct {
	Token        string `gorm:"type:uuid;primaryKey"`
	RefreshToken string
	UserId       string
	ExpiresAt    time.Time
}

type RefreshToken struct {
	Token     string `gorm:"type:uuid;primaryKey"`
	UserId    string
	ExpiresAt time.Time
}
