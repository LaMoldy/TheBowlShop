package models

import "time"

type Rating struct {
	RatingID  uint `gorm:"primaryKey"`
	Rating    uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}
