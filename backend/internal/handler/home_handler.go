package handler

import (
	"net/http"

	"github.com/tse/PulseOS/backend/internal/service"
)

type HomeHandler struct {
	service *service.ScoringService
}

func NewHomeHandler(service *service.ScoringService) *HomeHandler {
	return &HomeHandler{service: service}
}

func (h *HomeHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/home/dashboard", h.handleDashboard)
}

func (h *HomeHandler) handleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	writeJSON(w, http.StatusOK, h.service.BuildDashboard(r.Context()))
}
