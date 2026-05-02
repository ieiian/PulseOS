package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/domain/user"
	"github.com/tse/PulseOS/backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/users/onboarding", h.handleOnboarding)
	mux.HandleFunc("/api/v1/users/profile", h.handleProfile)
	mux.HandleFunc("/api/v1/users/settings", h.handleSettings)
	mux.HandleFunc("/api/v1/users/stats", h.handleStats)
}

func (h *UserHandler) handleOnboarding(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var profile user.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	writeJSON(w, http.StatusCreated, h.service.Onboard(r.Context(), profile))
}

func (h *UserHandler) handleProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.GetProfile(r.Context()))
	case http.MethodPut:
		var profile user.Profile
		if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		writeJSON(w, http.StatusOK, h.service.UpdateProfile(r.Context(), profile))
	default:
		methodNotAllowed(w)
	}
}

func (h *UserHandler) handleSettings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.GetSettings(r.Context()))
	case http.MethodPut:
		var settings user.Settings
		if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		writeJSON(w, http.StatusOK, h.service.UpdateSettings(r.Context(), settings))
	default:
		methodNotAllowed(w)
	}
}

func (h *UserHandler) handleStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	writeJSON(w, http.StatusOK, h.service.GetStats(r.Context()))
}

func methodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

