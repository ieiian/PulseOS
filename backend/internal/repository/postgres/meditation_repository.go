package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/meditation"
)

type MeditationRepository struct {
	pool *pgxpool.Pool
}

func NewMeditationRepository(pool *pgxpool.Pool) *MeditationRepository {
	return &MeditationRepository{pool: pool}
}

func (r *MeditationRepository) SaveSession(ctx context.Context, session meditation.Session) meditation.Session {
	var id int64
	err := r.pool.QueryRow(ctx,
		`INSERT INTO meditation_sessions (user_id, mode_key, duration_s, audio_key)
		 VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		session.UserID, session.ModeKey, session.DurationS, session.AudioKey,
	).Scan(&id, &session.CreatedAt)
	if err != nil {
		session.ID = fmt.Sprintf("meditation-err-%v", err)
		return session
	}
	session.ID = fmt.Sprintf("%d", id)
	return session
}

func (r *MeditationRepository) ListSessions(ctx context.Context) []meditation.Session {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, mode_key, duration_s, audio_key, created_at
		 FROM meditation_sessions ORDER BY created_at DESC`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var sessions []meditation.Session
	for rows.Next() {
		var s meditation.Session
		var id int64
		if err := rows.Scan(&id, &s.UserID, &s.ModeKey, &s.DurationS, &s.AudioKey, &s.CreatedAt); err != nil {
			continue
		}
		s.ID = fmt.Sprintf("%d", id)
		sessions = append(sessions, s)
	}
	return sessions
}
