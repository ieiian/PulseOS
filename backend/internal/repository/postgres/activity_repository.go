package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
)

type ActivityRepository struct {
	mu      sync.RWMutex
	records []activity.Record
}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{
		records: []activity.Record{
			{ID: "activity-1", Source: "device", ActivityType: "walking", Steps: 4200, Minutes: 35, Intensity: activity.IntensityModerate, CardioPoints: 35},
			{ID: "activity-2", Source: "manual", ActivityType: "cycling", Steps: 0, Minutes: 20, Intensity: activity.IntensityVigorous, CardioPoints: 40},
		},
	}
}

func (r *ActivityRepository) SaveRecord(_ context.Context, record activity.Record) activity.Record {
	r.mu.Lock()
	defer r.mu.Unlock()

	record.ID = fmt.Sprintf("activity-%d", len(r.records)+1)
	r.records = append(r.records, record)
	return record
}

func (r *ActivityRepository) ListRecords(_ context.Context) []activity.Record {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]activity.Record, len(r.records))
	copy(out, r.records)
	return out
}
