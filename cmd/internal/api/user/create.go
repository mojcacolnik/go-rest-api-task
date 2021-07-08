package user

import (
	"net/http"

	"github.com/asaskevich/govalidator"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email     string `json:"email" valid:"email, required"`
		FirstName string `json:"firstname" valid:"stringlength(2|50),required"`
		LastName  string `json:"lastname" valid:"stringlength(2|50),required"`
		IsActive  bool   `json:"is_active" valid:"required"`
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

	if valid, err := govalidator.ValidateStruct(f); err != nil || !valid {
		BadRequestError(w, err.Error())
		return
	}
	DB.Create(&user)

	RenderJSON(w, http.StatusOK, user)
}
