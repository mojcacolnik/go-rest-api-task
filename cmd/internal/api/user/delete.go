package user

import "net/http"

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	params := chi.URLParam(r, "id")

	if err := DB.Delete(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}

	RenderJSON(w, http.StatusOK, "User is successfully deleted!")
}
