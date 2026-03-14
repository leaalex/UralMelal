package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
	"uralmelal/backend/internal/services"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.List()
	if err != nil {
		http.Error(w, `{"error":"failed to list"}`, http.StatusInternalServerError)
		return
	}
	for i := range users {
		users[i].PasswordHash = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		http.Error(w, `{"error":"username and password required"}`, http.StatusBadRequest)
		return
	}
	role := models.UserRole(req.Role)
	if role != models.RoleAdmin && role != models.RoleManager && role != models.RoleEditor {
		role = models.RoleEditor
	}
	exists, _ := h.repo.ExistsByUsername(req.Username)
	if exists {
		http.Error(w, `{"error":"username exists"}`, http.StatusConflict)
		return
	}
	hash, err := services.HashPassword(req.Password)
	if err != nil {
		http.Error(w, `{"error":"failed"}`, http.StatusInternalServerError)
		return
	}
	user := &models.User{Username: req.Username, PasswordHash: hash, Role: role}
	if err := h.repo.Create(user); err != nil {
		http.Error(w, `{"error":"failed to create"}`, http.StatusInternalServerError)
		return
	}
	user.PasswordHash = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	user, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Role != "" {
		role := models.UserRole(input.Role)
		if role == models.RoleAdmin || role == models.RoleManager || role == models.RoleEditor {
			user.Role = role
		}
	}
	if input.Password != "" {
		hash, err := services.HashPassword(input.Password)
		if err != nil {
			http.Error(w, `{"error":"failed"}`, http.StatusInternalServerError)
			return
		}
		user.PasswordHash = hash
	}
	if err := h.repo.Update(user); err != nil {
		http.Error(w, `{"error":"failed to update"}`, http.StatusInternalServerError)
		return
	}
	user.PasswordHash = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		http.Error(w, `{"error":"failed to delete"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
