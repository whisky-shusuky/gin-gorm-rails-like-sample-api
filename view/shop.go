package view

import (
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

type responseShops struct {
	Shops []shop `json:"shops"`
}

// RenderShops render shops.
func RenderShops(c *gin.Context, shops []*entity.Shop) {
	c.JSON(200, responseShops{Shops: convertToViewShops(shops)})
}

// RenderShop render shop.
func RenderShop(c *gin.Context, shop *entity.Shop) {
	c.JSON(200, convertToViewShop(shop))
}
