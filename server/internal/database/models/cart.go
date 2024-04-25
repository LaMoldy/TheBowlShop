package models

type Cart struct {
	CartID     uint `gorm:"primaryKey"`
	UserID     uint
	ProductIDs []uint
}
