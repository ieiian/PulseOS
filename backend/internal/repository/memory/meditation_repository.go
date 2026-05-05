package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/meditation"
)

type MeditationRepository struct {
	mu       sync.RWMutex
	sessions []meditation.Session
}

func NewMeditationRepository() *MeditationRepository {
	return &MeditationRepository{
		sessions: []meditation.Session{
			{
				ID:        "meditation-1",
				ModeKey:   "box_breathing",
				DurationS: 300,
				AudioKey:  "calm_waves",
			},
			{
				ID:        "meditation-2",
				ModeKey:   "deep_breathing",
				DurationS: 600,
				AudioKey:  "forest_sounds",
			},
		},
	}
}

func (r *MeditationRepository) SaveSession(_ context.Context, session meditation.Session) meditation.Session {
	r.mu.Lock()
	defer r.mu.Unlock()

	session.ID = fmt.Sprintf("meditation-%d", len(r.sessions)+1)
	r.sessions = append(r.sessions, session)
	return session
}

func (r *MeditationRepository) ListSessions(_ context.Context) []meditation.Session {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]meditation.Session, len(r.sessions))
	copy(out, r.sessions)
	return out
}
