package ruleengine

import (
	"strings"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
	"github.com/tse/PulseOS/backend/internal/domain/user"
)

func EvaluateDiet(profile user.Profile, foods []diet.FoodItem, targetCalories int) diet.Recommendation {
	totalCalories := 0
	hasHighSugar := false
	hasFried := false
	hasStrictFlag := false

	for _, food := range foods {
		totalCalories += food.Calories
		hasHighSugar = hasHighSugar || food.SugarHigh
		hasFried = hasFried || food.Fried
	}

	for _, flag := range profile.HealthFlags {
		if strings.Contains(strings.ToLower(flag), "diabetes") {
			hasStrictFlag = true
			break
		}
	}

	if hasStrictFlag && hasHighSugar {
		return diet.RecommendationForbidden
	}

	if totalCalories > targetCalories+200 {
		return diet.RecommendationNotRecommended
	}

	if hasFried || hasHighSugar {
		return diet.RecommendationCaution
	}

	return diet.RecommendationRecommended
}

