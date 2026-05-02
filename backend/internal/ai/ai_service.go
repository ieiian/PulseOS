package ai

import (
	"github.com/tse/PulseOS/backend/internal/ai/provider"
	"github.com/tse/PulseOS/backend/internal/domain/diet"
)

type DietExplainer interface {
	ExplainDiet(result diet.AnalyzeResult) string
}

type Service struct {
	explainer DietExplainer
}

func NewService() *Service {
	return &Service{
		explainer: provider.MockProvider{},
	}
}

func (s *Service) ExplainDiet(result diet.AnalyzeResult) string {
	return s.explainer.ExplainDiet(result)
}

