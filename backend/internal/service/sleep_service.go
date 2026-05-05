package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/sleep"
	"github.com/tse/PulseOS/backend/internal/ruleengine"
)

type SleepService struct {
	repo SleepRepo
}

func NewSleepService(repo SleepRepo) *SleepService {
	return &SleepService{repo: repo}
}

func (s *SleepService) StartSession(ctx context.Context, req sleep.StartRequest) sleep.Session {
	return s.repo.StartSession(ctx, req.AudioURL)
}

func (s *SleepService) EndSession(ctx context.Context, _ sleep.EndRequest) sleep.TodaySummary {
	durationM := 430
	events := ruleengine.DetectSleepEvents()
	score := ruleengine.CalculateSleepScore(durationM, events)
	advice := ruleengine.BuildSleepAdvice(durationM, events)
	session := s.repo.EndSession(ctx, score, advice, durationM, events)
	return sleep.TodaySummary{
		Session: session,
		Events:  events,
	}
}

func (s *SleepService) GetToday(ctx context.Context) sleep.TodaySummary {
	session, events := s.repo.GetToday(ctx)
	return sleep.TodaySummary{
		Session: session,
		Events:  events,
	}
}
