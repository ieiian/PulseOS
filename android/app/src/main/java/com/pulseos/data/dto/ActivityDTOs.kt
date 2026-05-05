package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class ActivityTodaySummaryDTO(
    val steps: Int,
    @SerializedName("cardio_points") val cardioPoints: Int,
    @SerializedName("step_goal") val stepGoal: Int,
    @SerializedName("weekly_goal") val weeklyGoal: Int,
    @SerializedName("remaining_points") val remainingPoints: Int,
    val reminder: String,
    @SerializedName("recent_activities") val recentActivities: List<ActivityRecordDTO>,
)

data class ActivityWeekSummaryDTO(
    @SerializedName("daily_points") val dailyPoints: List<Int>,
    @SerializedName("total_points") val totalPoints: Int,
    @SerializedName("weekly_goal") val weeklyGoal: Int,
    val status: String,
    val tips: List<String>,
)

data class ActivityRecordDTO(
    val id: String,
    val source: String,
    @SerializedName("activity_type") val activityType: String,
    val steps: Int,
    val minutes: Int,
    val intensity: String,
    @SerializedName("cardio_points") val cardioPoints: Int,
)

data class ManualActivityRequestDTO(
    @SerializedName("activity_type") val activityType: String,
    val minutes: Int,
    val intensity: String,
    val steps: Int,
)
