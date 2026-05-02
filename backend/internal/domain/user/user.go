package user

type Goal string

const (
	GoalFatLoss      Goal = "fat_loss"
	GoalMaintain     Goal = "maintain"
	GoalMuscleGain   Goal = "muscle_gain"
	GoalSleepRecover Goal = "sleep_recovery"
	GoalStressRelief Goal = "stress_relief"
)

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	HeightCM      int      `json:"height_cm"`
	WeightKG      int      `json:"weight_kg"`
	PrimaryGoal   Goal     `json:"primary_goal"`
	SecondaryGoal []Goal   `json:"secondary_goals"`
	HealthFlags   []string `json:"health_flags"`
}

type Settings struct {
	NotificationsEnabled bool `json:"notifications_enabled"`
	StepPermissionGranted bool `json:"step_permission_granted"`
	MicrophonePermissionGranted bool `json:"microphone_permission_granted"`
	SleepReminderEnabled bool `json:"sleep_reminder_enabled"`
}

type Stats struct {
	CurrentStreak int `json:"current_streak"`
	DaysTracked   int `json:"days_tracked"`
}

