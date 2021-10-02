package server

import (
	"gin-gorm-rails-like-sample-api/controller"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := Router()
	r.Run()
}

func Router() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		shopCtrl := controller.ShopController{}
		apiv1.GET("/shops", shopCtrl.IndexShop)
		apiv1.GET("/shops/:id", shopCtrl.ShowShop)
		apiv1.POST("/shops", shopCtrl.CreateShop)
		apiv1.PUT("/shops/:id", shopCtrl.UpdateShop)
		apiv1.DELETE("/shops/:id", shopCtrl.DeleteShop)

		bookCtrl := controller.BookController{}
		apiv1.GET("/books", bookCtrl.IndexBook)
		apiv1.GET("/books/:id", bookCtrl.ShowBook)
		apiv1.POST("/books", bookCtrl.CreateBook)
		apiv1.PUT("/books/:id", bookCtrl.UpdateBook)
		apiv1.DELETE("/books/:id", bookCtrl.DeleteBook)
	}
	return r
}
