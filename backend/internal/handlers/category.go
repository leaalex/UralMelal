package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
)

type CategoryHandler struct {
	repo *repository.CategoryRepository
}

func NewCategoryHandler(repo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

func slugify(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

func (h *CategoryHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	cats, err := h.repo.GetTree()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch categories"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if cat.Name == "" {
		http.Error(w, `{"error":"name required"}`, http.StatusBadRequest)
		return
	}
	if cat.Slug == "" {
		cat.Slug = slugify(cat.Name)
	}
	if err := h.repo.Create(&cat); err != nil {
		http.Error(w, `{"error":"failed to create"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	cat, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	var input models.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if input.Name != "" {
		cat.Name = input.Name
	}
	if input.Slug != "" {
		cat.Slug = input.Slug
	} else if input.Name != "" {
		cat.Slug = slugify(input.Name)
	}
	cat.ParentID = input.ParentID
	if input.SortOrder != 0 || input.Name != "" {
		cat.SortOrder = input.SortOrder
	}
	if err := h.repo.Update(cat); err != nil {
		http.Error(w, `{"error":"failed to update"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cat)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
