package models

type User struct {
	UserID       uint `gorm:"primaryKey"`
	Email        string
	Password     string
	FirstName    string
	LastName     string
	ProfileImage string
	IsAdmin      bool
	IsDeleted    bool
}
