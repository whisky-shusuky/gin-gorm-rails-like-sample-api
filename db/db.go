package db

import (
	"gin-gorm-rails-like-sample-api/config"
	"gin-gorm-rails-like-sample-api/model/entity"

	"gorm.io/driver/mysql" // Use <ysql in gorm
	"gorm.io/gorm"
)

var (
	// Db is database.
	Db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	configs, err := config.GetConfigs()
	if err != nil {
		panic(err)
	}

	Db, err = gorm.Open(mysql.Open(configs.Database.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return Db
}

// Close is closing db
func Close() {
	db, _ := Db.DB()
	defer db.Close()
}

func autoMigration() {
	Db.AutoMigrate(&entity.Shop{})
	Db.AutoMigrate(&entity.Book{})
}
