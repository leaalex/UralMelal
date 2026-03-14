package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
)

const HoneypotField = "website_url"

type LeadHandler struct {
	repo *repository.LeadRepository
}

func NewLeadHandler(repo *repository.LeadRepository) *LeadHandler {
	return &LeadHandler{repo: repo}
}

type CreateLeadRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	// Honeypot - must be empty. Bots often fill all fields.
	WebsiteURL string `json:"website_url"`
	Consent   bool   `json:"consent"`
}

func (h *LeadHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateLeadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if req.WebsiteURL != "" {
		http.Error(w, `{"error":"invalid"}`, http.StatusBadRequest)
		return
	}
	if !req.Consent {
		http.Error(w, `{"error":"consent required"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Name) == "" {
		http.Error(w, `{"error":"name required"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Phone) == "" {
		http.Error(w, `{"error":"phone required"}`, http.StatusBadRequest)
		return
	}
	lead := &models.Lead{
		Name:    strings.TrimSpace(req.Name),
		Phone:   strings.TrimSpace(req.Phone),
		Subject: strings.TrimSpace(req.Subject),
		Status:  models.LeadStatusNew,
	}
	if err := h.repo.Create(lead); err != nil {
		http.Error(w, `{"error":"failed to save"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": lead.ID, "message": "ok"})
}

func (h *LeadHandler) List(w http.ResponseWriter, r *http.Request) {
	params := repository.LeadListParams{Page: 1, PerPage: 20}
	if s := r.URL.Query().Get("status"); s != "" {
		st := models.LeadStatus(s)
		if st == models.LeadStatusNew || st == models.LeadStatusProcessed {
			params.Status = &st
		}
	}
	if from := r.URL.Query().Get("date_from"); from != "" {
		if t, err := time.Parse("2006-01-02", from); err == nil {
			params.DateFrom = &t
		}
	}
	if to := r.URL.Query().Get("date_to"); to != "" {
		if t, err := time.Parse("2006-01-02", to); err == nil {
			params.DateTo = &t
		}
	}
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			params.Page = v
		}
	}
	if pp := r.URL.Query().Get("per_page"); pp != "" {
		if v, err := strconv.Atoi(pp); err == nil && v > 0 {
			params.PerPage = v
		}
	}
	leads, total, err := h.repo.List(params)
	if err != nil {
		http.Error(w, `{"error":"failed to list"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"leads": leads, "total": total})
}

func (h *LeadHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	status := models.LeadStatus(input.Status)
	if status != models.LeadStatusNew && status != models.LeadStatusProcessed {
		http.Error(w, `{"error":"invalid status"}`, http.StatusBadRequest)
		return
	}
	if err := h.repo.UpdateStatus(uint(id), status); err != nil {
		http.Error(w, `{"error":"failed to update"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": string(status)})
}
