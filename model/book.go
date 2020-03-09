package model

import (
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/gin-gonic/gin"
)

// GetBookAll is get all Book
func GetBookAll() ([]*entity.Book, error) {
	db := db.GetDB()
	var u []*entity.Book
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// GetBookByID is get a Book
func GetBookByID(id string) (*entity.Book, error) {
	db := db.GetDB()
	var u entity.Book
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return &u, err
	}

	return &u, nil
}

// CreateBook is create Book model
func CreateBook(c *gin.Context) (*entity.Book, error) {
	db := db.GetDB()
	var u entity.Book

	if err := c.BindJSON(&u); err != nil {
		return nil, err
	}

	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateBookByID is update a Book
func UpdateBookByID(id string, c *gin.Context) (*entity.Book, error) {
	db := db.GetDB()
	var u entity.Book

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	if err := c.BindJSON(&u); err != nil {
		return nil, err
	}

	db.Save(&u)

	return &u, nil
}

// DeleteBookByID is delete a Book
func DeleteBookByID(id string) error {
	db := db.GetDB()
	var u entity.Book

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
