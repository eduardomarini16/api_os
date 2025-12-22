package auth

import (
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo *Repository
}

func NesHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "erro ao gerar senha", http.StatusInternalServerError)
		return
	}

	err = h.repo.CreateUser(
		context.Background(),
		req.Name,
		req.Email,
		hash,
		req.Role,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
