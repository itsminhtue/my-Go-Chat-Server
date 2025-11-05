package user

import (
	"encoding/json"
	auth "goChat/internal/common/hash"
	"net/http"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input User
	json.NewDecoder(r.Body).Decode(&input)

	hashed, _ := auth.HashPassword(input.Password)
	input.Password = hashed

	h.Repo.Create(r.Context(), &input)
	json.NewEncoder(w).Encode(map[string]string{"message": "Register successfully!"})
}
