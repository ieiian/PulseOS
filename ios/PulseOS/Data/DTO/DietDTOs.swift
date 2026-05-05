import Foundation

// MARK: - Today Plan (matches backend diet.TodayPlan JSON)

struct TodayPlanDTO: Codable {
    let targetCalories: Int
    let fastingPlan: FastingPlanDTO
    let options: [PlanOptionDTO]
    let commonMeals: [PlanOptionDTO]
    let quickActions: [String]

    enum CodingKeys: String, CodingKey {
        case targetCalories = "target_calories"
        case fastingPlan = "fasting_plan"
        case options
        case commonMeals = "common_meals"
        case quickActions = "quick_actions"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        targetCalories = try container.decodeIfPresent(Int.self, forKey: .targetCalories) ?? 0
        fastingPlan = try container.decodeIfPresent(FastingPlanDTO.self, forKey: .fastingPlan) ?? FastingPlanDTO(name: "", window: "", description: "")
        options = try container.decodeIfPresent([PlanOptionDTO].self, forKey: .options) ?? []
        commonMeals = try container.decodeIfPresent([PlanOptionDTO].self, forKey: .commonMeals) ?? []
        quickActions = try container.decodeIfPresent([String].self, forKey: .quickActions) ?? []
    }
}

struct FastingPlanDTO: Codable {
    let name: String
    let window: String
    let description: String
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        name = try container.decodeIfPresent(String.self, forKey: .name) ?? ""
        window = try container.decodeIfPresent(String.self, forKey: .window) ?? ""
        description = try container.decodeIfPresent(String.self, forKey: .description) ?? ""
    }
    
    init(name: String, window: String, description: String) {
        self.name = name
        self.window = window
        self.description = description
    }
}

struct PlanOptionDTO: Codable {
    let title: String
    let description: String
    let items: [String]
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        title = try container.decodeIfPresent(String.self, forKey: .title) ?? ""
        description = try container.decodeIfPresent(String.self, forKey: .description) ?? ""
        items = try container.decodeIfPresent([String].self, forKey: .items) ?? []
    }
}

// MARK: - Analyze (matches backend diet.AnalyzeResult JSON)

struct AnalyzeResultDTO: Codable {
    let imageURL: String?
    let detectedFoods: [FoodItemDTO]
    let recommendation: String
    let summary: String
    let explanation: String
    let targetCalories: Int?
    let totalCalories: Int

    enum CodingKeys: String, CodingKey {
        case imageURL = "image_url"
        case detectedFoods = "detected_foods"
        case recommendation, summary, explanation
        case targetCalories = "target_calories"
        case totalCalories = "total_calories"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        imageURL = try container.decodeIfPresent(String.self, forKey: .imageURL)
        detectedFoods = try container.decodeIfPresent([FoodItemDTO].self, forKey: .detectedFoods) ?? []
        recommendation = try container.decodeIfPresent(String.self, forKey: .recommendation) ?? ""
        summary = try container.decodeIfPresent(String.self, forKey: .summary) ?? ""
        explanation = try container.decodeIfPresent(String.self, forKey: .explanation) ?? ""
        targetCalories = try container.decodeIfPresent(Int.self, forKey: .targetCalories)
        totalCalories = try container.decodeIfPresent(Int.self, forKey: .totalCalories) ?? 0
    }
}

struct FoodItemDTO: Codable {
    let name: String
    let calories: Int
    let proteinG: Double
    let fatG: Double
    let carbsG: Double
    let sugarHigh: Bool
    let fried: Bool

    enum CodingKeys: String, CodingKey {
        case name, calories
        case proteinG = "protein_g"
        case fatG = "fat_g"
        case carbsG = "carbs_g"
        case sugarHigh = "sugar_high"
        case fried
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        name = try container.decodeIfPresent(String.self, forKey: .name) ?? ""
        calories = try container.decodeIfPresent(Int.self, forKey: .calories) ?? 0
        proteinG = try container.decodeIfPresent(Double.self, forKey: .proteinG) ?? 0.0
        fatG = try container.decodeIfPresent(Double.self, forKey: .fatG) ?? 0.0
        carbsG = try container.decodeIfPresent(Double.self, forKey: .carbsG) ?? 0.0
        sugarHigh = try container.decodeIfPresent(Bool.self, forKey: .sugarHigh) ?? false
        fried = try container.decodeIfPresent(Bool.self, forKey: .fried) ?? false
    }
}

// MARK: - Analyze Request

struct AnalyzeRequestDTO: Codable {
    let imageURL: String?
    let mealType: String?
    let manualItems: [String]

    enum CodingKeys: String, CodingKey {
        case imageURL = "image_url"
        case mealType = "meal_type"
        case manualItems = "manual_items"
    }
}

// MARK: - Record (matches backend diet.Record JSON)

struct DietRecordDTO: Codable {
    let id: String
    let imageURL: String?
    let mealType: String?
    let foods: [FoodItemDTO]
    let recommendation: String
    let totalCalories: Int

    enum CodingKeys: String, CodingKey {
        case id
        case imageURL = "image_url"
        case mealType = "meal_type"
        case foods, recommendation
        case totalCalories = "total_calories"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = try container.decodeIfPresent(String.self, forKey: .id) ?? ""
        imageURL = try container.decodeIfPresent(String.self, forKey: .imageURL)
        mealType = try container.decodeIfPresent(String.self, forKey: .mealType)
        foods = try container.decodeIfPresent([FoodItemDTO].self, forKey: .foods) ?? []
        recommendation = try container.decodeIfPresent(String.self, forKey: .recommendation) ?? ""
        totalCalories = try container.decodeIfPresent(Int.self, forKey: .totalCalories) ?? 0
    }
}

// MARK: - Photo Upload

struct PhotoUploadDTO: Codable {
    let url: String
    let filename: String
}
