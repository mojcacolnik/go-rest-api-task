package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

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
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	RenderJSON(w, http.StatusOK, users)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	var user User

	if err := DB.First(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	RenderJSON(w, http.StatusOK, user)

}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		IsActive  bool   `json:"is_active"`
	}

	var f form
	defer r.Body.Close()

	if err := DecodeJSON(r.Body, &f); err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	user := User{
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		IsActive:  f.IsActive,
	}

	if err := DB.Create(&user).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	RenderJSON(w, http.StatusOK, user)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email     string
		FirstName string
		LastName  string
		IsActive  bool
	}

	params := chi.URLParam(r, "id")
	var f form
	defer r.Body.Close()

	user := User{
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		IsActive:  f.IsActive,
	}

	if err := DB.First(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	if err := DB.Save(&user).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
	}
	RenderJSON(w, http.StatusOK, user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	params := chi.URLParam(r, "id")

	if err := DB.Delete(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	RenderJSON(w, http.StatusOK, "User is successfully deleted!")
}
