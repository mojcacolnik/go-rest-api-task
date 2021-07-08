package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model
	Email     string `json:"email" valid:"email, required"`
	FirstName string `json:"firstname" valid:"stringlength(2|50),required"`
	LastName  string `json:"lastname" valid:"stringlength(2|50),required"`
	IsActive  bool   `json:"is_active" valid:"required"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connetct to DB")
	}
	DB.AutoMigrate(&User{})
}
