package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/memory"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestActivityRecordEndpoint(t *testing.T) {
	repo := memory.NewActivityRepository()
	svc := service.NewActivityService(repo)
	h := handler.NewActivityHandler(svc)
	mux := http.NewServeMux()
	h.Register(mux)

	body, _ := json.Marshal(activity.ManualRecordRequest{
		ActivityType: "running",
		Minutes:      30,
		Intensity:    activity.IntensityVigorous,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/activity/records", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	var out activity.Record
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if out.CardioPoints != 60 {
		t.Fatalf("expected 60 cardio points, got %d", out.CardioPoints)
	}
}
