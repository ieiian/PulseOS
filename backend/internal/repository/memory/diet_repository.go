package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/diet"
)

type DietRepository struct {
	mu      sync.RWMutex
	records []diet.Record
}

func NewDietRepository() *DietRepository {
	return &DietRepository{
		records: []diet.Record{
			{
				ID:        "diet-1",
				ImageURL:  "",
				MealType:  "breakfast",
				Foods: []diet.FoodItem{
					{Name: "燕麦粥", Calories: 150, ProteinG: 5, FatG: 3, CarbsG: 27, SugarHigh: false, Fried: false},
					{Name: "香蕉", Calories: 105, ProteinG: 1, FatG: 0, CarbsG: 27, SugarHigh: true, Fried: false},
				},
				Recommendation: "营养均衡的早餐",
				TotalCalories:  255,
			},
			{
				ID:        "diet-2",
				ImageURL:  "",
				MealType:  "lunch",
				Foods: []diet.FoodItem{
					{Name: "鸡胸肉", Calories: 165, ProteinG: 31, FatG: 3.6, CarbsG: 0, SugarHigh: false, Fried: false},
					{Name: "糙米饭", Calories: 216, ProteinG: 5, FatG: 1.8, CarbsG: 45, SugarHigh: false, Fried: false},
					{Name: "西兰花", Calories: 55, ProteinG: 3.7, FatG: 0.6, CarbsG: 11, SugarHigh: false, Fried: false},
				},
				Recommendation: "高蛋白低脂午餐",
				TotalCalories:  436,
			},
		},
	}
}

func (r *DietRepository) SaveRecord(_ context.Context, record diet.Record) diet.Record {
	r.mu.Lock()
	defer r.mu.Unlock()

	record.ID = fmt.Sprintf("diet-%d", len(r.records)+1)
	r.records = append(r.records, record)
	return record
}

func (r *DietRepository) ListRecords(_ context.Context) []diet.Record {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]diet.Record, len(r.records))
	copy(out, r.records)
	return out
}
