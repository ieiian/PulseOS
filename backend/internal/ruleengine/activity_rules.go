package ruleengine

import "github.com/tse/PulseOS/backend/internal/domain/activity"

func CalculateCardioPoints(minutes int, intensity activity.Intensity) int {
	switch intensity {
	case activity.IntensityVigorous:
		return minutes * 2
	case activity.IntensityModerate:
		return minutes
	default:
		return minutes / 2
	}
}

func BuildActivityReminder(steps int, totalPoints int, weeklyGoal int) string {
	if totalPoints >= weeklyGoal {
		return "本周心肺强化目标已达成，继续保持稳定节律。"
	}
	if steps < 6000 {
		return "今天步数偏低，建议补一段 20 分钟快走。"
	}
	if totalPoints < 60 {
		return "本周积分还在起步阶段，优先安排 30 分钟中等强度运动。"
	}
	return "今天完成度不错，继续补足本周目标。"
}
