package db

import (
	"gin-gorm-rails-like-sample-api/config"
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use <ysql in gorm
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

	Db, err = gorm.Open(configs.Database.Dialect, configs.Database.DataSource)
	Db.LogMode(true)
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
	if err := Db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	Db.AutoMigrate(&entity.Shop{})
}
