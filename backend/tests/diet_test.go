package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tse/PulseOS/backend/internal/ai"
	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/domain/user"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/repository/memory"
	"github.com/tse/PulseOS/backend/internal/service"
)

func TestDietAnalyzeEndpoint(t *testing.T) {
	userRepo := memory.NewUserRepository()
	userRepo.SaveProfile(t.Context(), user.Profile{
		Name:        "Test",
		PrimaryGoal: user.GoalFatLoss,
		HealthFlags: []string{"diabetes"},
	})

	dietRepo := memory.NewDietRepository()
	dietService := service.NewDietService(dietRepo, userRepo, ai.NewService())
	dietHandler := handler.NewDietHandler(dietService)

	mux := http.NewServeMux()
	dietHandler.Register(mux)

	body, _ := json.Marshal(diet.AnalyzeRequest{
		ImageURL:    "/mock/uploads/tea.jpg",
		MealType:    "snack",
		ManualItems: []string{"奶茶", "炸鸡"},
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/diet/analyze", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var result diet.AnalyzeResult
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if result.Recommendation != diet.RecommendationForbidden {
		t.Fatalf("expected forbidden, got %s", result.Recommendation)
	}
}
