package auth

import (
	"encoding/json"
	"goChat/internal/common/hash"
	"goChat/internal/user"
	"net/http"
)

type Handler struct {
	UserRepo *user.Repository
}

func NewHandler(repo *user.Repository) *Handler {
	return &Handler{UserRepo: repo}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	u, err := h.UserRepo.FindByEmail(r.Context(), input.Email)

	if err != nil || !hash.CheckPassword(u.Password, input.Password) {
		http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
	}

	token, err := GenerateToken(u.ID.Hex())

	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
