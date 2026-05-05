package scoring

type DailyScore struct {
	Date          string `json:"date"`
	UserID        int64  `json:"user_id"`
	DietScore     int    `json:"diet_score"`
	ActivityScore int    `json:"activity_score"`
	SleepScore    int    `json:"sleep_score"`
	TotalScore    int    `json:"total_score"`
}

type Dashboard struct {
	Today           DailyScore `json:"today"`
	ActionItem      string     `json:"action_item"`
	DietSummary     string     `json:"diet_summary"`
	ActivitySummary string     `json:"activity_summary"`
	SleepSummary    string     `json:"sleep_summary"`
	MeditationNote  string     `json:"meditation_note"`
	Trends          []int      `json:"trends"`
}
