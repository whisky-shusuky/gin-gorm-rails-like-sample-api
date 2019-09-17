package main

import (
	"gin-gorm-viron/db"
	"gin-gorm-viron/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
