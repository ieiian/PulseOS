package service

import (
	"context"
	"time"

	"github.com/tse/PulseOS/backend/internal/domain/scoring"
	"github.com/tse/PulseOS/backend/internal/repository/postgres"
)

type ScoringService struct {
	repo              *postgres.ScoringRepository
	dietService       *DietService
	activityService   *ActivityService
	sleepService      *SleepService
	meditationService *MeditationService
}

func NewScoringService(
	repo *postgres.ScoringRepository,
	dietService *DietService,
	activityService *ActivityService,
	sleepService *SleepService,
	meditationService *MeditationService,
) *ScoringService {
	return &ScoringService{
		repo:              repo,
		dietService:       dietService,
		activityService:   activityService,
		sleepService:      sleepService,
		meditationService: meditationService,
	}
}

func (s *ScoringService) CalculateToday(ctx context.Context) scoring.DailyScore {
	dietScore := 72

	activitySummary := s.activityService.GetTodaySummary(ctx)
	activityScore := activitySummary.CardioPoints
	if activityScore > 100 {
		activityScore = 100
	}

	sleepSummary := s.sleepService.GetToday(ctx)
	sleepScore := sleepSummary.Session.Score
	if sleepScore == 0 {
		sleepScore = 55
	}

	total := (dietScore + activityScore + sleepScore) / 3
	score := scoring.DailyScore{
		Date:          time.Now().Format("2006-01-02"),
		DietScore:     dietScore,
		ActivityScore: activityScore,
		SleepScore:    sleepScore,
		TotalScore:    total,
	}
	return s.repo.Save(ctx, score)
}

func (s *ScoringService) BuildDashboard(ctx context.Context) scoring.Dashboard {
	score := s.CalculateToday(ctx)
	dietPlan := s.dietService.GetTodayPlan(ctx)
	activitySummary := s.activityService.GetTodaySummary(ctx)
	sleepSummary := s.sleepService.GetToday(ctx)
	meditationSummary := s.meditationService.GetTodaySummary(ctx)

	return scoring.Dashboard{
		Today:           score,
		ActionItem:      "优先完成今天的运动积分和一次饮食记录。",
		DietSummary:     "目标热量 " + itoa(dietPlan.TargetCalories) + " 千卡，已配置轻断食窗口。",
		ActivitySummary: "今日 " + itoa(activitySummary.Steps) + " 步，心肺强化 " + itoa(activitySummary.CardioPoints) + " 分。",
		SleepSummary:    "昨夜评分 " + itoa(sleepSummary.Session.Score) + "，时长 " + itoa(sleepSummary.Session.DurationM) + " 分钟。",
		MeditationNote:  "今日冥想 " + itoa(meditationSummary.TotalDurationS/60) + " 分钟。",
		Trends:          []int{58, 63, 61, 70, 68, 74, score.TotalScore},
	}
}
