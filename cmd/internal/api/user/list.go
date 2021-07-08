package user

import (
	"net/http"

	"github.com/bleenco/go-kit/render"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/db"
)

func HandleListUsers(w http.ResponseWriter, r *http.Request) {
	var users []db.User

	if err := db.DB.Find(&users).Error; err != nil {
		render.JSON(w, http.StatusInternalServerError, db.ErrorResponse{Message: err.Error()})
		return
	}

	render.JSON(w, http.StatusOK, users)
}
