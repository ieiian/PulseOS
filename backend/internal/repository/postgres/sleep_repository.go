package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/sleep"
)

type SleepRepository struct {
	pool *pgxpool.Pool
}

func NewSleepRepository(pool *pgxpool.Pool) *SleepRepository {
	return &SleepRepository{pool: pool}
}

func (r *SleepRepository) StartSession(ctx context.Context, audioURL string) sleep.Session {
	now := time.Now()
	var id int64
	var createdAt time.Time
	err := r.pool.QueryRow(ctx,
		`INSERT INTO sleep_records (user_id, started_at, status, audio_url)
		 VALUES (1, $1, 'recording', $2) RETURNING id, created_at`,
		now, audioURL,
	).Scan(&id, &createdAt)
	if err != nil {
		return sleep.Session{}
	}
	return sleep.Session{
		ID:        fmtID(id),
		UserID:    1,
		Status:    "recording",
		StartedAt: now.Format(time.RFC3339),
		AudioURL:  audioURL,
		CreatedAt: createdAt,
	}
}

func (r *SleepRepository) EndSession(ctx context.Context, score int, advice string, durationM int, events []sleep.Event) sleep.Session {
	now := time.Now()

	var id int64
	var startedAt time.Time
	err := r.pool.QueryRow(ctx,
		`UPDATE sleep_records SET status='completed', ended_at=$1, score=$2, advice=$3, duration_m=$4
		 WHERE status='recording' RETURNING id, started_at`,
		now, score, advice, durationM,
	).Scan(&id, &startedAt)
	if err != nil {
		return sleep.Session{}
	}

	for _, e := range events {
		_, _ = r.pool.Exec(ctx,
			`INSERT INTO sleep_events (sleep_record_id, event_type, event_timestamp, level) VALUES ($1, $2, $3, $4)`,
			id, e.Type, e.Timestamp, e.Level)
	}

	return sleep.Session{
		ID:        fmtID(id),
		UserID:    1,
		Status:    "completed",
		StartedAt: startedAt.Format(time.RFC3339),
		EndedAt:   now.Format(time.RFC3339),
		DurationM: durationM,
		Score:     score,
		Advice:    advice,
		CreatedAt: now,
	}
}

func (r *SleepRepository) GetToday(ctx context.Context) (sleep.Session, []sleep.Event) {
	var s sleep.Session
	var id int64
	var startedAt time.Time
	var endedAt *time.Time
	var createdAt time.Time

	err := r.pool.QueryRow(ctx,
		`SELECT id, started_at, ended_at, duration_m, score, audio_url, advice, created_at
		 FROM sleep_records ORDER BY created_at DESC LIMIT 1`,
	).Scan(&id, &startedAt, &endedAt, &s.DurationM, &s.Score, &s.AudioURL, &s.Advice, &createdAt)
	if err != nil {
		return sleep.Session{}, nil
	}

	s.ID = fmtID(id)
	s.UserID = 1
	s.StartedAt = startedAt.Format(time.RFC3339)
	if endedAt != nil {
		s.EndedAt = endedAt.Format(time.RFC3339)
	}
	s.CreatedAt = createdAt

	rows, err := r.pool.Query(ctx,
		`SELECT event_type, event_timestamp, level FROM sleep_events WHERE sleep_record_id = $1`, id)
	if err != nil {
		return s, nil
	}
	defer rows.Close()

	var events []sleep.Event
	for rows.Next() {
		var e sleep.Event
		if err := rows.Scan(&e.Type, &e.Timestamp, &e.Level); err != nil {
			continue
		}
		events = append(events, e)
	}
	return s, events
}

func fmtID(id int64) string {
	return fmt.Sprintf("sleep-%d", id)
}
