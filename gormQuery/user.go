package gormQuery

import (
	"gorm/database"
	"gorm/helper"
	"gorm/models/userModel"
	"log"
	"time"
)

func CreateUser() {
	now := time.Now()
	user := userModel.User{
		Id:        helper.GenerateUUID(),
		Firstname: "suttipong",
		Lastname:  "Sak",
		Age:       17,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	if err := userModelHelper.Insert(&user); err != nil {
		log.Println("Error inserting user: ", err)
	}
	log.Println("Created user: ", user)
}

func GetUserById() {
	id := "58db5c65-d01d-471d-b7dc-a2c84de3fd79"
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.FindById(id)
	if err != nil {
		log.Println("Error inserting user: ", err)
	}

	log.Println("user:", user)
}
