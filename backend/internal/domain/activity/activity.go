package activity

import "time"

type Intensity string

const (
	IntensityLight    Intensity = "light"
	IntensityModerate Intensity = "moderate"
	IntensityVigorous Intensity = "vigorous"
)

type Record struct {
	ID           string     `json:"id"`
	UserID       int64      `json:"user_id"`
	Source       string     `json:"source"`
	ActivityType string     `json:"activity_type"`
	Steps        int        `json:"steps"`
	Minutes      int        `json:"minutes"`
	Intensity    Intensity  `json:"intensity"`
	CardioPoints int        `json:"cardio_points"`
	CreatedAt    time.Time  `json:"created_at"`
}

type ManualRecordRequest struct {
	ActivityType string    `json:"activity_type"`
	Minutes      int       `json:"minutes"`
	Intensity    Intensity `json:"intensity"`
	Steps        int       `json:"steps"`
}

type TodaySummary struct {
	Steps            int      `json:"steps"`
	CardioPoints     int      `json:"cardio_points"`
	StepGoal         int      `json:"step_goal"`
	WeeklyGoal       int      `json:"weekly_goal"`
	RemainingPoints  int      `json:"remaining_points"`
	Reminder         string   `json:"reminder"`
	RecentActivities []Record `json:"recent_activities"`
}

type WeekSummary struct {
	DailyPoints []int    `json:"daily_points"`
	TotalPoints int      `json:"total_points"`
	WeeklyGoal  int      `json:"weekly_goal"`
	Status      string   `json:"status"`
	Tips        []string `json:"tips"`
}
