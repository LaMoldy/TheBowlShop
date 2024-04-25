package models

import "time"

type Comment struct {
	CommentID uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}
