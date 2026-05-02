package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
	"github.com/tse/PulseOS/backend/internal/service"
)

type ActivityHandler struct {
	service *service.ActivityService
}

func NewActivityHandler(service *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

func (h *ActivityHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/activity/records", h.handleRecords)
	mux.HandleFunc("/api/v1/activity/today", h.handleToday)
	mux.HandleFunc("/api/v1/activity/week", h.handleWeek)
}

func (h *ActivityHandler) handleRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var req activity.ManualRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	writeJSON(w, http.StatusCreated, h.service.RecordManualActivity(r.Context(), req))
}

func (h *ActivityHandler) handleToday(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, h.service.GetTodaySummary(r.Context()))
}

func (h *ActivityHandler) handleWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, h.service.GetWeekSummary(r.Context()))
}
