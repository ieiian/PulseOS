package postgres

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tse/PulseOS/backend/internal/domain/user"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) SaveProfile(ctx context.Context, profile user.Profile) user.Profile {
	secondaryGoals, _ := json.Marshal(profile.SecondaryGoal)
	healthFlags, _ := json.Marshal(profile.HealthFlags)

	if profile.ID == 0 {
		err := r.pool.QueryRow(ctx,
			`INSERT INTO users (name, age, gender, height_cm, weight_kg, primary_goal, secondary_goals, health_flags)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`,
			profile.Name, profile.Age, profile.Gender, profile.HeightCM, profile.WeightKG,
			profile.PrimaryGoal, secondaryGoals, healthFlags,
		).Scan(&profile.ID, &profile.CreatedAt, &profile.UpdatedAt)
		if err != nil {
			return profile
		}
	} else {
		_ = r.pool.QueryRow(ctx,
			`UPDATE users SET name=$1, age=$2, gender=$3, height_cm=$4, weight_kg=$5,
			 primary_goal=$6, secondary_goals=$7, health_flags=$8, updated_at=NOW()
			 WHERE id=$9 RETURNING updated_at`,
			profile.Name, profile.Age, profile.Gender, profile.HeightCM, profile.WeightKG,
			profile.PrimaryGoal, secondaryGoals, healthFlags, profile.ID,
		).Scan(&profile.UpdatedAt)
	}
	return profile
}

func (r *UserRepository) GetProfile(ctx context.Context) user.Profile {
	var p user.Profile
	var secondaryGoals, healthFlags []byte

	err := r.pool.QueryRow(ctx,
		`SELECT id, name, age, gender, height_cm, weight_kg, primary_goal, secondary_goals, health_flags, created_at, updated_at
		 FROM users ORDER BY id LIMIT 1`,
	).Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.HeightCM, &p.WeightKG,
		&p.PrimaryGoal, &secondaryGoals, &healthFlags, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p
	}
	_ = json.Unmarshal(secondaryGoals, &p.SecondaryGoal)
	_ = json.Unmarshal(healthFlags, &p.HealthFlags)
	return p
}

func (r *UserRepository) SaveSettings(ctx context.Context, settings user.Settings) user.Settings {
	if settings.ID == 0 {
		_ = r.pool.QueryRow(ctx,
			`INSERT INTO user_settings (user_id, notifications_enabled, step_permission_granted, microphone_permission_granted, sleep_reminder_enabled)
			 VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`,
			settings.UserID, settings.NotificationsEnabled, settings.StepPermissionGranted,
			settings.MicrophonePermissionGranted, settings.SleepReminderEnabled,
		).Scan(&settings.ID, &settings.CreatedAt, &settings.UpdatedAt)
	} else {
		_ = r.pool.QueryRow(ctx,
			`UPDATE user_settings SET notifications_enabled=$1, step_permission_granted=$2,
			 microphone_permission_granted=$3, sleep_reminder_enabled=$4, updated_at=NOW()
			 WHERE id=$5 RETURNING updated_at`,
			settings.NotificationsEnabled, settings.StepPermissionGranted,
			settings.MicrophonePermissionGranted, settings.SleepReminderEnabled, settings.ID,
		).Scan(&settings.UpdatedAt)
	}
	return settings
}

func (r *UserRepository) GetSettings(ctx context.Context) user.Settings {
	var s user.Settings
	_ = r.pool.QueryRow(ctx,
		`SELECT id, user_id, notifications_enabled, step_permission_granted, microphone_permission_granted, sleep_reminder_enabled, created_at, updated_at
		 FROM user_settings ORDER BY id LIMIT 1`,
	).Scan(&s.ID, &s.UserID, &s.NotificationsEnabled, &s.StepPermissionGranted,
		&s.MicrophonePermissionGranted, &s.SleepReminderEnabled, &s.CreatedAt, &s.UpdatedAt)
	return s
}

func (r *UserRepository) GetStats(ctx context.Context) user.Stats {
	var s user.Stats
	_ = r.pool.QueryRow(ctx,
		`SELECT id, user_id, current_streak, days_tracked, updated_at
		 FROM user_stats ORDER BY id LIMIT 1`,
	).Scan(&s.ID, &s.UserID, &s.CurrentStreak, &s.DaysTracked, &s.UpdatedAt)
	return s
}
