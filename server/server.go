package server

import (
	"gin-gorm-rails-like-sample-api/controller"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		ctrl := controller.UserController{}
		users.GET("", ctrl.IndexUser)
		users.GET("/:id", ctrl.ShowUser)
		users.POST("/", ctrl.CreateUser)
		users.PUT("/:id", ctrl.UpdateUser)
		users.DELETE("/:id", ctrl.DeleteUser)
	}

	return r
}
