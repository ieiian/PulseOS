package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/domain/meditation"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestMeditationSessionEndpoint(t *testing.T) {
	repo := postgres.NewMeditationRepository()
	svc := service.NewMeditationService(repo)
	h := handler.NewMeditationHandler(svc)
	mux := http.NewServeMux()
	h.Register(mux)

	body, _ := json.Marshal(meditation.SessionRequest{
		ModeKey:   "sleep",
		DurationS: 600,
		AudioKey:  "night-rain",
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/meditation/sessions", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}
