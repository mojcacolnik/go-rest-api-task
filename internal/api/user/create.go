package user

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/bleenco/go-kit/lib"
	"github.com/bleenco/go-kit/render"

	"github.com/mojcacolnik/go-rest-api-task/internal/db"
	"github.com/mojcacolnik/go-rest-api-task/internal/db/models"
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

	if err := lib.DecodeJSON(r.Body, &f); err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}

	user := models.User{
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		IsActive:  f.IsActive,
	}

	if valid, err := govalidator.ValidateStruct(f); err != nil || !valid {
		render.BadRequestError(w, err.Error())
		return
	}
	db.DB.Create(&user)

	render.JSON(w, http.StatusOK, user)
}
