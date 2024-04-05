package main

import (
	"gorm/configs"
	"gorm/database"
	"gorm/gormQuery"
)

func main() {
	configs.Init()
	database.Init()

	gormQuery.CreateUser()
}
