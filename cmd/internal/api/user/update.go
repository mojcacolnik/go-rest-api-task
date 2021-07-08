package user

import "net/http"

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	type form struct {
		Email     string `json:"email" valid:"email, required"`
		FirstName string `json:"firstname" valid:"stringlength(2|50),optional"`
		LastName  string `json:"lastname" valid:"stringlength(2|50),optional"`
		IsActive  bool   `json:"is_active" valid:"required"`
	}

	params := chi.URLParam(r, "id")
	var f form
	defer r.Body.Close()

	user := User{
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		IsActive:  f.IsActive,
	}

	if err := DB.First(&user, params).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	if err := DB.Save(&user).Error; err != nil {
		RenderJSON(w, http.StatusInternalServerError, errorResponse{Message: err.Error()})
	}
	RenderJSON(w, http.StatusOK, user)
}
