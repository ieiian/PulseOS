package postgres

import (
	"context"
	"sync"
	"time"

	"github.com/tse/PulseOS/backend/internal/domain/sleep"
)

type SleepRepository struct {
	mu      sync.RWMutex
	session sleep.Session
	events  []sleep.Event
}

func NewSleepRepository() *SleepRepository {
	return &SleepRepository{
		session: sleep.Session{},
		events:  []sleep.Event{},
	}
}

func (r *SleepRepository) StartSession(_ context.Context, audioURL string) sleep.Session {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.session = sleep.Session{
		ID:        "sleep-1",
		Status:    "recording",
		StartedAt: time.Now().Format(time.RFC3339),
		AudioURL:  audioURL,
	}
	r.events = []sleep.Event{}
	return r.session
}

func (r *SleepRepository) EndSession(_ context.Context, score int, advice string, durationM int, events []sleep.Event) sleep.Session {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.session.Status = "completed"
	r.session.EndedAt = time.Now().Format(time.RFC3339)
	r.session.Score = score
	r.session.Advice = advice
	r.session.DurationM = durationM
	r.events = events
	return r.session
}

func (r *SleepRepository) GetToday(_ context.Context) (sleep.Session, []sleep.Event) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	session := r.session
	events := make([]sleep.Event, len(r.events))
	copy(events, r.events)
	return session, events
}
