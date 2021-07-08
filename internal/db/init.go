package db

import (
	"fmt"

	"github.com/mojcacolnik/go-rest-api-task/internal/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type ErrorResponse struct {
	Message string `json:"message"`
}

func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connetct to DB")
	}
	DB.AutoMigrate(&models.User{})
}
