package provider

import (
	"fmt"
	"strings"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
)

type MockProvider struct{}

func (p MockProvider) ExplainDiet(result diet.AnalyzeResult) string {
	names := make([]string, 0, len(result.DetectedFoods))
	for _, item := range result.DetectedFoods {
		names = append(names, item.Name)
	}

	return fmt.Sprintf("识别到 %s，总热量约 %d 千卡。当前判断为 %s，建议结合今天目标热量和进食频次控制份量。",
		strings.Join(names, "、"),
		result.TotalCalories,
		result.Recommendation,
	)
}

