package user

import (
	"net/http"

	"github.com/bleenco/go-kit/render"
	"github.com/go-chi/chi"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/db"
)

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email     string `json:"email" valid:"email, required"`
		FirstName string `json:"firstname" valid:"stringlength(2|50),optional"`
		LastName  string `json:"lastname" valid:"stringlength(2|50),optional"`
		IsActive  bool   `json:"is_active" valid:"required"`
	}

	params := chi.URLParam(r, "id")
	var f form
	defer r.Body.Close()

	user := db.User{
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		IsActive:  f.IsActive,
	}

	if err := db.DB.First(&user, params).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}
	if err := db.DB.Save(&user).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
	}
	render.JSON(w, http.StatusOK, user)
}
