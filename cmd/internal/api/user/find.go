package user

import (
	"net/http"

	"github.com/bleenco/go-kit/render"
	"github.com/go-chi/chi"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/db"
)

func HandleFindUser(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	var user db.User

	if err := db.DB.First(&user, params).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}
	render.JSON(w, http.StatusOK, user)

}
