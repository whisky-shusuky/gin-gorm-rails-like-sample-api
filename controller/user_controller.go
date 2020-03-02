package controller

import (
	"fmt"
	"gin-gorm-rails-like-sample-api/model"

	"github.com/gin-gonic/gin"
)

// UserController is user controlller
type UserController struct{}

// IndexUser action: GET /users
func (pc UserController) IndexUser(c *gin.Context) {
	users, err := model.GetUserAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

// CreateUser action: POST /users
func (pc UserController) CreateUser(c *gin.Context) {
	_, err := model.CreateUser(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(204, nil)
	}
}

// ShowUser action: GET /users/:id
func (pc UserController) ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := model.GetUserByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// UpdateUser action: PUT /users/:id
func (pc UserController) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := model.UpdateUserByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// DeleteUser action: DELETE /users/:id
func (pc UserController) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := model.DeleteUserByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
