package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
)

type DietRepository struct {
	mu      sync.RWMutex
	records []diet.Record
}

func NewDietRepository() *DietRepository {
	return &DietRepository{
		records: make([]diet.Record, 0),
	}
}

func (r *DietRepository) SaveRecord(_ context.Context, record diet.Record) diet.Record {
	r.mu.Lock()
	defer r.mu.Unlock()

	record.ID = fmt.Sprintf("diet-%d", len(r.records)+1)
	r.records = append(r.records, record)
	return record
}

func (r *DietRepository) ListRecords(_ context.Context) []diet.Record {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]diet.Record, len(r.records))
	copy(out, r.records)
	return out
}

