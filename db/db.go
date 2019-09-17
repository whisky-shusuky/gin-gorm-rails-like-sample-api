package db

import (
	"gin-gorm-viron/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use PostgreSQL in gorm
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// TODO: dockerに繋ぎこむのに失敗したためローカルのmysqlに繋ぎ混んでいる。
	DBMS := "mysql"
	CONNECT := "root:@/gin_gorm_viron"
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
