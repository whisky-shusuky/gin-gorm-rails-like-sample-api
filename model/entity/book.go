package entity

import (
	"time"
)

// Book is book models property
type Book struct {
	ID              uint64 `gorm:"primary_key"`
	BookName        string
	BookDescription string
	Sales           uint
	CreatedAt       time.Time
	DeletedAt       *time.Time
}

// Rank calculate rank of Book by sales.
func (book *Book) Rank() *string {

	if book.Sales > 10000 {
		rank := "ベストセラー"
		return &rank
	} else if book.Sales > 1000 {
		rank := "売れ筋"
		return &rank
	} else if book.Sales > 100 {
		rank := "そこそこ"
		return &rank
	} else if book.Sales > 10 {
		rank := "ちょいちょい"
		return &rank
	}
	return nil
}
