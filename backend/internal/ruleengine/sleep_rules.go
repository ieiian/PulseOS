package ruleengine

import "github.com/tse/PulseOS/backend/internal/domain/sleep"

func DetectSleepEvents() []sleep.Event {
	return []sleep.Event{
		{Type: sleep.EventSnore, Timestamp: "01:12", Level: "medium"},
		{Type: sleep.EventTalk, Timestamp: "03:46", Level: "low"},
		{Type: sleep.EventSnore, Timestamp: "05:21", Level: "high"},
	}
}

func CalculateSleepScore(durationM int, events []sleep.Event) int {
	score := 85

	if durationM < 420 {
		score -= 15
	}
	score -= len(events) * 5

	if score < 0 {
		return 0
	}
	return score
}

func BuildSleepAdvice(durationM int, events []sleep.Event) string {
	if durationM < 420 {
		return "昨夜睡眠时长偏短，今晚建议提前 30 分钟入睡，并减少临睡前进食。"
	}
	if len(events) >= 3 {
		return "昨夜呼噜/梦话事件偏多，建议观察连续几晚情况，并避免睡前饮酒。"
	}
	return "昨夜整体睡眠稳定，继续保持固定入睡时间。"
}
