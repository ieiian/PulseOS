package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/domain/meditation"
	"github.com/tse/PulseOS/backend/internal/service"
)

type MeditationHandler struct {
	service *service.MeditationService
}

func NewMeditationHandler(service *service.MeditationService) *MeditationHandler {
	return &MeditationHandler{service: service}
}

func (h *MeditationHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/meditation/sessions", h.handleSessions)
	mux.HandleFunc("/api/v1/meditation/today", h.handleToday)
}

func (h *MeditationHandler) handleSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var req meditation.SessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	writeJSON(w, http.StatusCreated, h.service.RecordSession(r.Context(), req))
}

func (h *MeditationHandler) handleToday(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, h.service.GetTodaySummary(r.Context()))
}
