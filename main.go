package main

import (
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
