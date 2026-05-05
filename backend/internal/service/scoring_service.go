package service

import (
	"context"
	"time"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/domain/scoring"
)

type ScoringService struct {
	repo              ScoringRepo
	dietService       *DietService
	activityService   *ActivityService
	sleepService      *SleepService
	meditationService *MeditationService
}

func NewScoringService(
	repo ScoringRepo,
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
	dietRecords := s.dietService.ListRecords(ctx)
	dietScore := calculateDietScore(dietRecords)

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
		Trends:          buildTrends(s.repo, score.TotalScore),
	}
}

func calculateDietScore(records []diet.Record) int {
	if len(records) == 0 {
		return 0
	}
	score := 50
	for _, r := range records {
		switch r.Recommendation {
		case diet.RecommendationRecommended:
			score += 10
		case diet.RecommendationCaution:
			score -= 5
		case diet.RecommendationNotRecommended, diet.RecommendationForbidden:
			score -= 15
		}
	}
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}
	return score
}

func buildTrends(repo ScoringRepo, todayScore int) []int {
	past := repo.GetHistory(context.Background())
	trends := make([]int, 0, len(past)+1)
	for _, s := range past {
		trends = append(trends, s.TotalScore)
	}
	trends = append(trends, todayScore)
	return trends
}
