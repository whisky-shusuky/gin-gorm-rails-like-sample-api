package db

import (
	"database/sql"
	"fmt"
	"gin-gorm-rails-like-sample-api/config"
	"gin-gorm-rails-like-sample-api/model/entity"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql" // Use <ysql in gorm
	"gorm.io/gorm"
)

var (
	// Db is database.
	Db  *gorm.DB
	err error
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	DBType   string
}

func buildDBConfig(host, port, user, name, password string, dbType string) *DBConfig {
	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   name,
		Password: password,
		DBType:   dbType,
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	if dbConfig.DBType == "cloudsql" {
		return fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.DBName,
		)
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// Init is initialize db from main function
func Init() {
	configs, err := config.GetConfigs()
	if err != nil {
		panic(err)
	}

	// databaseが無ければ作る
	db, err := sql.Open("mysql", "root@tcp(db)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS sample")
	if err != nil {
		panic(err)
	}

	// gorm設定
	Db, err = gorm.Open(mysql.Open(configs.Database.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// migration
	dbUser := "root"
	dbPassword := ""
	dbPort := "3306"
	dbHost := "db"
	dbName := "sample"
	dbType := "mysql"

	fmt.Println("--- Connecting Migrations ---")
	dbConfig := buildDBConfig(dbHost, dbPort, dbUser, dbName, dbPassword, dbType)
	dbURL := dbURL(dbConfig)

	m, err := migrate.New("file://db/migrations/", dbType+"://"+dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Running Migration ---")
	m.Steps(1000)

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
