package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/domain/sleep"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestSleepFlowEndpoints(t *testing.T) {
	repo := postgres.NewSleepRepository()
	svc := service.NewSleepService(repo)
	h := handler.NewSleepHandler(svc)
	mux := http.NewServeMux()
	h.Register(mux)

	startBody, _ := json.Marshal(sleep.StartRequest{AudioURL: "/mock/audio/night.wav"})
	startReq := httptest.NewRequest(http.MethodPost, "/api/v1/sleep/sessions/start", bytes.NewReader(startBody))
	startRec := httptest.NewRecorder()
	mux.ServeHTTP(startRec, startReq)
	if startRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", startRec.Code)
	}

	endBody, _ := json.Marshal(sleep.EndRequest{SessionID: "sleep-1"})
	endReq := httptest.NewRequest(http.MethodPost, "/api/v1/sleep/sessions/end", bytes.NewReader(endBody))
	endRec := httptest.NewRecorder()
	mux.ServeHTTP(endRec, endReq)
	if endRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", endRec.Code)
	}
}
