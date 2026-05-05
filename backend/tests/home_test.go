package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/ai"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/memory"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestHomeDashboardEndpoint(t *testing.T) {
	userRepo := memory.NewUserRepository()
	dietRepo := memory.NewDietRepository()
	activityRepo := memory.NewActivityRepository()
	sleepRepo := memory.NewSleepRepository()
	meditationRepo := memory.NewMeditationRepository()
	scoreRepo := memory.NewScoringRepository()

	dietService := service.NewDietService(dietRepo, userRepo, ai.NewService())
	activityService := service.NewActivityService(activityRepo)
	sleepService := service.NewSleepService(sleepRepo)
	meditationService := service.NewMeditationService(meditationRepo)
	scoringService := service.NewScoringService(scoreRepo, dietService, activityService, sleepService, meditationService)
	homeHandler := handler.NewHomeHandler(scoringService)

	mux := http.NewServeMux()
	homeHandler.Register(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/home/dashboard", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
