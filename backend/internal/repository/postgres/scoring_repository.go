package postgres

import (
	"context"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/scoring"
)

type ScoringRepository struct {
	mu    sync.RWMutex
	score scoring.DailyScore
}

func NewScoringRepository() *ScoringRepository {
	return &ScoringRepository{}
}

func (r *ScoringRepository) Save(_ context.Context, score scoring.DailyScore) scoring.DailyScore {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.score = score
	return r.score
}

func (r *ScoringRepository) Get(_ context.Context) scoring.DailyScore {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.score
}
