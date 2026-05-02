package diet

type Recommendation string

const (
	RecommendationRecommended    Recommendation = "recommended"
	RecommendationCaution        Recommendation = "caution"
	RecommendationNotRecommended Recommendation = "not_recommended"
	RecommendationForbidden      Recommendation = "forbidden"
)

type FoodItem struct {
	Name      string  `json:"name"`
	Calories  int     `json:"calories"`
	ProteinG  float64 `json:"protein_g"`
	FatG      float64 `json:"fat_g"`
	CarbsG    float64 `json:"carbs_g"`
	SugarHigh bool    `json:"sugar_high"`
	Fried     bool    `json:"fried"`
}

type AnalyzeRequest struct {
	ImageURL   string   `json:"image_url"`
	MealType   string   `json:"meal_type"`
	ManualItems []string `json:"manual_items"`
}

type AnalyzeResult struct {
	ImageURL        string         `json:"image_url"`
	DetectedFoods   []FoodItem     `json:"detected_foods"`
	Recommendation  Recommendation `json:"recommendation"`
	Summary         string         `json:"summary"`
	Explanation     string         `json:"explanation"`
	TargetCalories  int            `json:"target_calories"`
	TotalCalories   int            `json:"total_calories"`
}

type Record struct {
	ID             string         `json:"id"`
	ImageURL       string         `json:"image_url"`
	MealType       string         `json:"meal_type"`
	Foods          []FoodItem     `json:"foods"`
	Recommendation Recommendation `json:"recommendation"`
	TotalCalories  int            `json:"total_calories"`
}

type PlanOption struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
}

type FastingTemplate struct {
	Name        string `json:"name"`
	Window      string `json:"window"`
	Description string `json:"description"`
}

type TodayPlan struct {
	TargetCalories int               `json:"target_calories"`
	FastingPlan    FastingTemplate   `json:"fasting_plan"`
	Options        []PlanOption      `json:"options"`
	CommonMeals    []PlanOption      `json:"common_meals"`
	QuickActions   []string          `json:"quick_actions"`
}

