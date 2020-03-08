package controller

import (
	"fmt"
	"gin-gorm-rails-like-sample-api/model"

	"github.com/gin-gonic/gin"
)

// ShopController is shop controlller
type ShopController struct{}

// IndexShop action: GET /shops
func (pc ShopController) IndexShop(c *gin.Context) {
	shops, err := model.GetShopAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, shops)
	}
}

// CreateShop action: POST /shops
func (pc ShopController) CreateShop(c *gin.Context) {
	_, err := model.CreateShop(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(204, nil)
	}
}

// ShowShop action: GET /shops/:id
func (pc ShopController) ShowShop(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := model.GetShopByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// UpdateShop action: PUT /shops/:id
func (pc ShopController) UpdateShop(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := model.UpdateShopByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// DeleteShop action: DELETE /shops/:id
func (pc ShopController) DeleteShop(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := model.DeleteShopByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
