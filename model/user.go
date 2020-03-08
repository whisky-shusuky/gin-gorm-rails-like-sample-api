package model

import (
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

// GetUserAll is get all User
func GetUserAll() ([]entity.User, error) {
	db := db.GetDB()
	var u []entity.User
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserByID is get a User
func GetUserByID(id string) (entity.User, error) {
	db := db.GetDB()
	var u entity.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// CreateUser is create User model
func CreateUser(c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUserByID is update a User
func UpdateUserByID(id string, c *gin.Context) (entity.User, error) {
	db := db.GetDB()
	var u entity.User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

// DeleteUserByID is delete a User
func DeleteUserByID(id string) error {
	db := db.GetDB()
	var u entity.User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
