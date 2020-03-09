package entity

import "time"

// Book is book models property
type Book struct {
	ID              uint64 `gorm:"primary_key"`
	BookName        string
	BookDescription string
	Sales           uint
	CreatedAt       time.Time
	DeletedAt       *time.Time
}
