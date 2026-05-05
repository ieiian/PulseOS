package memory

import (
	"context"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/scoring"
)

type ScoringRepository struct {
	mu     sync.RWMutex
	scores []scoring.DailyScore
}

func NewScoringRepository() *ScoringRepository {
	return &ScoringRepository{
		scores: []scoring.DailyScore{
			{Date: "2026-05-03", DietScore: 75, ActivityScore: 80, SleepScore: 70, TotalScore: 75},
			{Date: "2026-05-04", DietScore: 80, ActivityScore: 85, SleepScore: 75, TotalScore: 80},
			{Date: "2026-05-05", DietScore: 85, ActivityScore: 90, SleepScore: 80, TotalScore: 85},
		},
	}
}

func (r *ScoringRepository) Save(_ context.Context, score scoring.DailyScore) scoring.DailyScore {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, s := range r.scores {
		if s.Date == score.Date {
			r.scores[i] = score
			return score
		}
	}
	r.scores = append(r.scores, score)
	return score
}

func (r *ScoringRepository) Get(_ context.Context) scoring.DailyScore {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.scores) == 0 {
		return scoring.DailyScore{}
	}
	return r.scores[len(r.scores)-1]
}

func (r *ScoringRepository) GetHistory(_ context.Context) []scoring.DailyScore {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]scoring.DailyScore, len(r.scores))
	copy(out, r.scores)
	return out
}
