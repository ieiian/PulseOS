package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/scoring"
)

type ScoringRepository struct {
	pool *pgxpool.Pool
}

func NewScoringRepository(pool *pgxpool.Pool) *ScoringRepository {
	return &ScoringRepository{pool: pool}
}

func (r *ScoringRepository) Save(ctx context.Context, score scoring.DailyScore) scoring.DailyScore {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO daily_scores (user_id, score_date, diet_score, activity_score, sleep_score, total_score)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 ON CONFLICT (user_id, score_date) DO UPDATE SET
		   diet_score = EXCLUDED.diet_score,
		   activity_score = EXCLUDED.activity_score,
		   sleep_score = EXCLUDED.sleep_score,
		   total_score = EXCLUDED.total_score`,
		score.UserID, score.Date, score.DietScore, score.ActivityScore, score.SleepScore, score.TotalScore)
	if err != nil {
		return score
	}
	return score
}

func scanDailyScore(scanner interface{ Scan(...interface{}) error }) (scoring.DailyScore, error) {
	var s scoring.DailyScore
	var scoreDate time.Time
	if err := scanner.Scan(&scoreDate, &s.UserID, &s.DietScore, &s.ActivityScore, &s.SleepScore, &s.TotalScore); err != nil {
		return s, err
	}
	s.Date = scoreDate.Format("2006-01-02")
	return s, nil
}

func (r *ScoringRepository) Get(ctx context.Context) scoring.DailyScore {
	s, _ := scanDailyScore(r.pool.QueryRow(ctx,
		`SELECT score_date, user_id, diet_score, activity_score, sleep_score, total_score
		 FROM daily_scores ORDER BY score_date DESC LIMIT 1`))
	return s
}

func (r *ScoringRepository) GetHistory(ctx context.Context) []scoring.DailyScore {
	rows, err := r.pool.Query(ctx,
		`SELECT score_date, user_id, diet_score, activity_score, sleep_score, total_score
		 FROM daily_scores ORDER BY score_date ASC`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var scores []scoring.DailyScore
	for rows.Next() {
		s, err := scanDailyScore(rows)
		if err != nil {
			continue
		}
		scores = append(scores, s)
	}
	return scores
}
