package models

import "time"

type User struct {
	UserID       uint `gorm:"primaryKey"`
	Email        string
	Password     string
	FirstName    string
	LastName     string
	ProfileImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	IsAdmin      bool
	IsDeleted    bool
}
