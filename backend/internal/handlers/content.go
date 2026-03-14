package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/microcosm-cc/bluemonday"

	"github.com/go-chi/chi/v5"

	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
)

var htmlPolicy = bluemonday.StrictPolicy()

type ContentHandler struct {
	repo *repository.ContentRepository
}

func NewContentHandler(repo *repository.ContentRepository) *ContentHandler {
	return &ContentHandler{repo: repo}
}

func (h *ContentHandler) GetByPage(w http.ResponseWriter, r *http.Request) {
	page := models.ContentPage(chi.URLParam(r, "page"))
	if page != models.PageHome && page != models.PageContacts {
		http.Error(w, `{"error":"invalid page"}`, http.StatusBadRequest)
		return
	}
	blocks, err := h.repo.GetByPage(page)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func (h *ContentHandler) GetByPageAdmin(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page != string(models.PageHome) && page != string(models.PageContacts) {
		http.Error(w, `{"error":"invalid page"}`, http.StatusBadRequest)
		return
	}
	blocks, err := h.repo.GetByPageAll(models.ContentPage(page))
	if err != nil {
		http.Error(w, `{"error":"failed to fetch"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func (h *ContentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	block, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	var input struct {
		Title     string `json:"title"`
		Body      string `json:"body"`
		SortOrder int    `json:"sort_order"`
		Hidden    *bool  `json:"hidden"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if input.Title != "" {
		block.Title = input.Title
	}
	if input.Body != "" {
		block.Body = htmlPolicy.Sanitize(input.Body) // strips XSS from HTML/Markdown
	}
	block.SortOrder = input.SortOrder
	if input.Hidden != nil {
		block.Hidden = *input.Hidden
	}
	if err := h.repo.Update(block); err != nil {
		http.Error(w, `{"error":"failed to update"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(block)
}
