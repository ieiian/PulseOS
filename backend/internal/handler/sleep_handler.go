package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/domain/sleep"
	"github.com/tse/PulseOS/backend/internal/service"
)

type SleepHandler struct {
	service *service.SleepService
}

func NewSleepHandler(service *service.SleepService) *SleepHandler {
	return &SleepHandler{service: service}
}

func (h *SleepHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/sleep/sessions/start", h.handleStart)
	mux.HandleFunc("/api/v1/sleep/sessions/end", h.handleEnd)
	mux.HandleFunc("/api/v1/sleep/today", h.handleToday)
}

func (h *SleepHandler) handleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	var req sleep.StartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	writeJSON(w, http.StatusCreated, h.service.StartSession(r.Context(), req))
}

func (h *SleepHandler) handleEnd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	var req sleep.EndRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	writeJSON(w, http.StatusOK, h.service.EndSession(r.Context(), req))
}

func (h *SleepHandler) handleToday(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, h.service.GetToday(r.Context()))
}
