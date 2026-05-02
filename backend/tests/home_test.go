package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/ai"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestHomeDashboardEndpoint(t *testing.T) {
	userRepo := postgres.NewUserRepository()
	dietRepo := postgres.NewDietRepository()
	activityRepo := postgres.NewActivityRepository()
	sleepRepo := postgres.NewSleepRepository()
	meditationRepo := postgres.NewMeditationRepository()
	scoreRepo := postgres.NewScoringRepository()

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
