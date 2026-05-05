package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
)

type DietRepository struct {
	pool *pgxpool.Pool
}

func NewDietRepository(pool *pgxpool.Pool) *DietRepository {
	return &DietRepository{pool: pool}
}

func (r *DietRepository) SaveRecord(ctx context.Context, record diet.Record) diet.Record {
	foods, _ := json.Marshal(record.Foods)

	var id int64
	err := r.pool.QueryRow(ctx,
		`INSERT INTO food_records (user_id, image_url, meal_type, recommendation, total_calories, detected_food)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`,
		record.UserID, record.ImageURL, record.MealType, record.Recommendation,
		record.TotalCalories, foods,
	).Scan(&id, &record.CreatedAt)
	if err != nil {
		record.ID = fmt.Sprintf("diet-err-%v", err)
		return record
	}
	record.ID = fmt.Sprintf("%d", id)
	return record
}

func (r *DietRepository) ListRecords(ctx context.Context) []diet.Record {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, image_url, meal_type, recommendation, total_calories, detected_food, created_at
		 FROM food_records ORDER BY created_at DESC`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var records []diet.Record
	for rows.Next() {
		var rec diet.Record
		var foods []byte
		var id int64
		if err := rows.Scan(&id, &rec.UserID, &rec.ImageURL, &rec.MealType,
			&rec.Recommendation, &rec.TotalCalories, &foods, &rec.CreatedAt); err != nil {
			continue
		}
		rec.ID = fmt.Sprintf("%d", id)
		_ = json.Unmarshal(foods, &rec.Foods)
		records = append(records, rec)
	}
	return records
}
