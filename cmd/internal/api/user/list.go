package user

import "net/http"

func HandleListUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	RenderJSON(w, http.StatusOK, users)
}
