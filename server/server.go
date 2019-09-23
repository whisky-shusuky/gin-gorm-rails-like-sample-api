package server

import (
	"gin-gorm-viron/controller"

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
		users.GET("", ctrl.Index)
		users.GET("/:id", ctrl.Show)
		users.POST("/", ctrl.Create)
		users.PUT("/:id", ctrl.Update)
		users.DELETE("/:id", ctrl.Delete)
	}

	swagger := r.Group("/swagger")
	{
		ctrl := controller.UserController{}
		swagger.GET("", ctrl.Index)
		swagger.GET("/:id", ctrl.Show)
		swagger.POST("/", ctrl.Create)
		swagger.PUT("/:id", ctrl.Update)
		swagger.DELETE("/:id", ctrl.Delete)
	}
	return r
}
