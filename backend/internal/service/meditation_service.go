package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/meditation"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
)

type MeditationService struct {
	repo *postgres.MeditationRepository
}

func NewMeditationService(repo *postgres.MeditationRepository) *MeditationService {
	return &MeditationService{repo: repo}
}

func (s *MeditationService) RecordSession(ctx context.Context, req meditation.SessionRequest) meditation.Session {
	return s.repo.SaveSession(ctx, meditation.Session{
		ModeKey:   req.ModeKey,
		DurationS: req.DurationS,
		AudioKey:  req.AudioKey,
	})
}

func (s *MeditationService) GetTodaySummary(ctx context.Context) meditation.TodaySummary {
	sessions := s.repo.ListSessions(ctx)
	total := 0
	for _, session := range sessions {
		total += session.DurationS
	}

	return meditation.TodaySummary{
		TotalDurationS: total,
		CompletedCount: len(sessions),
		RecentSessions: sessions,
		Modes:          defaultModes(),
	}
}

func defaultModes() []meditation.BreathMode {
	return []meditation.BreathMode{
		{Key: "calm", Title: "平静呼吸", InhaleSec: 4, HoldSec: 2, ExhaleSec: 6, Description: "适合放慢节奏，降低紧绷感。"},
		{Key: "focus", Title: "专注呼吸", InhaleSec: 4, HoldSec: 4, ExhaleSec: 4, Description: "适合进入工作或学习状态前。"},
		{Key: "sleep", Title: "睡前呼吸", InhaleSec: 4, HoldSec: 7, ExhaleSec: 8, Description: "适合睡前稳定呼吸节律。"},
	}
}
