package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" valid:"email, required"`
	FirstName string `json:"firstname" valid:"stringlength(2|50),required"`
	LastName  string `json:"lastname" valid:"stringlength(2|50),required"`
	IsActive  bool   `json:"is_active" valid:"required"`
}
