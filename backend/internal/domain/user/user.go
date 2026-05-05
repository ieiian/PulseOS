package user

import "time"

type Goal string

const (
	GoalFatLoss      Goal = "fat_loss"
	GoalMaintain     Goal = "maintain"
	GoalMuscleGain   Goal = "muscle_gain"
	GoalSleepRecover Goal = "sleep_recovery"
	GoalStressRelief Goal = "stress_relief"
)

type Profile struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Age           int       `json:"age"`
	Gender        string    `json:"gender"`
	HeightCM      int       `json:"height_cm"`
	WeightKG      int       `json:"weight_kg"`
	PrimaryGoal   Goal      `json:"primary_goal"`
	SecondaryGoal []Goal    `json:"secondary_goals"`
	HealthFlags   []string  `json:"health_flags"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Settings struct {
	ID                          int64     `json:"id"`
	UserID                      int64     `json:"user_id"`
	NotificationsEnabled        bool      `json:"notifications_enabled"`
	StepPermissionGranted       bool      `json:"step_permission_granted"`
	MicrophonePermissionGranted bool      `json:"microphone_permission_granted"`
	SleepReminderEnabled        bool      `json:"sleep_reminder_enabled"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
}

type Stats struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	CurrentStreak int       `json:"current_streak"`
	DaysTracked   int       `json:"days_tracked"`
	UpdatedAt     time.Time `json:"updated_at"`
}
