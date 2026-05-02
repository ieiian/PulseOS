package meditation

type BreathMode struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	InhaleSec   int    `json:"inhale_sec"`
	HoldSec     int    `json:"hold_sec"`
	ExhaleSec   int    `json:"exhale_sec"`
	Description string `json:"description"`
}

type SessionRequest struct {
	ModeKey   string `json:"mode_key"`
	DurationS int    `json:"duration_s"`
	AudioKey  string `json:"audio_key"`
}

type Session struct {
	ID        string `json:"id"`
	ModeKey   string `json:"mode_key"`
	DurationS int    `json:"duration_s"`
	AudioKey  string `json:"audio_key"`
}

type TodaySummary struct {
	TotalDurationS int          `json:"total_duration_s"`
	CompletedCount int          `json:"completed_count"`
	RecentSessions []Session    `json:"recent_sessions"`
	Modes          []BreathMode `json:"modes"`
}
