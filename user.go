package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const (
	dbName = "go-gorm-task"
	dbPass = "12345"
	dbHost = "localhost"
	dbPort = "9000"
)

var DNS = fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=true", dbPass, dbHost, dbPort, dbName)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	IsActive  bool   `json:"is_active"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func initialMigration() {
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connetct to DB")
	}
	DB.AutoMigrate(&User{})
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	err := DB.Find(&users).Error
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(users)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	var user User
	err := DB.First(&user, params).Error
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)

}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	err := DB.Create(&user).Error
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	var user User
	err := DB.First(&user, params).Error
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	DB.First(&user, params)
	json.NewDecoder(r.Body).Decode(&user)
	err = DB.Save(&user).Error
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	var user User
	DB.Delete(&user, params)
	json.NewEncoder(w).Encode("User is successfully deleted!")
}
