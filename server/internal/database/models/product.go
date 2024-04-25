package models

import "time"

type Product struct {
	ProductID    uint `gorm:"primaryKey"`
	Name         string
	ProductImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	IsDeleted    bool
}
