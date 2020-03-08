package model

import (
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

// GetShopAll is get all Shop
func GetShopAll() ([]entity.Shop, error) {
	db := db.GetDB()
	var u []entity.Shop
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// GetShopByID is get a Shop
func GetShopByID(id string) (entity.Shop, error) {
	db := db.GetDB()
	var u entity.Shop
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// CreateShop is create Shop model
func CreateShop(c *gin.Context) (entity.Shop, error) {
	db := db.GetDB()
	var u entity.Shop

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateShopByID is update a Shop
func UpdateShopByID(id string, c *gin.Context) (entity.Shop, error) {
	db := db.GetDB()
	var u entity.Shop

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteShopByID is delete a Shop
func DeleteShopByID(id string) error {
	db := db.GetDB()
	var u entity.Shop

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
