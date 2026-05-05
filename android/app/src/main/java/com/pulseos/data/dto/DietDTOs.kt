package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class TodayPlanDTO(
    @SerializedName("target_calories") val targetCalories: Int,
    @SerializedName("fasting_plan") val fastingPlan: FastingPlanDTO,
    val options: List<PlanOptionDTO>,
    @SerializedName("common_meals") val commonMeals: List<PlanOptionDTO>,
    @SerializedName("quick_actions") val quickActions: List<String>,
)

data class FastingPlanDTO(
    val name: String,
    val window: String,
    val description: String,
)

data class PlanOptionDTO(
    val title: String,
    val description: String,
    val items: List<String>,
)

data class AnalyzeResultDTO(
    @SerializedName("image_url") val imageURL: String?,
    @SerializedName("detected_foods") val detectedFoods: List<FoodItemDTO>,
    val recommendation: String,
    val summary: String,
    val explanation: String,
    @SerializedName("target_calories") val targetCalories: Int?,
    @SerializedName("total_calories") val totalCalories: Int,
)

data class FoodItemDTO(
    val name: String,
    val calories: Int,
    @SerializedName("protein_g") val proteinG: Double,
    @SerializedName("fat_g") val fatG: Double,
    @SerializedName("carbs_g") val carbsG: Double,
    @SerializedName("sugar_high") val sugarHigh: Boolean,
    val fried: Boolean,
)

data class AnalyzeRequestDTO(
    @SerializedName("image_url") val imageURL: String?,
    @SerializedName("meal_type") val mealType: String?,
    @SerializedName("manual_items") val manualItems: List<String>,
)

data class DietRecordDTO(
    val id: String,
    @SerializedName("image_url") val imageURL: String?,
    @SerializedName("meal_type") val mealType: String?,
    val foods: List<FoodItemDTO>,
    val recommendation: String,
    @SerializedName("total_calories") val totalCalories: Int,
)

data class PhotoUploadDTO(
    val url: String,
    val filename: String,
)
