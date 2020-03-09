package controller

import (
	"fmt"
	"gin-gorm-rails-like-sample-api/model"
	"gin-gorm-rails-like-sample-api/view"

	"github.com/gin-gonic/gin"
)

// BookController is book controlller
type BookController struct{}

// IndexBook action: GET /books
func (pc BookController) IndexBook(c *gin.Context) {
	books, err := model.GetBookAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		view.RenderBooks(c, books)
	}
}

// CreateBook action: POST /books
func (pc BookController) CreateBook(c *gin.Context) {
	_, err := model.CreateBook(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(204, nil)
	}
}

// ShowBook action: GET /books/:id
func (pc BookController) ShowBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, err := model.GetBookByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		view.RenderBook(c, book)
	}
}

// UpdateBook action: PUT /books/:id
func (pc BookController) UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, err := model.UpdateBookByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		view.RenderBook(c, book)
	}
}

// DeleteBook action: DELETE /books/:id
func (pc BookController) DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := model.DeleteBookByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, nil)
	}
}
