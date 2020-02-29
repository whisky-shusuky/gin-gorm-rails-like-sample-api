package controller

import (
	"fmt"
	user "gin-gorm-rails-like-sample-api/service"

	"github.com/gin-gonic/gin"
)

// SwaggerController is swagger controlller
type SwaggerController struct{}

// ShowSwagger action: GET /users/:id
func (pc SwaggerController) ShowSwagger(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
