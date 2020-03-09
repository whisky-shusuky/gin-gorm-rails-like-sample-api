package view

import (
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

type responseBooks struct {
	Books []bookWithRank `json:"books"`
}

// RenderBooks render books.
func RenderBooks(c *gin.Context, books []*entity.Book) {
	c.JSON(200, responseBooks{Books: convertToViewBooks(books)})
}

// RenderBook render book.
func RenderBook(c *gin.Context, book *entity.Book) {
	bookWithRank := bookWithRank{
		ID:              book.ID,
		BookName:        book.BookName,
		BookDescription: book.BookDescription,
		Sales:           book.Sales,
		Rank:            book.Rank(),
	}
	c.JSON(200, bookWithRank)
}
