package user

import "net/http"

func HandleFindUser(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	var user User

	if err := DB.First(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	RenderJSON(w, http.StatusOK, user)

}
