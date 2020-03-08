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

	shops := r.Group("/shops")
	{
		ctrl := controller.ShopController{}
		shops.GET("", ctrl.IndexShop)
		shops.GET("/:id", ctrl.ShowShop)
		shops.POST("", ctrl.CreateShop)
		shops.PUT("/:id", ctrl.UpdateShop)
		shops.DELETE("/:id", ctrl.DeleteShop)
	}

	return r
}
