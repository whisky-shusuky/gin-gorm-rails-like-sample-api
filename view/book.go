package view

import (
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

type responseBooks struct {
	Books []book `json:"books"`
}

// RenderBooks render books.
func RenderBooks(c *gin.Context, books []*entity.Book) {
	c.JSON(200, responseBooks{Books: convertToViewBooks(books)})
}

// RenderBook render book.
func RenderBook(c *gin.Context, book *entity.Book) {
	c.JSON(200, convertToViewBook(book))
}
