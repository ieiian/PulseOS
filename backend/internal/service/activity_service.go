package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
	"github.com/tse/PulseOS/backend/internal/ruleengine"
)

type ActivityService struct {
	repo ActivityRepo
}

func NewActivityService(repo ActivityRepo) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) ListRecords(ctx context.Context) []activity.Record {
	return s.repo.ListRecords(ctx)
}

func (s *ActivityService) RecordManualActivity(ctx context.Context, req activity.ManualRecordRequest) activity.Record {
	record := activity.Record{
		Source:       "manual",
		ActivityType: req.ActivityType,
		Steps:        req.Steps,
		Minutes:      req.Minutes,
		Intensity:    req.Intensity,
		CardioPoints: ruleengine.CalculateCardioPoints(req.Minutes, req.Intensity),
	}
	return s.repo.SaveRecord(ctx, record)
}

func (s *ActivityService) GetTodaySummary(ctx context.Context) activity.TodaySummary {
	records := s.repo.ListRecords(ctx)
	totalSteps := 0
	totalPoints := 0
	for _, record := range records {
		totalSteps += record.Steps
		totalPoints += record.CardioPoints
	}

	weeklyGoal := 150
	remaining := weeklyGoal - totalPoints
	if remaining < 0 {
		remaining = 0
	}

	return activity.TodaySummary{
		Steps:            totalSteps,
		CardioPoints:     totalPoints,
		StepGoal:         8000,
		WeeklyGoal:       weeklyGoal,
		RemainingPoints:  remaining,
		Reminder:         ruleengine.BuildActivityReminder(totalSteps, totalPoints, weeklyGoal),
		RecentActivities: records,
	}
}

func (s *ActivityService) GetWeekSummary(ctx context.Context) activity.WeekSummary {
	records := s.repo.ListRecords(ctx)
	total := 0
	for _, record := range records {
		total += record.CardioPoints
	}

	weeklyGoal := 150
	status := "on_track"
	if total >= weeklyGoal {
		status = "achieved"
	}

	dailyPoints := s.repo.ListDailyPoints(ctx)

	var tips []string
	if total < weeklyGoal/2 {
		tips = []string{
			"本周活动量偏低，建议每天安排 20 分钟快走。",
			"工作日安排短时快走，比周末突击更稳定。",
		}
	} else if total < weeklyGoal {
		tips = []string{
			"距离周目标还差 " + itoa(weeklyGoal-total) + " 分，继续加油。",
		}
	} else {
		tips = []string{"本周目标已达成，保持节奏。"}
	}

	return activity.WeekSummary{
		DailyPoints: dailyPoints,
		TotalPoints: total,
		WeeklyGoal:  weeklyGoal,
		Status:      status,
		Tips:        tips,
	}
}
