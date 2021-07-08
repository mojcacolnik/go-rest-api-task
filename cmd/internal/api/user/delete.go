package user

import (
	"net/http"

	"github.com/bleenco/go-kit/render"
	"github.com/go-chi/chi"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/db"
)

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	params := chi.URLParam(r, "id")

	if err := db.DB.Delete(&user, params).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}

	render.JSON(w, http.StatusOK, "User is successfully deleted!")
}
