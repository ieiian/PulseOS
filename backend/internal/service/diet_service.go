package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/tse/PulseOS/backend/internal/ai"
	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/domain/user"
	"github.com/tse/PulseOS/backend/internal/ruleengine"
)

type DietService struct {
	repo        DietRepo
	userRepo    UserRepo
	aiService   *ai.Service
}

func NewDietService(repo DietRepo, userRepo UserRepo, aiService *ai.Service) *DietService {
	return &DietService{
		repo:      repo,
		userRepo:  userRepo,
		aiService: aiService,
	}
}

func (s *DietService) ListRecords(ctx context.Context) []diet.Record {
	return s.repo.ListRecords(ctx)
}

func (s *DietService) GetTodayPlan(ctx context.Context) diet.TodayPlan {
	profile := s.userRepo.GetProfile(ctx)
	target := calculateTargetCalories(profile)

	return diet.TodayPlan{
		TargetCalories: target,
		FastingPlan: diet.FastingTemplate{
			Name:        "16:8 轻断食",
			Window:      "12:00 - 20:00",
			Description: "适合日常节律管理，先保持稳定窗口，不追求极端断食。",
		},
		Options: []diet.PlanOption{
			{Title: "高蛋白午餐", Description: "控制主食份量，优先蛋白质和蔬菜。", Items: []string{"鸡胸肉", "西兰花", "糙米饭"}},
			{Title: "均衡晚餐", Description: "避免过晚大份量进食。", Items: []string{"三文鱼", "南瓜", "菠菜"}},
		},
		CommonMeals: []diet.PlanOption{
			{Title: "常用早餐", Description: "快速记录模板", Items: []string{"无糖酸奶", "鸡蛋", "香蕉"}},
			{Title: "轻食午餐", Description: "快速记录模板", Items: []string{"鸡肉沙拉", "玉米", "牛油果"}},
		},
		QuickActions: []string{"拍照分析", "快速记录", "保存常用餐"},
	}
}

func (s *DietService) UploadPhoto(_ context.Context, filename string) map[string]string {
	if filename == "" {
		filename = "meal.jpg"
	}

	return map[string]string{
		"image_url": "/mock/uploads/" + filename,
		"object_key": "diet/" + filename,
	}
}

func (s *DietService) Analyze(ctx context.Context, req diet.AnalyzeRequest) diet.AnalyzeResult {
	profile := s.userRepo.GetProfile(ctx)
	foods := recognizeFoods(req)
	target := calculateTargetCalories(profile)
	total := 0

	for _, item := range foods {
		total += item.Calories
	}

	result := diet.AnalyzeResult{
		ImageURL:       req.ImageURL,
		DetectedFoods:  foods,
		Recommendation: ruleengine.EvaluateDiet(profile, foods, target),
		Summary:        buildSummary(foods, total),
		TargetCalories: target,
		TotalCalories:  total,
	}
	result.Explanation = s.aiService.ExplainDiet(result)

	s.repo.SaveRecord(ctx, diet.Record{
		ImageURL:       result.ImageURL,
		MealType:       req.MealType,
		Foods:          foods,
		Recommendation: result.Recommendation,
		TotalCalories:  total,
	})

	return result
}

func (s *DietService) QuickRecord(ctx context.Context, req diet.AnalyzeRequest) diet.Record {
	foods := recognizeFoods(req)
	total := 0
	for _, item := range foods {
		total += item.Calories
	}

	record := diet.Record{
		ImageURL:       req.ImageURL,
		MealType:       req.MealType,
		Foods:          foods,
		Recommendation: ruleengine.EvaluateDiet(s.userRepo.GetProfile(ctx), foods, calculateTargetCalories(s.userRepo.GetProfile(ctx))),
		TotalCalories:  total,
	}
	return s.repo.SaveRecord(ctx, record)
}

func recognizeFoods(req diet.AnalyzeRequest) []diet.FoodItem {
	items := req.ManualItems
	if len(items) == 0 {
		items = []string{"鸡胸肉沙拉", "糙米饭"}
	}

	foods := make([]diet.FoodItem, 0, len(items))
	for _, item := range items {
		name := strings.TrimSpace(item)
		lower := strings.ToLower(name)
		food := diet.FoodItem{Name: name, Calories: 180, ProteinG: 12, FatG: 8, CarbsG: 18}

		switch {
		case strings.Contains(lower, "salad") || strings.Contains(name, "沙拉"):
			food.Calories = 220
			food.ProteinG = 18
			food.CarbsG = 12
		case strings.Contains(name, "炸") || strings.Contains(lower, "fried"):
			food.Calories = 380
			food.Fried = true
		case strings.Contains(name, "奶茶") || strings.Contains(lower, "bubble tea") || strings.Contains(lower, "soda"):
			food.Calories = 320
			food.SugarHigh = true
		case strings.Contains(name, "蛋糕") || strings.Contains(lower, "cake"):
			food.Calories = 360
			food.SugarHigh = true
		case strings.Contains(name, "鸡胸") || strings.Contains(lower, "chicken"):
			food.Calories = 210
			food.ProteinG = 28
			food.CarbsG = 4
		case strings.Contains(name, "米饭") || strings.Contains(lower, "rice"):
			food.Calories = 260
			food.CarbsG = 48
		}

		foods = append(foods, food)
	}

	return foods
}

func calculateTargetCalories(profile user.Profile) int {
	if profile.PrimaryGoal == user.GoalFatLoss {
		return 1600
	}
	if profile.PrimaryGoal == user.GoalMuscleGain {
		return 2300
	}
	return 1900
}

func buildSummary(foods []diet.FoodItem, total int) string {
	names := make([]string, 0, len(foods))
	for _, food := range foods {
		names = append(names, food.Name)
	}

	return "本次记录包含：" + strings.Join(names, "、") + "。总热量约 " + strconv.Itoa(total) + " 千卡。"
}
