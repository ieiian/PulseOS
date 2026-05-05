package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/domain/meditation"
	"github.com/tse/PulseOS/backend/internal/domain/scoring"
	"github.com/tse/PulseOS/backend/internal/domain/sleep"
	"github.com/tse/PulseOS/backend/internal/domain/user"
)

type UserRepo interface {
	SaveProfile(ctx context.Context, profile user.Profile) user.Profile
	GetProfile(ctx context.Context) user.Profile
	SaveSettings(ctx context.Context, settings user.Settings) user.Settings
	GetSettings(ctx context.Context) user.Settings
	GetStats(ctx context.Context) user.Stats
}

type DietRepo interface {
	SaveRecord(ctx context.Context, record diet.Record) diet.Record
	ListRecords(ctx context.Context) []diet.Record
}

type ActivityRepo interface {
	SaveRecord(ctx context.Context, record activity.Record) activity.Record
	ListRecords(ctx context.Context) []activity.Record
	ListDailyPoints(ctx context.Context) []int
}

type MeditationRepo interface {
	SaveSession(ctx context.Context, session meditation.Session) meditation.Session
	ListSessions(ctx context.Context) []meditation.Session
}

type SleepRepo interface {
	StartSession(ctx context.Context, audioURL string) sleep.Session
	EndSession(ctx context.Context, score int, advice string, durationM int, events []sleep.Event) sleep.Session
	GetToday(ctx context.Context) (sleep.Session, []sleep.Event)
}

type ScoringRepo interface {
	Save(ctx context.Context, score scoring.DailyScore) scoring.DailyScore
	Get(ctx context.Context) scoring.DailyScore
	GetHistory(ctx context.Context) []scoring.DailyScore
}
