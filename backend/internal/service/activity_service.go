package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/activity"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
	"github.com/tse/PulseOS/backend/internal/ruleengine"
)

type ActivityService struct {
	repo *postgres.ActivityRepository
}

func NewActivityService(repo *postgres.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
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

	status := "on_track"
	if total >= 150 {
		status = "achieved"
	}

	return activity.WeekSummary{
		DailyPoints: []int{12, 18, 0, 35, 20, 24, total},
		TotalPoints: total,
		WeeklyGoal:  150,
		Status:      status,
		Tips: []string{
			"优先补足中等强度分钟数。",
			"工作日安排短时快走，比周末突击更稳定。",
		},
	}
}
