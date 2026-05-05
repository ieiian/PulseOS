package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/service"
)

type DietHandler struct {
	service *service.DietService
}

func NewDietHandler(service *service.DietService) *DietHandler {
	return &DietHandler{service: service}
}

func (h *DietHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/diet/plan/today", h.handleTodayPlan)
	mux.HandleFunc("/api/v1/diet/photo-upload", h.handlePhotoUpload)
	mux.HandleFunc("/api/v1/diet/analyze", h.handleAnalyze)
	mux.HandleFunc("/api/v1/diet/records", h.handleRecords)
}

func (h *DietHandler) handleTodayPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, h.service.GetTodayPlan(r.Context()))
}

func (h *DietHandler) handlePhotoUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	filename := r.URL.Query().Get("filename")
	writeJSON(w, http.StatusCreated, h.service.UploadPhoto(r.Context(), filename))
}

func (h *DietHandler) handleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var req diet.AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	writeJSON(w, http.StatusOK, h.service.Analyze(r.Context(), req))
}

func (h *DietHandler) handleRecords(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.ListRecords(r.Context()))
	case http.MethodPost:
		var req diet.AnalyzeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		writeJSON(w, http.StatusCreated, h.service.QuickRecord(r.Context(), req))
	default:
		methodNotAllowed(w)
	}
}

