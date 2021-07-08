package user

import (
	"net/http"

	"github.com/bleenco/go-kit/render"
	"github.com/go-chi/chi"
	"github.com/mojcacolnik/go-rest-api-task/internal/db"
	"github.com/mojcacolnik/go-rest-api-task/internal/db/models"
)

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := chi.URLParam(r, "id")

	if err := db.DB.Delete(&user, params).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}

	render.JSON(w, http.StatusOK, "User is successfully deleted!")
}
