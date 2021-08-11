package db

import (
	"database/sql"
	"gin-gorm-rails-like-sample-api/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker" // Use <ysql in gorm
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testdb *gorm.DB
var container_name = "TestDBContainer"

func Init() (*gorm.DB, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Name:         container_name,
		Repository:   "mysql",
		Tag:          "5.7",
		Env:          []string{"MYSQL_ROOT_PASSWORD=secret"},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: "0.0.0.0", HostPort: "32804"},
			},
		},
	}

	//resource, err := pool.RunWithOptions(&opts)
	pool.RunWithOptions(&opts)

	if err != nil {
		if err.Error() == ": container already exists" {
			configs, err := config.GetConfigs()
			if err != nil {
				panic(err)
			}
			testdb, _ = gorm.Open(mysql.Open(configs.Database.DataSource), &gorm.Config{})
			return testdb, pool
		} else {
			log.Fatalf("Could not start resource: %s", err)
		}
	}
	if err := pool.Retry(func() error {
		// create database
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		_, err = db.Exec("CREATE DATABASE sample")
		if err != nil {
			panic(err)
		}

		dsn := "root@tcp(127.0.0.1:3306)/sample?parseTime=true"
		testdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}

		sqlDB, err := testdb.DB()
		if err != nil {
			log.Println("Ping for Database is failed.")
			return err
		}

		return sqlDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return testdb, pool
}
