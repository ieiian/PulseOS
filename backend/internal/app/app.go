package app

import (
	"fmt"
	"net/http"

	"github.com/tse/PulseOS/backend/internal/ai"
	"github.com/tse/PulseOS/backend/internal/bootstrap"
	"github.com/tse/PulseOS/backend/internal/handler"
	"github.com/tse/PulseOS/backend/internal/middleware"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
	"github.com/tse/PulseOS/backend/internal/service"
)

type Server struct {
	config bootstrap.Config
}

func NewServer() (*Server, error) {
	cfg, err := bootstrap.LoadConfig("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	return &Server{config: cfg}, nil
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.config.Server.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	userRepo := postgres.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userHandler.Register(mux)

	dietRepo := postgres.NewDietRepository()
	aiService := ai.NewService()
	dietService := service.NewDietService(dietRepo, userRepo, aiService)
	dietHandler := handler.NewDietHandler(dietService)
	dietHandler.Register(mux)

	activityRepo := postgres.NewActivityRepository()
	activityService := service.NewActivityService(activityRepo)
	activityHandler := handler.NewActivityHandler(activityService)
	activityHandler.Register(mux)

	meditationRepo := postgres.NewMeditationRepository()
	meditationService := service.NewMeditationService(meditationRepo)
	meditationHandler := handler.NewMeditationHandler(meditationService)
	meditationHandler.Register(mux)

	sleepRepo := postgres.NewSleepRepository()
	sleepService := service.NewSleepService(sleepRepo)
	sleepHandler := handler.NewSleepHandler(sleepService)
	sleepHandler.Register(mux)

	scoringRepo := postgres.NewScoringRepository()
	scoringService := service.NewScoringService(scoringRepo, dietService, activityService, sleepService, meditationService)
	homeHandler := handler.NewHomeHandler(scoringService)
	homeHandler.Register(mux)

	return http.ListenAndServe(addr, middleware.WithAccessLog(mux))
}
