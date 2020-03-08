package entity

import "time"

// Shop is shop models property
type Shop struct {
	ID              uint64 `gorm:"primary_key"`
	ShopName        string
	ShopDescription string
	CreatedAt       time.Time
	DeletedAt       *time.Time
}
