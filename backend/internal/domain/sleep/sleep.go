package sleep

type EventType string

const (
	EventSnore EventType = "snore"
	EventTalk  EventType = "talk"
)

type Session struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
	DurationM int    `json:"duration_m"`
	Score     int    `json:"score"`
	AudioURL  string `json:"audio_url"`
	Advice    string `json:"advice"`
}

type Event struct {
	Type      EventType `json:"type"`
	Timestamp string    `json:"timestamp"`
	Level     string    `json:"level"`
}

type StartRequest struct {
	AudioURL string `json:"audio_url"`
}

type EndRequest struct {
	SessionID string `json:"session_id"`
}

type TodaySummary struct {
	Session Session `json:"session"`
	Events  []Event `json:"events"`
}
