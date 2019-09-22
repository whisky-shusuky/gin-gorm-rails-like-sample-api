package server

import (
	user "gin-gorm-viron/controller"

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
		ctrl := user.Controller{}
		users.GET("", ctrl.Index)
		users.GET("/:id", ctrl.Show)
		users.POST("/", ctrl.Create)
		users.PUT("/:id", ctrl.Update)
		users.DELETE("/:id", ctrl.Delete)
	}
	return r
}
