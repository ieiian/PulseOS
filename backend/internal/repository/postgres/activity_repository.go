package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
)

type ActivityRepository struct {
	pool *pgxpool.Pool
}

func NewActivityRepository(pool *pgxpool.Pool) *ActivityRepository {
	return &ActivityRepository{pool: pool}
}

func (r *ActivityRepository) SaveRecord(ctx context.Context, record activity.Record) activity.Record {
	var id int64
	err := r.pool.QueryRow(ctx,
		`INSERT INTO activity_records (user_id, source, activity_type, steps, minutes, intensity, cardio_points)
		 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at`,
		record.UserID, record.Source, record.ActivityType, record.Steps,
		record.Minutes, record.Intensity, record.CardioPoints,
	).Scan(&id, &record.CreatedAt)
	if err != nil {
		return record
	}
	record.ID = fmt.Sprintf("%d", id)
	return record
}

func (r *ActivityRepository) ListRecords(ctx context.Context) []activity.Record {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, source, activity_type, steps, minutes, intensity, cardio_points, created_at
		 FROM activity_records ORDER BY created_at DESC`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var records []activity.Record
	for rows.Next() {
		var rec activity.Record
		var id int64
		if err := rows.Scan(&id, &rec.UserID, &rec.Source, &rec.ActivityType,
			&rec.Steps, &rec.Minutes, &rec.Intensity, &rec.CardioPoints, &rec.CreatedAt); err != nil {
			continue
		}
		rec.ID = fmt.Sprintf("%d", id)
		records = append(records, rec)
	}
	return records
}

func (r *ActivityRepository) ListDailyPoints(ctx context.Context) []int {
	rows, err := r.pool.Query(ctx,
		`SELECT DATE(created_at) AS d, SUM(cardio_points) FROM activity_records
		 GROUP BY d ORDER BY d ASC`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var points []int
	for rows.Next() {
		var p int
		var _date interface{}
		if err := rows.Scan(&_date, &p); err != nil {
			continue
		}
		points = append(points, p)
	}
	return points
}
