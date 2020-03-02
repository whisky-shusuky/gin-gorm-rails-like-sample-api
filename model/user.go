package model

import (
	"gin-gorm-rails-like-sample-api/db"

	"github.com/gin-gonic/gin"
)

// User is user models property
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// GetUserAll is get all User
func GetUserAll() ([]User, error) {
	db := db.GetDB()
	var u []User
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserByID is get a User
func GetUserByID(id string) (User, error) {
	db := db.GetDB()
	var u User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// CreateUser is create User model
func CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUserByID is update a User
func UpdateUserByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

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
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
